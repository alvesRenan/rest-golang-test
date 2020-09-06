package utils

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var ctx context.Context = context.Background()
var cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

func createClient(name, net string) (container.ContainerCreateCreatedBody, error) {

	return cli.ContainerCreate(
		ctx, &container.Config{
			Image: "renanalves/android-22:v2",
			Cmd:   []string{"/root/port_forward.sh"},
			Tty:   true,
		}, &container.HostConfig{
			Privileged: true,
		}, nil, name)
}

func createServer(name string) (container.ContainerCreateCreatedBody, error) {

	return cli.ContainerCreate(
		ctx, &container.Config{
			Image: "renanalves/server-testbed:configured",
			Tty:   true,
		}, nil, nil, name)
}

// CreateContainer creates a new container
func CreateContainer(name, net string, isServer bool) {

	var newContainer container.ContainerCreateCreatedBody
	var err error

	if isServer {
		newContainer, err = createServer(name)
	} else {
		newContainer, err = createClient(name, net)
	}

	if err != nil {
		panic(err)
	}

	cli.ContainerStart(ctx, newContainer.ID, types.ContainerStartOptions{})
}
