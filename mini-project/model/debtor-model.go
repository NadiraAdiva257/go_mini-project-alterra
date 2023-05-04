package model

import (
	"mini-project/utils"

	"gorm.io/gorm"
)

type Debtor struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Debts    []Debt
}

func (d *Debtor) BeforeCreate(tx *gorm.DB) error {
	hashPassword, err := utils.HashPassword(d.Password)
	if err != nil {
		return err
	}

	d.Password = hashPassword

	return err
}
