package main

import (
	"testing"
	"net"
	"google.golang.org/grpc"
	"github.com/jgimeno/go-imgrpc/protocolbuffer"
	"os"
	"context"
	"io/ioutil"
	"github.com/jgimeno/go-imgrpc/persistence/boltdb"
	"bytes"
)

const (
	port = ":8081"
)

func Server() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic("Error when listening to port.")
	}

	p := boltdb.New()
	s := grpc.NewServer()
	protocolbuffer.RegisterImageServiceServer(s, &ImageService{p})
	if err := s.Serve(lis); err != nil {
		panic("Failed to serve the server.")
	}
}

func TestMain(m *testing.M) {
	go Server()
	os.Exit(m.Run())
}

func TestMicroservice(t *testing.T) {
	const address = ":8081"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatal("Error connecting to microservice.")
	}
	defer conn.Close()

	c := protocolbuffer.NewImageServiceClient(conn)

	var savedImageId *protocolbuffer.ImageId

	initImage, err := ioutil.ReadFile("files/gopher.jpg")
	if err != nil {
		t.Fatal("Error opening jpg image.")
	}

	t.Run("We can upload an image and get an id", func(t *testing.T) {
		pImg := &protocolbuffer.Image{
			Data: initImage,
		}

		savedImageId, err = c.SaveImage(context.Background(), pImg)
		if err != nil {
			t.Fatal("Failed to call microservice to save image. ", err.Error())
		}

		if savedImageId == nil {
			t.Fatal("Failed when trying to save image.")
		}
	})

	t.Run("We can get the image by the id", func(t *testing.T) {
		img, err := c.GetImage(context.Background(), savedImageId)
		if err != nil {
			t.Fatal("Failed to get the image from the microservice.")
		}

		if !bytes.Equal(initImage, img.Data) {
			t.Fatal("Failed asserting that the image coming from the microservice is the expected one.")
		}
	})

	os.Remove("my.db")
}