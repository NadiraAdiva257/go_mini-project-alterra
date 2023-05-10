package service

import (
	"errors"
	"mini-project/model"

	"github.com/stretchr/testify/mock"
)

type DebtRepositoryMock struct {
	Mock mock.Mock
}

func (dm *DebtRepositoryMock) DeleteDebtController(id, debtor_id int) error {
	var debt []model.Debt
	var result error

	for _, val := range debt {
		if val.ID == id && val.DebtorID == debtor_id {
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
