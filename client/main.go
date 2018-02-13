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
)

var (
	ErrMissingFileName = errors.New("Missing image filename.")
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name: "save",
			Aliases: []string{"s"},
			Usage: "Save an image in the microservice",
			Action: func(c *cli.Context) error {
				fileName := c.Args().First()

				if len(c.Args()) < 1 {
					return ErrMissingFileName
				}

				conn, err := grpc.Dial(":8080", grpc.WithInsecure())
				if err != nil {
					return ErrMissingFileName
				}

				defer conn.Close()
				client := protocolbuffer.NewImageServiceClient(conn)

				imgBytes, _ := ioutil.ReadFile(fileName)

				img := &protocolbuffer.Image{
					Data: imgBytes,
				}

				id, _ := client.SaveImage(context.Background(), img)
				fmt.Printf(id.Id)
				return nil
			},
		},
		{
			Name: "get",
			Aliases: []string{"g"},
			Usage: "Use to get the files by id.",
			Action: func(c *cli.Context) error {
				fmt.Print("Getting the file.")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}
}
