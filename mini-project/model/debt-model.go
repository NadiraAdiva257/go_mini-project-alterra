package model

import (
	"gorm.io/gorm"

	"gorm.io/datatypes"
)

type Debt struct {
	gorm.Model
	CreditorName   string         `json:"creditor_name" form:"creditor_name"`
	Date           datatypes.Date `json:"date" form:"date"`
	Amount         int            `json:"amount" form:"amount"`
	DebtCategoryID int            `json:"debt_category_id" form:"debt_category_id"`
	Detail         string         `json:"detail" form:"detail"`
	DebtorID       int            `json:"debtor_id" form:"debtor_id"`
}
