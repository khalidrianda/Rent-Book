package controller

import "Rent-Book/model"

type LendBookControl struct {
	Model model.LendBookModel
}

func (mc LendBookControl) GetAll(Id uint) ([]model.LendBook, error) {
	res, err := mc.Model.GetAll(Id)

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

func (gc LendBookControl) Return(data model.LendBook) (model.LendBook, error) {
	res, err := gc.Model.Return(data)
	if err != nil {
		return model.LendBook{}, err
	}
	return res, nil
}

func (mc LendBookControl) CariPinjamUser(Id uint) (int64, int64) {
	res, ser := mc.Model.CariPinjamUser(Id)

	return res, ser
}
