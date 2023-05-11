package service

import (
	"log"
	"mini-project/config"
	"mini-project/model"

	"gorm.io/datatypes"
)

type IDebtService interface {
	CreateDebtController(debt *model.Debt) error
	UpdateDebtController(debtUpdate *model.Debt, id int, debtor_id int) error
	DeleteDebtController(id, debtor_id int) error
	GetDebtByCreditorController(creditor_name string, debtor_id int) (map[string]interface{}, error)
}

type DebtRepository struct {
	Func IDebtService
}

var debtRepository IDebtService

func init() {
	dr := &DebtRepository{}
	dr.Func = dr

	debtRepository = dr
}

func GetDebtRepository() IDebtService {
	return debtRepository
}

func SetDebtRepository(dr IDebtService) {
	debtRepository = dr
}

func (d *DebtRepository) CreateDebtController(debt *model.Debt) error {
	if err := config.DB.Save(&debt); err != nil {
		return err.Error
	}

	return nil
}

func (d *DebtRepository) UpdateDebtController(debtUpdate *model.Debt, id int, debtor_id int) error {
	var debt model.Debt

	if err := config.DB.Model(&debt).Where("id = ? AND debtor_id = ?", id, debtor_id).Updates(debtUpdate); err != nil {
		return err.Error
	}

	return nil
}

func (d *DebtRepository) DeleteDebtController(id, debtor_id int) error {
	var debt model.Debt

	if err := config.DB.Where("id = ? AND debtor_id = ?", id, debtor_id).Delete(&debt); err != nil {
		return err.Error
	}

	return nil
}

type Result1 struct {
	CreditorName string
	Amount       int
	Detail       string
}

type Result2 struct {
	Date   datatypes.Date
	Amount int
	Detail string
}

type ResultTotal struct {
	Total int
}

func (d *DebtRepository) GetDebtByCreditorController(creditor_name string, debtor_id int) (map[string]interface{}, error) {
	var debt model.Debt
	var resultTotal []ResultTotal
	var result2 []Result2

	if err := config.DB.Model(&debt).Select("sum(amount) AS total").Where("creditor_name = ? AND debtor_id = ?", creditor_name, debtor_id).Find(&resultTotal); err != nil {
		return nil, err.Error
	}

	if err := config.DB.Model(&debt).Where("creditor_name = ? AND debtor_id = ?", creditor_name, debtor_id).Find(&result2); err != nil {
		return nil, err.Error
	}

	log.Println(result2)

	return map[string]interface{}{
		"total debt":  resultTotal,
		creditor_name: result2,
	}, nil
}
