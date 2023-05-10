package service

import (
	"mini-project/config"
	"mini-project/model"
)

type IDebtorService interface {
	CreateDebtorController(*model.Debtor) error
}

type DebtorRepository struct {
	Func IDebtorService
}

var debtorRepository IDebtorService

func init() {
	dr := &DebtorRepository{}
	dr.Func = dr

	debtorRepository = dr
}

func GetDebtorRepository() IDebtorService {
	return debtorRepository
}

func SetDebtorRepository(dr IDebtorService) {
	debtorRepository = dr
}

func (u *DebtorRepository) CreateDebtorController(debtor *model.Debtor) error {
	if err := config.DB.Save(&debtor); err != nil {
		return err.Error
	}

	return nil
}
