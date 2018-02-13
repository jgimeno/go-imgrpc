package image

type Id string

type Image struct {
	Id   Id
	Data []byte
	Type string
}

func New(id Id, data []byte, fileType string) *Image {
	return &Image{
		Id:   id,
		Data: data,
		Type: fileType,
	}
}
