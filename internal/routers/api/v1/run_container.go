package v1

import (
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/oci"
	"github.com/gofiber/fiber/v2"
	"time"
)

const DOCKER_HUB_PREFIX = "docker.io"

func RunContainer(c *fiber.Ctx) error {
	imageName := c.Params("image")
	var ctx = c.UserContext()
	var client, err = new_client()
	if err != nil {
		return err
	}
	ls, err := client.ListImages(ctx, imageName)
	if err != nil {
		return err
	}
	var image containerd.Image
	if len(ls) < 1 {
		// pull the image
		image, err = client.Pull(ctx, DOCKER_HUB_PREFIX+imageName)
		if err != nil {
			return err
		}
	} else {
		image = ls[len(ls)-1]
	}
	// create the container with image spec
	ctn, err := client.NewContainer(ctx, "custom_container"+imageName,
		containerd.WithNewSnapshot(imageName+"-rootfs", image),
		containerd.WithNewSpec(oci.WithImageConfig(image)))
	if err != nil {
		return err
	}

	// create task
	task, err := ctn.NewTask(ctx, cio.NewCreator(cio.WithStdio))
	defer task.Delete(ctx)

	//pid := task.Pid()

	err = task.Start(ctx)
	if err != nil {
		return err
	}

	//status, err := task.Wait(ctx)

	time.Sleep(time.Second)

	return c.SendString("ok")
}

func new_client() (*containerd.Client, error) {
	return containerd.New("/run/containerd/containerd.sock", containerd.WithDefaultNamespace("ficus"))
}
