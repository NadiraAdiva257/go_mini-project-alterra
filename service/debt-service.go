package service

import (
	"mini-project/config"
	"mini-project/model"
)

type IDebtService interface {
	DeleteDebtController(id, debtor_id int) error
}

type DebtRepository struct {
	Func IDebtService
}

var debtRepository IDebtService

func init() {
	dr := &DebtorRepository{}
	dr.Func = dr

	debtRepository = dr
}

func GetDebtRepository() IDebtService {
	return debtRepository
}

func SetDebtRepository(dr IDebtService) {
	debtRepository = dr
}

func (d *DebtorRepository) DeleteDebtController(id, debtor_id int) error {
	var debt model.Debt

	if err := config.DB.Where("id = ? AND debtor_id = ?", id, debtor_id).Delete(&debt); err != nil {
		return err.Error
	}

	return nil
}
