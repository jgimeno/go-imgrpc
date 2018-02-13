package image

type Id string

type Image struct {
	Id   Id
	Data []byte
}

func New(id Id, data []byte) *Image {
	return &Image{
		Id:   id,
		Data: data,
	}
}
