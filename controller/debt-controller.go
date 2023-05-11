package controller

import (
	"mini-project/config"
	"mini-project/middleware"
	"mini-project/model"
	"mini-project/service"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// buat hutang
func CreateDebtController(c echo.Context) error {
	creditor_name := c.FormValue("creditor_name")

	formatDate := "2006-01-02"
	date, err := time.Parse(formatDate, c.FormValue("date"))
	if err != nil {
		return err
	}

	amount, err := strconv.Atoi(c.FormValue("amount"))
	if err != nil {
		return err
	}

	detail := c.FormValue("detail")

	debtor_id := middleware.GetClaims(c).Id

	debt := model.Debt{
		CreditorName: creditor_name,
		Date:         datatypes.Date(date),
		Amount:       amount,
		Detail:       detail,
		DebtorID:     debtor_id,
	}

	if err := service.GetDebtRepository().CreateDebtController(&debt); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new debt",
	})
}

// edit hutang
func UpdateDebtController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	creditor_name := c.FormValue("creditor_name")

	formatDate := "2006-01-02"
	date, err := time.Parse(formatDate, c.FormValue("date"))
	if err != nil {
		return err
	}

	amount, err := strconv.Atoi(c.FormValue("amount"))
	if err != nil {
		return err
	}

	detail := c.FormValue("detail")

	debtor_id := middleware.GetClaims(c).Id

	debt := model.Debt{
		CreditorName: creditor_name,
		Date:         datatypes.Date(date),
		Amount:       amount,
		Detail:       detail,
		DebtorID:     debtor_id,
	}

	if err := service.GetDebtRepository().UpdateDebtController(&debt, id, debtor_id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update debt by id",
	})
}

// hapus hutang
func DeleteDebtController(c echo.Context) error {
	debtor_id := middleware.GetClaims(c).Id

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := service.GetDebtRepository().DeleteDebtController(id, debtor_id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete debt by id",
	})
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

// lihat keseluruhan daftar hutang berdasarkan pengelompokan waktu
func GetAllDebtByTimeController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)

	var debts []model.Debt
	var debtByTime *gorm.DB
	var debtByTime2 *gorm.DB
	var resultErr error
	var result1 []Result1
	var resultTotal []ResultTotal

	timeDesc := TimeDesc(c)

	for _, value := range timeDesc {
		debtByTime = config.DB.Model(&debts).Where("date = ? AND debtor_id = ?", value, claims.Id).Find(&result1)
		debtByTime2 = config.DB.Model(&debts).Select("sum(amount) AS total").Where("date = ? AND debtor_id = ?", value, claims.Id).Find(&resultTotal)

		if err := debtByTime.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := debtByTime2.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		resultErr = c.JSON(http.StatusOK, map[string]interface{}{
			value:        result1,
			"total debt": resultTotal,
		})
	}

	return resultErr
}

func TimeDesc(c echo.Context) []string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)

	var debts []model.Debt
	var result []string

	dateDesc := config.DB.Order("date desc").Model(&debts).Distinct("date").Where("debtor_id = ?", claims.Id).Find(&result)

	if err := dateDesc.Error; err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return result
}

// melihat keseluruhan daftar hutang berdasarkan pengelompokan nama kreditur
func GetAllDebtByCreditorController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)

	var debts []model.Debt
	var debtByCreditor *gorm.DB
	var debtByCreditor2 *gorm.DB
	var resultErr error
	var result2 []Result2
	var resultTotal []ResultTotal

	creditorNameDesc := CreditorNameAsc(c)

	for _, value := range creditorNameDesc {
		debtByCreditor = config.DB.Model(&debts).Where("creditor_name = ? AND debtor_id = ?", value, claims.Id).Find(&result2)
		debtByCreditor2 = config.DB.Model(&debts).Select("sum(amount) AS total").Where("creditor_name = ? AND debtor_id = ?", value, claims.Id).Find(&resultTotal)

		if err := debtByCreditor.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := debtByCreditor2.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		resultErr = c.JSON(http.StatusOK, map[string]interface{}{
			value:        result2,
			"total debt": resultTotal,
		})
	}

	return resultErr
}

func CreditorNameAsc(c echo.Context) []string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)

	var debts []model.Debt
	var result []string

	creditorName := config.DB.Order("creditor_name asc").Model(&debts).Distinct("creditor_name").Where("debtor_id = ?", claims.Id).Find(&result)

	if err := creditorName.Error; err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return result
}

// cari daftar hutang berdasarkan waktu
func GetDebtByTimeController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)

	var debts []model.Debt
	var resultTotal []ResultTotal
	var result1 []Result1

	date := c.QueryParam("date")

	debtByTime := config.DB.Model(&debts).Select("sum(amount) AS total").Where("date = ? AND debtor_id = ?", date, claims.Id).Find(&resultTotal)
	debtByTime2 := config.DB.Model(&debts).Where("date = ? AND debtor_id = ?", date, claims.Id).Find(&result1)

	if err := debtByTime.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := debtByTime2.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"total debt": resultTotal,
		date:         result1,
	})
}

// cari daftar hutang berdasarkan nama kreditur
func GetDebtByCreditorController(c echo.Context) error {
	var debts []model.Debt
	var resultTotal []ResultTotal
	var result2 []Result2

	debtor_id := middleware.GetClaims(c).Id
	creditor := c.QueryParam("creditor_name")

	// result, err := service.GetDebtRepository().GetDebtByCreditorController(creditor, claims.Id)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": err.Error(),
	// 	})
	// } else {
	// 	return c.JSON(http.StatusOK, map[string]interface{}{
	// 		"total debt": result["total debt"],
	// 		creditor:     result[creditor],
	// 	})
	// }

	debtByCreditor := config.DB.Model(&debts).Select("sum(amount) AS total").Where("creditor_name = ? AND debtor_id = ?", creditor, debtor_id).Find(&resultTotal)
	debtByCreditor2 := config.DB.Model(&debts).Where("creditor_name = ? AND debtor_id = ?", creditor, debtor_id).Find(&result2)

	if err := debtByCreditor.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := debtByCreditor2.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"total debt": resultTotal,
		creditor:     result2,
	})
}

// lihat daftar hutang yang diurutkan dari hutang tertinggi berdasarkan pengelompokan nama kreditur
func GetAllDebtByTheHighest(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)

	var debts []model.Debt
	var debtByHighest *gorm.DB
	var resultErr error
	var resultTotal ResultTotal
	var result2 []Result2

	creditorNameHighest, creditorTotalHighest := DebtHighest(c)

	for i, value := range creditorNameHighest {
		resultTotal.Total = creditorTotalHighest[i]
		debtByHighest = config.DB.Model(&debts).Where("creditor_name = ? AND debtor_id = ?", value, claims.Id).Find(&result2)

		if err := debtByHighest.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		resultErr = c.JSON(http.StatusOK, map[string]interface{}{
			"total debt": resultTotal,
			value:        result2,
		})
	}

	return resultErr
}

func DebtHighest(c echo.Context) ([]string, []int) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)

	var debts []model.Debt
	var resultName []string
	var resultTotal []int

	creditorName := config.DB.Order("sum(amount) desc").Model(&debts).Select("creditor_name").Where("debtor_id = ?", claims.Id).Group("creditor_name").Find(&resultName)
	creditorTotal := config.DB.Order("sum(amount) desc").Model(&debts).Select("sum(amount)").Where("debtor_id = ?", claims.Id).Group("creditor_name").Find(&resultTotal)

	if err := creditorName.Error; err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := creditorTotal.Error; err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return resultName, resultTotal
}

// lihat daftar hutang yang diurutkan dari hutang terlama berdasarkan pengelompokan nama kreditur)
func GetAllDebtByTheLongest(c echo.Context) error {
	var debtor_id = middleware.GetClaims(c).Id
	var debts []model.Debt
	var creditorNameArray []string
	var creditorTotalArray []int

	var resultErr error
	var result2 []Result2

	creditorName := config.DB.Order("datediff(curdate(), min(date)) desc").Model(&debts).Select("creditor_name").Where("debtor_id = ?", debtor_id).Group("creditor_name").Find(&creditorNameArray)
	if err := creditorName.Error; err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	creditorTotal := config.DB.Order("datediff(curdate(), min(date)) desc").Model(&debts).Select("datediff(curdate(), min(date))").Where("debtor_id = ?", debtor_id).Group("creditor_name").Find(&creditorTotalArray)
	if err := creditorTotal.Error; err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for i, value := range creditorNameArray {
		debtByLongest := config.DB.Order("date asc").Model(&debts).Where("creditor_name = ? AND debtor_id = ?", value, debtor_id).Find(&result2)
		if err := debtByLongest.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		resultErr = c.JSON(http.StatusOK, map[string]interface{}{
			"total day": creditorTotalArray[i],
			value:       result2,
		})
	}

	return resultErr
}
