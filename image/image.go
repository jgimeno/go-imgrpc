package image

import (
	"github.com/pkg/errors"
	"image/jpeg"
	"bytes"
	"image/png"
	image2 "image"
)

var (
	ErrUnsupportedType = errors.New("Unsupported type.")
	supportedTypes = map[string]bool{
		"jpg" : true,
		"png" : true,
	}
)

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

func (i *Image) TransformTo(fileType string) (*Image, error) {
	if !supportedTypes[fileType] {
		return nil, ErrUnsupportedType
	}

	if fileType == i.Type {
		return i, nil
	}

	var srcImg image2.Image
		srcImg, err := jpeg.Decode(bytes.NewReader(i.Data))
		if err != nil {
			return nil, err
		}

	var tData []byte
	b := bytes.NewBuffer(tData)
	png.Encode(b, srcImg)

	convertedImg := &Image{
		Type: "png",
		Data: b.Bytes(),
	}

	return convertedImg, nil
}
