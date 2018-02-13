package boltdb

import (
	"testing"

	"io/ioutil"
	"github.com/jgimeno/go-imgrpc/image"
	"bytes"
)

func TestNew(t *testing.T) {
	p := New()

	imgBytes, err := ioutil.ReadFile("gopher.jpg")
	if err != nil {
		t.Fatalf("Failing to open file")
	}

	img := image.New("id", imgBytes)

	t.Run("We can save the image and retrieve later.", func(t *testing.T) {
		p.Save(img)
	})

	t.Run("We can get a saved image", func(t *testing.T) {
		img := p.GetById("id")
		if !bytes.Equal(imgBytes, img.Data) {
			t.Fatalf("Failed to get data.")
		}
	})
}
