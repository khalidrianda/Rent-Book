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
	Create_at      time.Time  `gorm:"column:created_at;autoCreateTime"`
	Updated_at     time.Time  `gorm:"column:updated_at;autoUpdateTime"`
	Bukus          []Buku     `gorm:"foreignKey:Id_user"`
	LendBooks      []LendBook `gorm:"foreignKey:Id_peminjam"`
}

type UserModel struct {
	DB *gorm.DB
}

func (um UserModel) GetAll(data User) (User, error) {
	var res User
	err := um.DB.Where("email = ? && password = ? && status=1", data.Email, data.Password).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return User{}, err
	}
	return res, nil
}

func (um UserModel) LogIn(Id uint) ([]User, error) {
	var res []User
	err := um.DB.Where("id_user = ?", Id).Find(&res).Error
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

func (um UserModel) UpdateNama(newData User) {
	um.DB.Select("Nama_user").Where("Id_user = ?", newData.Id_user).Updates(&newData)
}

func (um UserModel) UpdateEmail(newData User) {
	um.DB.Select("Email").Where("Id_user = ?", newData.Id_user).Updates(&newData)
}

func (um UserModel) UpdatePassword(newData User) {
	um.DB.Select("Password").Where("Id_user = ?", newData.Id_user).Updates(&newData)
}

func (um UserModel) UpdateAlamat(newData User) {
	um.DB.Select("Alamat").Where("Id_User = ?", newData.Id_user).Updates(&newData)
}
func (um UserModel) UpdateFotoProfil(newData User) {
	um.DB.Select("Foto_profil").Where("Id_User = ?", newData.Id_user).Updates(&newData)
}
