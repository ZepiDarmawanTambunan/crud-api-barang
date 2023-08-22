package models

import (
	"time"

	"gorm.io/gorm"
)

type Barang struct {
	gorm.Model
	Nama      string    `gorm:"column:nama"`
	Harga     int       `gorm:"column:harga"`
	Deskripsi string    `gorm:"column:deskripsi"`
	Gambar    string    `gorm:"column:gambar"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (Barang) TableName() string {
	return "barang"
}
