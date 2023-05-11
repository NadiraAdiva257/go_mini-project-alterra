package service

import (
	"errors"
	"mini-project/model"

	"github.com/stretchr/testify/mock"
)

type DebtorRepositoryMock struct {
	Mock mock.Mock
}

func (drm *DebtorRepositoryMock) CreateDebtorController(debtor *model.Debtor) error {
	if debtor == nil {
		return errors.New("error")
	} else {
		return nil
	}
}

func (drm *DebtorRepositoryMock) UpdateDebtorController(debtorUpdate *model.Debtor, id int) error {
	var debtor []model.Debt
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
