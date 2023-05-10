package service

import (
	"mini-project/config"
	"mini-project/model"
)

type IDebtorService interface {
	CreateDebtorController(debtor *model.Debtor) error
	UpdateDebtorController(debtorUpdate *model.Debtor, id int) error
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

func (u *DebtorRepository) UpdateDebtorController(debtorUpdate *model.Debtor, id int) error {
	var debtor model.Debtor

	if err := config.DB.Model(&debtor).Where("id = ?", id).Updates(debtorUpdate); err != nil {
		return err.Error
	}

	return nil
}
