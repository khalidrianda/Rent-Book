package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Buku struct {
	Id_buku     int    `gorm:"column:id_buku"`
	Code_buku   string `gorm:"column:code_buku"`
	Nama_buku   string `gorm:"column:nama_buku"`
	Pengarang   string `gorm:"column:pengarang"`
	Gambar_buku string `gorm:"column:gambar_buku"`
	Deskripsi   string `gorm:"column:deskripsi"`
	Is_lend     bool   `gorm:"column:is_lend"`
}

type BukuModel struct {
	DB *gorm.DB
}

func (mm BukuModel) GetAll() ([]Buku, error) {
	var res []Buku
	err := mm.DB.Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (mm BukuModel) Insert(newData Buku) (Buku, error) {
	err := mm.DB.Create(&newData).Error
	if err != nil {
		fmt.Println("error on insert", err.Error())
		return Buku{}, err
	}
	return newData, nil
}

func (mm BukuModel) Update(newData Buku) (Buku, error) {
	err := mm.DB.Save(&newData).Error
	if err != nil {
		fmt.Println("error on insert", err.Error())
		return Buku{}, err
	}
	return newData, nil
}
