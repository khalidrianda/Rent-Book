package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Buku struct {
	Id_buku     int        `gorm:"column:id_buku;primaryKey;autoIncrement"`
	Id_user     uint       `gorm:"column:id_user"`
	Code_buku   string     `gorm:"column:code_buku"`
	Nama_buku   string     `gorm:"column:nama_buku"`
	Pengarang   string     `gorm:"column:pengarang"`
	Gambar_buku string     `gorm:"column:gambar_buku"`
	Deskripsi   string     `gorm:"column:deskripsi"`
	Is_lend     bool       `gorm:"column:is_lend"`
	LendBooks   []LendBook `gorm:"foreignKey:Id_buku"`
}

type BukuModel struct {
	DB *gorm.DB
}

func (mm BukuModel) GetAll(session uint) ([]Buku, error) {
	var res []Buku
	err := mm.DB.Where("is_lend = 0 && id_user != ?", session).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (mm BukuModel) GetMyBook(id uint) ([]Buku, error) {
	var res []Buku
	err := mm.DB.Where("id_user = ?", id).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (mm BukuModel) GetName(id uint) (Buku, error) {
	var res Buku
	err := mm.DB.Select("nama_buku").Where("id_buku = ?", id).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return Buku{}, err
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
	err := mm.DB.Where("id_buku = ?", newData.Id_buku).Updates(&newData).Error
	if err != nil {
		fmt.Println("error on insert", err.Error())
		return Buku{}, err
	}
	return newData, nil
}

func (mm BukuModel) Delete(newData Buku) (Buku, error) {
	err := mm.DB.Where("id_buku = ?", newData.Id_buku).Delete(&newData).Error
	if err != nil {
		fmt.Println("error on insert", err.Error())
		return Buku{}, err
	}
	return newData, nil
}

func (mm BukuModel) Dipinjam(newData Buku) {
	mm.DB.Select("is_lend").Where("id_buku = ?", newData.Id_buku).Updates(&newData)
}

func (mm BukuModel) UpdateCode(newData Buku) {
	mm.DB.Select("code_buku").Where("id_buku = ?", newData.Id_buku).Updates(&newData)
}

func (mm BukuModel) UpdateNama(newData Buku) {
	mm.DB.Select("nama_buku").Where("id_buku = ?", newData.Id_buku).Updates(&newData)
}

func (mm BukuModel) UpdatePengarang(newData Buku) {
	mm.DB.Select("pengarang").Where("id_buku = ?", newData.Id_buku).Updates(&newData)
}

func (mm BukuModel) UpdateGambar(newData Buku) {
	mm.DB.Select("gambar_buku").Where("id_buku = ?", newData.Id_buku).Updates(&newData)
}

func (mm BukuModel) UpdateDeskripsi(newData Buku) {
	mm.DB.Select("deskripsi").Where("id_buku = ?", newData.Id_buku).Updates(&newData)
}
