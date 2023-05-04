package model

import "gorm.io/gorm"

type DebtCategory struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Debts []Debt
}
