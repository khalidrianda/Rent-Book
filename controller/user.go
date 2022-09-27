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

func (uc UserControll) Add(data model.User) (model.User, error) {
	res, err := uc.Model.Insert(data)

	if err != nil {
		return model.User{}, err
	}
	return res, nil
}
