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
