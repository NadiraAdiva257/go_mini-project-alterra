package controller

import (
	"mini-project/config"
	"mini-project/model"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type JwtCustomClaims struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// buat debt
func CreateDebtController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

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

	debt := model.Debt{
		CreditorName: creditor_name,
		Date:         datatypes.Date(date),
		Amount:       amount,
		Detail:       detail,
		DebtorID:     claims.Id,
	}

	if err := config.DB.Save(&debt).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success buat new debt",
	})
}

// edit debt
func UpdateDebtController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	var debts []model.Debt

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

	debtById := config.DB.Model(&debts).Where("id = ? AND debtor_id = ?", id, claims.Id).Updates(model.Debt{
		CreditorName: creditor_name, Date: datatypes.Date(date), Amount: amount, Detail: detail})

	if err := debtById.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update debt by id",
	})
}

// hapus debt
func DeleteDebtController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	var debts []model.Debt

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	debtById := config.DB.Where("id = ? AND debtor_id = ?", id, claims.Id).Delete(&debts)

	if err := debtById.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete debt by id",
	})
}

// lihat semua debts berdasarkan waktu
func GetAllDebtByTimeController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	type Result struct {
		CreditorName string
		Date         datatypes.Date
		Amount       int
		Detail       string
	}

	type Result2 struct {
		Total int
	}

	var debts []model.Debt
	var debtByTime *gorm.DB
	var debtByTime2 *gorm.DB
	var resultErr error
	var result []Result
	var result2 []Result2

	timeDesc := TimeDesc(c)

	for _, value := range timeDesc {
		debtByTime = config.DB.Model(&debts).Where("date = ? AND debtor_id = ?", value, claims.Id).Find(&result)
		debtByTime2 = config.DB.Model(&debts).Select("sum(amount) AS total").Where("date = ? AND debtor_id = ?", value, claims.Id).Find(&result2)

		if err := debtByTime.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := debtByTime2.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		resultErr = c.JSON(http.StatusOK, map[string]interface{}{
			value:        result,
			"total debt": result2,
		})
	}

	return resultErr
}

func TimeDesc(c echo.Context) []string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	var debts []model.Debt
	var result []string

	dateDesc := config.DB.Order("date desc").Model(&debts).Distinct("date").Where("debtor_id = ?", claims.Id).Find(&result)

	if err := dateDesc.Error; err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return result
}

// lihat semua debts berdasarkan kreditor
func GetAllDebtByCreditorController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	type Result struct {
		CreditorName string
		Date         datatypes.Date
		Amount       int
		Detail       string
	}

	type Result2 struct {
		Total int
	}

	var debts []model.Debt
	var debtByCreditor *gorm.DB
	var debtByCreditor2 *gorm.DB
	var resultErr error
	var result []Result
	var result2 []Result2

	creditorNameDesc := CreditorNameAsc(c)

	for _, value := range creditorNameDesc {
		debtByCreditor = config.DB.Model(&debts).Where("creditor_name = ? AND debtor_id = ?", value, claims.Id).Find(&result)
		debtByCreditor2 = config.DB.Model(&debts).Select("sum(amount) AS total").Where("creditor_name = ? AND debtor_id = ?", value, claims.Id).Find(&result2)

		if err := debtByCreditor.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := debtByCreditor2.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		resultErr = c.JSON(http.StatusOK, map[string]interface{}{
			value:        result,
			"total debt": result2,
		})
	}

	return resultErr
}

func CreditorNameAsc(c echo.Context) []string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	var debts []model.Debt
	var result []string

	creditorName := config.DB.Order("creditor_name asc").Model(&debts).Distinct("creditor_name").Where("debtor_id = ?", claims.Id).Find(&result)

	if err := creditorName.Error; err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return result
}

func GetDebtByTimeController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	type Result struct {
		Total int
	}

	type Result2 struct {
		CreditorName string
		Date         datatypes.Date
		Amount       int
		Detail       string
	}

	var debts []model.Debt
	var result []Result
	var result2 []Result2

	date := c.QueryParam("date")

	debtByTime := config.DB.Model(&debts).Select("sum(amount) AS total").Where("date = ? AND debtor_id = ?", date, claims.Id).Find(&result)
	debtByTime2 := config.DB.Model(&debts).Where("date = ? AND debtor_id = ?", date, claims.Id).Find(&result2)

	if err := debtByTime.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := debtByTime2.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"total debt": result,
		date:         result2,
	})
}

// cari debt berdasarkan kreditor
func GetDebtByCreditorController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	type Result struct {
		Total int
	}

	type Result2 struct {
		CreditorName string
		Date         datatypes.Date
		Amount       int
		Detail       string
	}

	var debts []model.Debt
	var result []Result
	var result2 []Result2

	creditor := c.QueryParam("creditor_name")

	debtByCreditor := config.DB.Model(&debts).Select("sum(amount) AS total").Where("creditor_name = ? AND debtor_id = ?", creditor, claims.Id).Find(&result)
	debtByCreditor2 := config.DB.Model(&debts).Where("creditor_name = ? AND debtor_id = ?", creditor, claims.Id).Find(&result2)

	if err := debtByCreditor.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := debtByCreditor2.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"total debt": result,
		creditor:     result2,
	})
}

// urutkan debts terbanyak (per orang)
func GetAllDebtByTheHighest(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	type Result struct {
		Total int
	}

	type Result2 struct {
		CreditorName string
		Date         datatypes.Date
		Amount       int
		Detail       string
	}

	var debts []model.Debt
	var debtByHighest *gorm.DB
	var resultErr error
	var result Result
	var result2 []Result2

	creditorNameHighest, creditorTotalHighest := DebtHighest(c)

	for i, value := range creditorNameHighest {
		result.Total = creditorTotalHighest[i]
		debtByHighest = config.DB.Model(&debts).Where("creditor_name = ? AND debtor_id = ?", value, claims.Id).Find(&result2)

		if err := debtByHighest.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		resultErr = c.JSON(http.StatusOK, map[string]interface{}{
			"total debt": result,
			"debts":      result2,
		})
	}

	return resultErr
}

func DebtHighest(c echo.Context) ([]string, []int) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

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

// urutkan debts terlama (per orang)
func GetAllDebtByTheLongest(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	type Result struct {
		Total int
	}

	type Result2 struct {
		CreditorName string
		Date         datatypes.Date
		Amount       int
		Detail       string
	}

	var debts []model.Debt
	var debtByLongest *gorm.DB
	var resultErr error
	var result Result
	var result2 []Result2

	creditorNameLongest, creditorTotalLongest := DebtLongest(c)

	for i, value := range creditorNameLongest {
		// result.CreditorName = value
		result.Total = creditorTotalLongest[i]
		debtByLongest = config.DB.Order("date asc").Model(&debts).Where("creditor_name = ? AND debtor_id = ?", value, claims.Id).Find(&result2)

		if err := debtByLongest.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		resultErr = c.JSON(http.StatusOK, map[string]interface{}{
			"total day": result,
			"debts":     result2,
		})
	}

	return resultErr
}

func DebtLongest(c echo.Context) ([]string, []int) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	var debts []model.Debt
	var resultName []string
	var resultTotal []int

	subQuery := config.DB.Model(&debts).Select("min(date)").Where("debtor_id = ?", claims.Id).Group("creditor_name")
	creditorName := config.DB.Order("datediff(curdate(), date) desc").Model(&debts).Select("creditor_name").Where("debtor_id = ?", claims.Id).Group("creditor_name").Find(&resultName)
	creditorTotal := config.DB.Order("datediff(curdate(), date) desc").Model(&debts).Select("datediff(curdate(), date)").Where("date = any (?) AND debtor_id = ?", subQuery, claims.Id).Group("creditor_name").Find(&resultTotal)

	if err := creditorName.Error; err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := creditorTotal.Error; err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return resultName, resultTotal
}