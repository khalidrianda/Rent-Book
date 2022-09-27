package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id_user        int       `gorm:"column:id_user;primaryKey"`
	Nama_user      string    `gorm:"column:nama_user"`
	Email          string    `gorm:"column:email;unique"`
	Password       string    `gorm:"column:password"`
	Alamat         string    `gorm:"column:alamat"`
	Foto_profil    string    `gorm:"column:foto_profil"`
	Status_boolean bool      `gorm:"column:status;default:false"`
	Create_at      time.Time `gorm:"created_at;autoCreateTime"`
	Updated_at     time.Time `gorm:"column:updated_at;autoUpdateTime"`
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
