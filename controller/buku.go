package controller

import "Rent-Book/model"

type BukuControll struct {
	Model model.BukuModel
}

func (mc BukuControll) GetAll(session uint) ([]model.Buku, error) {
	res, err := mc.Model.GetAll(session)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (mc BukuControll) GetMyBook(id uint) ([]model.Buku, error) {
	res, err := mc.Model.GetMyBook(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (mc BukuControll) GetName(id uint) (model.Buku, error) {
	res, err := mc.Model.GetName(id)
	if err != nil {
		return model.Buku{}, err
	}
	return res, nil
}

func (mc BukuControll) CariBuku(namaBuku string) ([]model.Buku, error) {
	res, err := mc.Model.CariBuku(namaBuku)
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

func (gc BukuControll) Update(data model.Buku) (model.Buku, error) {
	res, err := gc.Model.Update(data)
	if err != nil {
		return model.Buku{}, err
	}
	return res, nil
}

func (gc BukuControll) Delete(data model.Buku) (model.Buku, error) {
	res, err := gc.Model.Delete(data)
	if err != nil {
		return model.Buku{}, err
	}
	return res, nil
}
func (gc BukuControll) DeleteBukuUser(id uint) {
	gc.Model.DeleteBukuUser(id)

}

func (gc BukuControll) Dikembalikan(data model.Buku) {
	gc.Model.Dikembalikan(data)

}

func (gc BukuControll) Dipinjam(data model.Buku) {
	gc.Model.Dipinjam(data)
}

func (gc BukuControll) UpdateCode(data model.Buku) {
	gc.Model.UpdateCode(data)
}

func (gc BukuControll) UpdateNama(data model.Buku) {
	gc.Model.UpdateNama(data)
}

func (gc BukuControll) UpdatePengarang(data model.Buku) {
	gc.Model.UpdatePengarang(data)
}

func (gc BukuControll) UpdateGambar(data model.Buku) {
	gc.Model.UpdateGambar(data)
}

func (gc BukuControll) UpdateDeskripsi(data model.Buku) {
	gc.Model.UpdateDeskripsi(data)
}
