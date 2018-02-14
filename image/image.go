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
	var err error
	if i.Type == "jpg" {
		srcImg, err = jpeg.Decode(bytes.NewReader(i.Data))
		if err != nil {
			return nil, err
		}
	} else if i.Type == "png" {
		srcImg, err = png.Decode(bytes.NewReader(i.Data))
		if err != nil {
			return nil, err
		}
	}

	var tData []byte
	b := bytes.NewBuffer(tData)

	switch fileType {
	case "png":
		png.Encode(b, srcImg)
	case "jpg":
		jpeg.Encode(b, srcImg, nil)
	}

	convertedImg := &Image{
		Type: fileType,
		Data: b.Bytes(),
	}

	return convertedImg, nil
}
