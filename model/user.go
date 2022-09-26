package model

import "fmt"

type User struct {
	id_user        string
	nama_user      string
	email          string
	password       string
	alamat         string
	foto_profil    string
	status_boolean bool
}

func (um UserModel) GetAll() ([]User, error) {
	var res []User
	err := um.DB.Table("user").Select("id_user", "nama_user", "email", "password", "alamat", "foto_profil", "status_boolean").Model(&User{}).Find(&res).Error
	if err != nil {
		fmt.Println("error on query", err.Error())
		return nil, err
	}
	return res, nil
}

func (um UserModel) Insert(newData User) (User, error) {
	newData.id_user = "usr001"

	res, err := um.DB.Exec("INSERT INTO User (id_user, Nama_user, email, password, alamat, foto_profil, status_boolean) values (?,?,?,?,?,?,?) ",
		newData.id_user, newData.nama_user, newData.email, newData.password, newData.alamat, newData.foto_profil, newData.status_boolean)

	if err != nil {
		fmt.Println(err.Error())
		return User{}, err
	}
	affectted, err := res.RowsAffected()

	if affectted < 1 {
		return User{}, err
	}
	return newData, nil

}
