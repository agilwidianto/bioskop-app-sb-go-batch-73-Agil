package models

type Bioskop struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float64 `json:"rating"`
}

var bioskops = []Bioskop{
	{ID: 1, Nama: "CGV Transmart", Lokasi: "Jakarta", Rating: 4.5},
	{ID: 2, Nama: "Ciwalk XXI", Lokasi: "Bandung", Rating: 4.7},
	{ID: 3, Nama: "Botani Square XXI", Lokasi: "Bogor", Rating: 4.6},
}