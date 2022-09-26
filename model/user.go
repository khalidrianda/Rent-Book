package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Id_user        int
	Nama_user      string
	Email          string
	Password       string
	Alamat         string
	Foto_profil    string
	Status_boolean bool
}

type UserModel struct {
	DB *gorm.DB
}

func (um UserModel) GetAll() ([]User, error) {
	var res []User
	err := um.DB.Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (um UserModel) Insert(newData User) (User, error) {
	err := um.DB.Create(&newData).Error

	if err != nil {
		fmt.Println(err.Error())
		return User{}, err
	}

	return newData, nil
}
