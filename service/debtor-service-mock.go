package service

import (
	"errors"
	"mini-project/model"

	"github.com/stretchr/testify/mock"
)

type DebtorRepositoryMock struct {
	Mock mock.Mock
}

func (d *DebtorRepositoryMock) CreateDebtorController(debtor *model.Debtor) error {
	if debtor == nil {
		return errors.New("error")
	} else {
		return nil
	}
}

func (drm *DebtorRepositoryMock) UpdateDebtorController(debtorUpdate *model.Debtor, id int) error {
	debtor := []model.Debtor{
		{
			ID:       1,
			Name:     "nadira",
			Email:    "nadira123@gmail.com",
			Password: "nadira123",
		},
		{
			ID:       2,
			Name:     "adiva",
			Email:    "adiva123@gmail.com",
			Password: "adiva123",
		},
	}

	var result error

	for _, val := range debtor {
		if val.ID == id {
			result = nil
			break
		}
	}

	if result == nil {
		return result
	} else {
		return errors.New("error")
	}
}
