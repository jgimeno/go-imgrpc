package boltdb

import (
	"github.com/jgimeno/go-imgrpc/image"
	"github.com/coreos/bbolt"
	"bytes"
	"encoding/gob"
)

const bucketName = "images"

func New() image.Persistence {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		panic("It was not possible to open DB.")
	}

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			_, err := tx.CreateBucket([]byte(bucketName))
			if err != nil {
				panic("It was not possible to create the db bucket.")
			}
		}

		return nil
	})

	return &boltDB{db}
}

type boltDB struct {
	db *bolt.DB
}

func (b *boltDB) Save(image *image.Image) {
	b.db.Update(func(tx *bolt.Tx) error {
		var buf bytes.Buffer
		e := gob.NewEncoder(&buf)

		err := e.Encode(image)
		if err != nil {
			panic("Error encoding the object to save on db.")
		}

		b := tx.Bucket([]byte(bucketName))
		b.Put([]byte(image.Id), buf.Bytes())

		return nil
	})
}

func (b *boltDB) GetById(id image.Id) *image.Image {
	var img *image.Image

	b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		result := b.Get([]byte(id))

		reader := bytes.NewReader(result)
		decoder := gob.NewDecoder(reader)
		decoder.Decode(img)

		return nil
	})

	return img
}

func (b *boltDB) Close() {
	b.db.Close()
}