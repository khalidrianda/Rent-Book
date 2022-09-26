package controller

import (
	"rent-book/model"
)

type UserControll struct {
	Model model.UserModel
}

func (uc UserControll) GetAll() ([]model.User, error) {
	res, err := uc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc UserControll) add(data model.User) (model.User, error) {
	res, err := uc.Model.Insert(data)

	if err != nil {
		return model.User{}, err
	}
	return res, nil
}
