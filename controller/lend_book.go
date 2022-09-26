package controller

import "Rent-Book/model"

type LendBookControl struct {
	Model model.LendBookModel
}

func (mc LendBookControl) GetAll() ([]model.LendBook, error) {
	res, err := mc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (gc LendBookControl) Add(data model.LendBook) (model.LendBook, error) {
	res, err := gc.Model.Insert(data)
	if err != nil {
		return model.LendBook{}, err
	}
	return res, nil
}
