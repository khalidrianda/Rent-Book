package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type LendBook struct {
	Id_peminjaman int
	Id_peminjam   int
	Id_buku       int
	Nama_buku     string
	Batas_waktu   time.Time
	Return_at     time.Time
}

type LendBookModel struct {
	DB *gorm.DB
}

func (mm LendBookModel) GetAll() ([]LendBook, error) {
	var res []LendBook
	err := mm.DB.Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (mm LendBookModel) Insert(newData LendBook) (LendBook, error) {
	err := mm.DB.Create(&newData).Error
	if err != nil {
		fmt.Println("error on insert", err.Error())
		return LendBook{}, err
	}
	return newData, nil
}

func (mm LendBookModel) Update(newData LendBook) (LendBook, error) {
	err := mm.DB.Save(&newData).Error
	if err != nil {
		fmt.Println("error on insert", err.Error())
		return LendBook{}, err
	}
	return newData, nil
}
