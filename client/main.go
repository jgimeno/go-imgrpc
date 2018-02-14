package main

import (
	"google.golang.org/grpc"
	"github.com/jgimeno/go-imgrpc/protocolbuffer"
	"context"
	"io/ioutil"
	"github.com/urfave/cli"
	"os"
	"github.com/pkg/errors"
	"fmt"
	"path"
)

var (
	ErrMissingFileName = errors.New("Missing image filename.")
	ErrMissingImageId  = errors.New("Missing image id.")
	ErrConnectingMS    = errors.New("Error connecting to Microservice.")
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "save",
			Aliases: []string{"s"},
			Usage:   "Save an image in the microservice",
			Action: func(c *cli.Context) error {
				fileName := c.Args().First()

				if len(c.Args()) < 1 {
					return ErrMissingFileName
				}

				conn, err := connectMicroservice()
				if err != nil {
					return err
				}
				defer conn.Close()

				client := protocolbuffer.NewImageServiceClient(conn)

				imgBytes, _ := ioutil.ReadFile(fileName)

				img := &protocolbuffer.Image{
					Data: imgBytes,
					Type: path.Ext(fileName)[1:],
				}

				id, _ := client.SaveImage(context.Background(), img)
				fmt.Printf(id.Id)

				return nil
			},
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Use to get the files by id.",
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 1 {
					return ErrMissingImageId
				}

				conn, err := connectMicroservice()
				if err != nil {
					return ErrConnectingMS
				}
				defer conn.Close()

				client := protocolbuffer.NewImageServiceClient(conn)

				id := &protocolbuffer.ImageId{Id: c.Args().First()}

				var fileType string
				if fileType = c.Args().Get(1); fileType != "" {
					id.Type = fileType
				}

				img, err := client.GetImage(context.Background(), id)
				if err != nil {
					return errors.New("Error getting the image." + err.Error())
				}

				ioutil.WriteFile("file." + img.Type, img.Data, 0644)

				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}
}

func connectMicroservice() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		return nil, ErrConnectingMS
	}

	return conn, nil
}
