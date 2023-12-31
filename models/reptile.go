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
	RegularPrice string `json:"regularPrice"`
	Price        string `json:"price"`
	AgeCategory  string `json:"age"`
	pictures     [][]byte
	Description  string `json:"description"`
	Genre        string `json:"genre"`
}
