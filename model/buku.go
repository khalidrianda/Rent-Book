package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Buku struct {
	Id_buku     int
	code_buku   string
	nama_buku   string
	pengarang   string
	gambar_buku string
	deskripsi   string
	is_lend     bool
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
