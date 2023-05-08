package route

import (
	"mini-project/controller"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func New() *echo.Echo {
	e := echo.New()

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(controller.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	groupDebts := e.Group("/debts")
	groupDebts.Use(echojwt.WithConfig(config))

	groupDebts.POST("", controller.CreateDebtController)
	groupDebts.PUT("/:id", controller.UpdateDebtController)
	groupDebts.DELETE("/:id", controller.DeleteDebtController)

	groupDebts.GET("/alltime", controller.GetAllDebtByTimeController)
	groupDebts.GET("/allcreditor", controller.GetAllDebtByCreditorController)
	groupDebts.GET("/time", controller.GetDebtByTimeController)
	groupDebts.GET("/creditor", controller.GetDebtByCreditorController)
	groupDebts.GET("/thehighest", controller.GetAllDebtByTheHighest)
	groupDebts.GET("/thelongest", controller.GetAllDebtByTheLongest)

	groupDebtors := e.Group("/debtors")
	groupDebtors.Use(echojwt.WithConfig(config))

	e.POST("/debtors", controller.CreateDebtorController)
	e.POST("/debtors/login", controller.LoginDebtorController)
	groupDebtors.PUT("", controller.UpdateDebtorController)

	return e
}
