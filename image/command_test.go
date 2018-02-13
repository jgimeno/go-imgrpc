package image_test

import (
	"testing"
	"github.com/jgimeno/go-imgrpc/image/mocks"
	"io/ioutil"
	"bytes"
	"github.com/jgimeno/go-imgrpc/image"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
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

	cmd.SaveImage(img)
}

func instanceId(id image.Id) {}

func TestGetImageCommand(t *testing.T) {
	p := mocks.Persistence{}
	cmd := image.GetImageCommand{&p}

	fakeImage := &image.Image{}
	p.On("GetById", image.Id("theImageId")).Return(fakeImage)

	img := cmd.GetImage("theImageId")

	assert.Equal(t, fakeImage, img)
}