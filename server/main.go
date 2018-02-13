package main

import (
	"github.com/jgimeno/go-imgrpc/protocolbuffer"
	"golang.org/x/net/context"
	"flag"
	"net"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

type ImageService struct {
}

func (*ImageService) SaveImage(context.Context, *protocolbuffer.Image) (*protocolbuffer.ImageId, error) {
	panic("implement me")
}

func (*ImageService) GetImage(context.Context, *protocolbuffer.ImageId) (*protocolbuffer.Image, error) {
	panic("implement me")
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen to :8080")
	}

	grpcServer := grpc.NewServer()
	protocolbuffer.RegisterImageServiceServer(grpcServer, &ImageService{})
	grpcServer.Serve(lis)
}