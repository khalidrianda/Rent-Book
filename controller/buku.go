package controller

import "Rent-Book/model"

type BukuControll struct {
	Model model.BukuModel
}

func (mc BukuControll) GetAll() ([]model.Buku, error) {
	res, err := mc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (mc BukuControll) GetMyBook(id uint) ([]model.Buku, error) {
	res, err := mc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (gc BukuControll) Add(data model.Buku) (model.Buku, error) {
	res, err := gc.Model.Insert(data)
	if err != nil {
		return model.Buku{}, err
	}
	return res, nil
}
