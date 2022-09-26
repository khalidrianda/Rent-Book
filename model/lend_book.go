package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type LendBook struct {
	Id_peminjaman int `gorm:"column:id_peminjaman;primaryKey;autoIncrement"`
	Id_peminjam   int `gorm:"column:id_peminjam"`
	// User          User      `gorm:"references:id_user"`
	Id_buku int `gorm:"column:id_buku"`
	// Buku          Buku      `gorm:"references:id_buku"`
	Nama_buku   string    `gorm:"column:nama_buku"`
	Batas_waktu time.Time `gorm:"column:batas_waktu"`
	Create_at   time.Time `gorm:"created_at;autoCreateTime"`
	Return_at   time.Time `gorm:"column:return_at;autoUpdateTime"`
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
