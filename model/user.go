package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	id_user        string
	nama_user      string
	email          string
	password       string
	alamat         string
	foto_profil    string
	status_boolean bool
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
