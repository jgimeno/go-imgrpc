package main

import (
	"github.com/jgimeno/go-imgrpc/protocolbuffer"
	"golang.org/x/net/context"
	"net"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"github.com/jgimeno/go-imgrpc/persistence/boltdb"
	"github.com/jgimeno/go-imgrpc/image"
)

type ImageService struct {
	p image.Persistence
}

func (is *ImageService) SaveImage(c context.Context, pImg *protocolbuffer.Image) (*protocolbuffer.ImageId, error) {
	cmd := image.SaveImageCommand{P: is.p}
	id := cmd.SaveImage(pImg.Data, pImg.Type)

	return &protocolbuffer.ImageId{Id:string(id)}, nil
}

func (is *ImageService) GetImage(c context.Context, pImg *protocolbuffer.ImageId) (*protocolbuffer.Image, error) {
	cmd := image.GetImageCommand{P: is.p}

	img, err := cmd.GetImage(pImg.Id, pImg.Type)
	if err != nil {
		return nil, err
	}

	return &protocolbuffer.Image{
		Id: string(img.Id),
		Data: img.Data,
		Type: img.Type,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen to :8080")
	}

	p := boltdb.New()

	grpcServer := grpc.NewServer()
	protocolbuffer.RegisterImageServiceServer(grpcServer, &ImageService{p})
	grpcServer.Serve(lis)
}