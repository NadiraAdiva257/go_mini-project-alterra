package model

import (
	"gorm.io/gorm"

	"gorm.io/datatypes"
)

type Debt struct {
	gorm.Model
	ID           int            `json:"id" form:"id"`
	CreditorName string         `json:"creditor_name" form:"creditor_name"`
	Date         datatypes.Date `json:"date" form:"date"`
	Amount       int            `json:"amount" form:"amount"`
	Detail       string         `json:"detail" form:"detail"`
	DebtorID     int            `json:"debtor_id" form:"debtor_id"`
}
