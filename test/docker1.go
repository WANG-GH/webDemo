package main

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	//"fmt"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	binds_arg := []string{"/home/yeye/program/webDemo/user_submit:/go/volume"}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Env:   []string{"PROGRAM=test1.go"},
		Image: "web:2",
		Cmd:   []string{"/go/check.sh"},
	},
		&container.HostConfig{
			Binds:        binds_arg,
			VolumeDriver: "/home/yeye/program/webDemo/user_submit",
		}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	//statusCh
	_, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	// case <-statusCh:
	case <-time.NewTimer(time.Second * 10).C:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}
