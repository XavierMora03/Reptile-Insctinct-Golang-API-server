package models

const (
	MACHO  = "M"
	HEMBRA = "H"
)

const (
	Cria    = "Cria"
	Juvenil = "Juvenil"
	Adulta  = "Adulta"
)

// const MAX_PHOTOS_PER_REPTILE_PRODUCT int = 10
type Reptile struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	RegularPrice uint64 `json:"regularPrice"`
	Price        uint64 `json:"price"`
	AgeCategory  string `json:"age"`
	pictures     *ReptilePictures
	Description  string `json:"description"`
	Genre        string `json:"genre"`
}

type ReptilePictures struct {
}
