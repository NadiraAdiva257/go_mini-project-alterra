package controller

import (
	"mini-project/config"
	"mini-project/model"
	"mini-project/service"
	"mini-project/utils"
	"net/http"

	"mini-project/middleware"

	"github.com/labstack/echo/v4"
)

func CreateDebtorController(c echo.Context) error {
	debtor := model.Debtor{}
	c.Bind(&debtor)

	if err := service.GetDebtorRepository().CreateDebtorController(&debtor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new debtor",
	})
}

func LoginDebtorController(c echo.Context) error {
	var debtors model.Debtor
	var idDebtor int

	debtor := model.Debtor{}
	c.Bind(&debtor)

	cekEmail := config.DB.First(&debtors, "email = ?", debtor.Email)
	if err := cekEmail.Error; err != nil {
		return echo.ErrUnauthorized
	} else {
		cekPasswordError := utils.ComparePassword(debtors.Password, debtor.Password)
		if cekPasswordError != nil {
			return echo.ErrUnauthorized
		} else {
			idDebtor = int(debtors.ID)
		}
	}

	token, err := middleware.CreateToken(idDebtor, debtor.Email, debtor.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "fail login",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "login success",
		"your token": token,
	})
}

func UpdateDebtorController(c echo.Context) error {
	debtor_id := middleware.GetClaims(c).Id
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	debtor := model.Debtor{
		Name:     name,
		Email:    email,
		Password: hashPassword,
	}

	if err := service.GetDebtorRepository().UpdateDebtorController(&debtor, debtor_id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes update debtor by id",
	})
}
