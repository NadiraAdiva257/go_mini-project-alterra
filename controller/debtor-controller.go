package controller

import (
	"mini-project/config"
	"mini-project/model"
	"mini-project/utils"
	"net/http"

	"mini-project/middleware"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func CreateDebtorController(c echo.Context) error {
	debtor := model.Debtor{}
	c.Bind(&debtor)

	if err := config.DB.Save(&debtor).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	var debtors []model.Debtor

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	debtorById := config.DB.Model(&debtors).Where("id = ?", claims.Id).Updates(model.Debtor{Name: name, Email: email, Password: hashPassword})

	if err := debtorById.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes update debtor by id",
	})
}
