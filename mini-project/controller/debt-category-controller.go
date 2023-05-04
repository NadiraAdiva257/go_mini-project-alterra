package controller

import (
	"mini-project/config"
	"mini-project/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateDebtCategoryController(c echo.Context) error {
	debtCategory := model.DebtCategory{}
	c.Bind(&debtCategory)

	if err := config.DB.Save(&debtCategory).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success create new debt category",
		"debt category": debtCategory,
	})
}
