package image

import "github.com/satori/go.uuid"

type SaveImageCommand struct {
	P Persistence
}

func (c *SaveImageCommand) SaveImage(image []byte) Id {
	var id = Id(uuid.NewV1().String())
	i := New(id, image)
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
