package models

type Barang struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Code      string `gorm:"unique;size:10;not null" json:"code"`
	Nama      string `gorm:"size:20;not null" json:"nama"`
	Jumlah    int    `gorm:"not null" json:"jumlah"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
	IsActive  bool   `gorm:"not null" json:"isActive"`
}
type Stok struct {
	Code   string `gorm:"unique;size:10;not null" json:"code"`
	Tipe   string `gorm:"type:text" json:"tipe"`
	Jumlah int    `gorm:"type:number" json:"jumlah"`
}