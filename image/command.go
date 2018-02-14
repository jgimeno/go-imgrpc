package image

import (
	"github.com/satori/go.uuid"
)

type SaveImageCommand struct {
	P Persistence
}

func (c *SaveImageCommand) SaveImage(image []byte, fileType string) Id {
	var id = Id(uuid.NewV1().String())
	i := New(id, image, fileType)
	c.P.Save(i)

	return id
}

type GetImageCommand struct {
	P Persistence
}

func (c *GetImageCommand) GetImage(imageId string) *Image {
	id := Id(imageId)
	image := c.P.GetById(id)
	return image
}

func (c *GetImageCommand) GetImageWithType(imageId string, fileType string) (*Image, error) {
	id := Id(imageId)
	image := c.P.GetById(id)

	tImg, err := image.TransformTo(fileType)
	if err != nil {
		return nil, err
	}

	return tImg, nil
}

