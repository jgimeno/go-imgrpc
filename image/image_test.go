package image

import (
	"testing"
	"io/ioutil"
	"bytes"
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

	t.Run("It cannot converted to an unsupported type.", func(t *testing.T) {
		_, err := jpgImg.TransformTo("xls")
		if err != ErrUnsupportedType {
			t.Fatal("Error asserting unssupported type conversions.")
		}
	})

	t.Run("It can convert an jpg image to a png", func(t *testing.T) {
		tImage, err := jpgImg.TransformTo("png")
		if err != nil {
			t.Fatal("Error converting image.")
		}

		dataPngImage, err := ioutil.ReadFile("files/gopher.png")
		if err != nil {
			t.Fatal("Error opening png sample file.")
		}

		if !bytes.Equal(dataPngImage, tImage.Data) {
			t.Fatal("Error converting file to png.")
		}
	})
}
