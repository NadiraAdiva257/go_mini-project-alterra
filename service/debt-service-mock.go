package service

import (
	"errors"
	"mini-project/model"
	"time"

	"github.com/stretchr/testify/mock"
	"gorm.io/datatypes"
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

func (dm *DebtRepositoryMock) GetDebtByCreditorController(creditor_name string, debtor_id int) (map[string]interface{}, error) {
	var debt []model.Debt
	var result = make(map[string]interface{})

	for _, val := range debt {
		if val.CreditorName == creditor_name && val.DebtorID == debtor_id {
			result["total debt"] = 23000
			result["mayla"] = []model.Debt{
				{
					Date:   datatypes.Date(time.Now()),
					Amount: 23000,
					Detail: "cilok",
				},
				{
					Date:   datatypes.Date(time.Now()),
					Amount: 23000,
					Detail: "cilok",
				},
			}
			break
		}
	}

	if len(result) != 0 {
		return result, nil
	} else {
		return nil, errors.New("data empty")
	}
}
