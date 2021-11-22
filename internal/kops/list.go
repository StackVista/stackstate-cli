package kops

import (
	"context"
	"fmt"
	"strings"

	"gitlab.com/stackvista/stackstate-cli2/internal/aws"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/rs/zerolog/log"
)

const KopsClusterMarkerObject = "/config"

func ListClusters(ctx context.Context, profile string) ([]string, error) {
	cfg, err := LoadKopsConfig()
	if err != nil {
		return nil, err
	}

	clusters := []string{}

	if profile != "" {
		c, err := listClustersInProfile(ctx, cfg, profile)
		if err != nil {
			return nil, err
		}

		clusters = append(clusters, c...)
	} else {
		profiles, err := aws.ListAwsProfiles(ctx)
		if err != nil {
			return nil, err
		}

		for _, profile := range profiles {
			c, err := listClustersInProfile(ctx, cfg, profile)
			if err != nil {
				return nil, err
			}

			clusters = append(clusters, c...)
		}
	}

	return clusters, nil
}

func listClustersInProfile(ctx context.Context, cfg *Config, profile string) ([]string, error) {
	logger := log.Ctx(ctx).With().Str("profile", profile).Logger()
	ctx = logger.WithContext(ctx)
	logger.Info().Msg("Listing Kops clusters")

	awsConfig, err := config.LoadDefaultConfig(ctx, config.WithRegion(""), config.WithSharedConfigProfile(profile))
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(awsConfig)

	out, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		logger.Debug().Err(err).Msg("An error occurred listing buckets for profile")
		logger.Warn().Msg("Cannot list S3 Buckets")
		return []string{}, nil
	}

	bucketNames := []string{}
	for _, bucket := range out.Buckets {
		bucketNames = append(bucketNames, *bucket.Name)
	}

	logger.Debug().Interface("buckets", bucketNames).Msg("Found S3 buckets in profile")

	clusters := []string{}

	for _, b := range out.Buckets {
		logger.Trace().Str("bucket", *b.Name).Msg("Scanning bucket for Kops clusters")

		if strings.HasSuffix(*b.Name, cfg.Kops.BucketSuffix) {
			bl, err := client.GetBucketLocation(ctx, &s3.GetBucketLocationInput{Bucket: b.Name})
			if err != nil {
				logger.Debug().Str("bucket", *b.Name).Err(err).Msg("An error occurred getting the S3 Bucket location")
				logger.Warn().Str("bucket", *b.Name).Msg("Cannot get S3 Bucket location")
				continue
			}

			klusters, err := listKopsClusters(ctx, b, bl.LocationConstraint, client)
			if err != nil {
				logger.Debug().Str("bucket", *b.Name).Err(err).Msg("An error occurred getting the S3 Bucket contents")
				logger.Warn().Str("bucket", *b.Name).Msg("Cannot get S3 Bucket contents")
				continue
			}

			logger.Debug().Str("bucket", *b.Name).Interface("clusters", klusters).Msg("Found kops clusters")

			clusters = append(clusters, klusters...)
		}
	}

	return clusters, nil
}

func listKopsClusters(ctx context.Context, bucket s3types.Bucket, loc s3types.BucketLocationConstraint, client *s3.Client) ([]string, error) {
	// Special case, if the bucket is in us-east-1, the BucketLocationConstraint will be empty
	region := string(loc)
	if region == "" {
		region = "us-east-1"
	}

	logger := log.Ctx(ctx).With().Str("bucket", *bucket.Name).Str("location", region).Logger()

	d := "/"

	setopts := func(o *s3.Options) {
		o.Region = region
		o.UsePathStyle = true
	}
	names := []string{}

	logger.Info().Msg("Checking bucket contents ")
	output, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:    bucket.Name,
		Delimiter: &d,
	}, setopts)
	if err != nil {
		return nil, err
	}

	logger.Debug().Interface("prefixes", output.CommonPrefixes).Msg("Root contents of bucket")

	for _, prefix := range output.CommonPrefixes {
		clusterName := *prefix.Prefix
		if strings.HasSuffix(clusterName, "/") {
			clusterName = clusterName[0 : len(clusterName)-1]
		}

		logger.Info().Str("cluster-name", clusterName).Msg("Checking to see whether it is a KOPS cluster definition")

		key := fmt.Sprintf("%s/config", clusterName)

		res, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
			Bucket: bucket.Name,
			Prefix: &key,
		}, setopts)

		if err != nil {
			return nil, err
		}

		if len(res.Contents) == 1 {
			names = append(names, clusterName)
		}
	}

	return names, nil
}
