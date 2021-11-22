package docker

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gosuri/uilive"
)

type Docker struct {
	dockerClient *client.Client
}

func New() (*Docker, error) {
	dockerCli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return &Docker{
		dockerClient: dockerCli,
	}, nil
}

func (d *Docker) Tag(context context.Context, source string, target string) error {
	return d.dockerClient.ImageTag(context, source, target)
}

type PushResponse struct {
	Status   string
	ID       string
	Progress string
}

func (d *Docker) PushImageAnonymous(context context.Context, image string) error {
	authConfig := types.AuthConfig{
		Username: "docker",
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	reader, errPush := d.dockerClient.ImagePush(context, image, types.ImagePushOptions{RegistryAuth: authStr})

	if errPush != nil {
		return errPush
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	var response PushResponse
	writer := uilive.New()
	writer.Start()
	for scanner.Scan() {
		jsonErr := json.Unmarshal(scanner.Bytes(), &response)
		if jsonErr != nil {
			fmt.Printf("Docker push: Error while parsing json response. %s\n", jsonErr)
		} else {
			fmt.Fprintf(writer, "Docker push: %s %s %s\n", response.ID, response.Status, response.Progress)
		}
	}
	writer.Stop()

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
