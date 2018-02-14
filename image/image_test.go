package image

import (
	"testing"
	"io/ioutil"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestAnImageCanBeConvertedToOtherFileTypes(t *testing.T) {
	data, err := ioutil.ReadFile("files/gopher.jpg")
	if err != nil {
		t.Fatal("Error opening image jpg.")
	}

	jpgImg := Image{
		Data: data,
		Type: "jpg",
	}

	data, err = ioutil.ReadFile("files/beatles.png")
	if err != nil {
		t.Fatal("Error opening png file.")
	}

	pngImg := Image{
		Data:data,
		Type:"png",
	}

	t.Run("It cannot converted to an unsupported type.", func(t *testing.T) {
		_, err := jpgImg.TransformTo("xls")
		if err != ErrUnsupportedType {
			t.Fatal("Error asserting unssupported type conversions.")
		}
	})

	t.Run("It can convert an jpg image to a png", func(t *testing.T) {
		tImage, err := jpgImg.TransformTo("png")
		if err != nil {
			t.Fatal("Error transforming image.")
		}

		dataPngImage, err := ioutil.ReadFile("files/gopher.png")
		if err != nil {
			t.Fatal("Error opening png sample file.")
		}

		if !bytes.Equal(dataPngImage, tImage.Data) {
			t.Fatal("Error converting file to png.")
		}

		assert.Equal(t, "png", tImage.Type)
	})

	t.Run("It can convert a png image to jpg.", func(t *testing.T) {
		tImage, err := pngImg.TransformTo("jpg")
		if err != nil {
			t.Fatal("Error transforming image.")
		}

		dataJpgImage, err := ioutil.ReadFile("files/beatles.jpg")
		if !bytes.Equal(dataJpgImage, tImage.Data) {
			t.Fatal("Error converting file to jpg.")
		}

		assert.Equal(t, "jpg", tImage.Type)
	})
}
