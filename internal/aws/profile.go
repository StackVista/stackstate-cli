package aws

import (
	"bufio"
	"context"
	"io"
	"os"
	"regexp"

	home "github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog/log"
)

var (
	ConfigProfileRegex = regexp.MustCompile(`^\s*\[profile ([^\[\]]+)\]\s*$`)
	CredProfileRegex   = regexp.MustCompile(`^\s*\[([^\[\]]+\])\s*$`)
)

func sharedConfigFile() (string, error) {
	return home.Expand("~/.aws/config")
}

func parseProfilesFromConfig(ctx context.Context) ([]string, error) {
	f, err := sharedConfigFile()
	if err != nil {
		return nil, err
	}

	log.Ctx(ctx).Debug().Str("file", f).Msg("Parsing AWS config file")

	lines, err := readFile(f)
	if err != nil {
		return nil, err
	}

	return parseProfiles(lines, ConfigProfileRegex)
}

func parseProfilesFromCredentials(ctx context.Context) ([]string, error) {
	f, err := sharedCredentialsFile()
	if err != nil {
		return nil, err
	}

	log.Ctx(ctx).Debug().Str("file", f).Msg("Parsing AWS credentials file")

	lines, err := readFile(f)
	if err != nil {
		return nil, err
	}

	return parseProfiles(lines, CredProfileRegex)
}

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	r := bufio.NewReader(file)
	lines := []string{}

	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			return lines, err
		}

		lines = append(lines, line)

		if err != nil {
			break
		}
	}

	return lines, nil
}

func parseProfiles(lines []string, regex *regexp.Regexp) ([]string, error) {
	profiles := []string{}

	for _, l := range lines {
		if regex.MatchString(l) {
			m := regex.FindStringSubmatch(l)
			profiles = append(profiles, m[1])
		}
	}

	return profiles, nil
}

func sharedCredentialsFile() (string, error) {
	return home.Expand("~/.aws/credentials")
}

func ListAwsProfiles(ctx context.Context) ([]string, error) {
	log.Ctx(ctx).Debug().Msg("Parsing AWS config files for profiles")

	profiles, err := parseProfilesFromConfig(ctx)
	if err != nil && os.IsNotExist(err) {
		profiles, err = parseProfilesFromCredentials(ctx)
	}

	if err != nil && os.IsNotExist(err) {
		return []string{}, nil
	}

	if err != nil {
		return nil, err
	}

	log.Ctx(ctx).Debug().Interface("profiles", profiles).Msg("Parsed profiles")

	return profiles, nil
}
