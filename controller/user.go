package controller

import (
	"Rent-Book/model"
)

type UserControll struct {
	Model model.UserModel
}

func (uc UserControll) GetAll(data model.User) (model.User, error) {
	res, err := uc.Model.GetAll(data)

	if err != nil {
		return model.User{}, err
	}
	return res, nil
}
func (uc UserControll) LogIn(Id uint) ([]model.User, error) {
	res, err := uc.Model.LogIn(Id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc UserControll) Add(data model.User) (model.User, error) {
	res, err := uc.Model.Insert(data)
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}

func (uc UserControll) Update(data model.User) (model.User, error) {
	res, err := uc.Model.Update(data)
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}

func (uc UserControll) UpdateId(data model.User) {
	uc.Model.UpdateId(data)
}

func (uc UserControll) UpdateNama(data model.User) {
	uc.Model.UpdateNama(data)
}

func (uc UserControll) UpdateEmail(data model.User) {
	uc.Model.UpdateEmail(data)
}

func (uc UserControll) UpdatePassword(data model.User) {
	uc.Model.UpdatePassword(data)
}

func (uc UserControll) UpdateAlamat(data model.User) {
	uc.Model.UpdateAlamat(data)
}
func (uc UserControll) UpdateFotoProfil(data model.User) {
	uc.Model.UpdateFotoProfil(data)
}
