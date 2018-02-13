package main

import (
	"google.golang.org/grpc"
	"github.com/jgimeno/go-imgrpc/protocolbuffer"
	"context"
	"io/ioutil"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic("Error connecting to server." + err.Error())
	}

	defer conn.Close()
	client := protocolbuffer.NewImageServiceClient(conn)

	imgBytes, _ := ioutil.ReadFile("gopher.jpg")

	img := &protocolbuffer.Image{
		Data: imgBytes,
	}

	client.SaveImage(context.Background(), img)
}
