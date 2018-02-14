package image_test

import (
	"testing"
	"github.com/jgimeno/go-imgrpc/image/mocks"
	"io/ioutil"
	"bytes"
	"github.com/jgimeno/go-imgrpc/image"
	"github.com/stretchr/testify/mock"
)

func TestSaveImageCommand(t *testing.T) {
	p := mocks.Persistence{}

	cmd := image.SaveImageCommand{&p}
	img, _ := ioutil.ReadFile("files/gopher.jpg")

	p.On("Save", mock.MatchedBy(
		func(i *image.Image) bool {
			if !bytes.Equal(img, i.Data) {
				return false
			}

			instanceId(i.Id)

			return true
		},
	))

	cmd.SaveImage(img, "jpg")
}

func instanceId(id image.Id) {}

func TestGetImage(t *testing.T) {
	p := mocks.Persistence{}
	cmd := image.GetImageCommand{&p}

	baseImgData, err := ioutil.ReadFile("files/gopher.jpg")
	if err != nil {
		t.Fatal("Error loading image jpg")
	}

	convertedImageData, err := ioutil.ReadFile("files/gopher.png")
	if err != nil {
		t.Fatal("Error loading image png")
	}

	returnImg := &image.Image{
		Data: baseImgData,
		Type: "jpg",
	}

	t.Run("We get the image without conversion when it is nil", func(t *testing.T) {
		p.On("GetById", image.Id("theImageId")).Return(returnImg)
		img, err := cmd.GetImage("theImageId", nil)
		if err != nil {
			t.Fatal("Error getting the image.")
		}

		if !bytes.Equal(baseImgData, img.Data) {
			t.Fatal("Error asserting that image is not being converted.")
		}
	})

	t.Run("We get the image converted when we set a fileType", func(t *testing.T) {
		p.On("GetById", image.Id("theImageId")).Return(returnImg)
		img, err := cmd.GetImage("theImageId", "png")
		if err != nil {
			t.Fatal("Error getting the converted image.")
		}

		if !bytes.Equal(convertedImageData, img.Data) {
			t.Fatal("Error asserting that image is being converted.")
		}
	})
}