package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type LendBook struct {
	Id_peminjaman int       `gorm:"column:id_peminjaman;primaryKey;autoIncrement"`
	Id_Pemilik    uint      `gorm:"column:id_pemilik"`
	Id_peminjam   uint      `gorm:"column:id_peminjam"`
	Id_buku       uint      `gorm:"column:id_buku"`
	Nama_buku     string    `gorm:"column:nama_buku"`
	Batas_waktu   time.Time `gorm:"column:batas_waktu"`
	Kembalikan    bool      `gorm:"column:kembalikan;default:false"`
	Create_at     time.Time `gorm:"created_at;autoCreateTime"`
	Return_at     time.Time `gorm:"column:return_at;autoUpdateTime"`
}

type LendBookModel struct {
	DB *gorm.DB
}

func (mm LendBookModel) GetAll(Id uint) ([]LendBook, error) {
	var res []LendBook
	err := mm.DB.Where("id_peminjam = ? && kembalikan=0", Id).Find(&res).Error
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

func (mm LendBookModel) Return(newData LendBook) (LendBook, error) {
	err := mm.DB.Select("kembalikan").Where("id_buku=?", newData.Id_buku).Updates(&newData).Error
	if err != nil {
		fmt.Println("error on insert", err.Error())
		return LendBook{}, err
	}
	return newData, nil
}

func (mm LendBookModel) CariPinjamUser(Id uint) (int64, int64) {
	var res, ser int64
	var lb LendBook
	mm.DB.Where("id_peminjam = ? && kembalikan=0", Id).Find(&lb).Count(&res)
	mm.DB.Where("id_pemilik = ? && kembalikan=0", Id).Find(&lb).Count(&ser)
	return res, ser
}
