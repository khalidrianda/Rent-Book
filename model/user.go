package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id_user        uint       `gorm:"column:id_user;primaryKey;autoIncrement"`
	Nama_user      string     `gorm:"column:nama_user"`
	Email          string     `gorm:"column:email;unique"`
	Password       string     `gorm:"column:password"`
	Alamat         string     `gorm:"column:alamat"`
	Foto_profil    string     `gorm:"column:foto_profil"`
	Status_boolean bool       `gorm:"column:status;default:false"`
	Create_at      time.Time  `gorm:"created_at;autoCreateTime"`
	Updated_at     time.Time  `gorm:"column:updated_at;autoUpdateTime"`
	Bukus          []Buku     `gorm:"foreignKey:Id_user"`
	LendBooks      []LendBook `gorm:"foreignKey:Id_peminjam"`
}

type UserModel struct {
	DB *gorm.DB
}

func (um UserModel) GetAll(newData User) (User, error) {
	var res User
	err := um.DB.Where("email = ? && password = ?", newData.Email, newData.Password).Find(&res).Error

	if err != nil {
		fmt.Println("error on query", err.Error())
		return User{}, err
	}
	return res, nil
}

func (mm UserModel) Insert(newData User) (User, error) {
	err := mm.DB.Create(&newData).Error
	if err != nil {
		fmt.Println("error on insert", err.Error())
		return User{}, err
	}

	return newData, nil
}

func (um UserModel) Update(newData User) (User, error) {
	err := um.DB.Save(&newData).Error
	if err != nil {
		fmt.Println("error on insert", err.Error())
		return User{}, err
	}

	return newData, nil
}
