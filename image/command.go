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

func (c *GetImageCommand) GetImage(imageId string, fileType interface{}) (*Image, error) {
	var image *Image
	var err error

	id := Id(imageId)
	image = c.P.GetById(id)

	if fileType != nil {
		image, err = image.TransformTo(fileType.(string))
		if err != nil {
			return nil, err
		}
	}

	return image, nil
}

