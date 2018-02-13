package image

type Persistence interface {
	Save(image *Image)
	GetById(id Id) *Image
}
