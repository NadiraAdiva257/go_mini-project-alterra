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
	arguments := drm.Mock.Called(debtorUpdate, id)
	if arguments.Get(0) == nil {
		return errors.New("error")
	} else {
		return nil
	}
}
