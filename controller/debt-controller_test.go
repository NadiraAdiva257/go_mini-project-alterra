package controller

import (
	"bytes"
	"encoding/json"
	"mini-project/middleware"
	"mini-project/model"
	"mini-project/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/datatypes"
)

func TestCreateDebtController(t *testing.T) {
	debtRepository := &service.DebtRepositoryMock{Mock: mock.Mock{}}
	service.SetDebtRepository(debtRepository)

	dataDebt := model.Debt{
		CreditorName: "mayla",
		Date:         datatypes.Date(time.Now()),
		Amount:       23000,
		Detail:       "ayam",
		DebtorID:     1,
	}

	debtRepository.Mock.On("CreateDebtController", &dataDebt).Return(nil)

	e := echo.New()

	bDataDebt, _ := json.Marshal(dataDebt)
	req := httptest.NewRequest(http.MethodPost, "/debts", bytes.NewReader(bDataDebt))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	user := &middleware.JwtCustomClaims{}
	c.Set("user", &jwt.Token{
		Claims: user,
	})

	CreateDebtController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateDebtController(t *testing.T) {
	debtRepository := &service.DebtRepositoryMock{Mock: mock.Mock{}}
	service.SetDebtRepository(debtRepository)

	dataDebt := model.Debt{
		CreditorName: "mayla",
		Date:         datatypes.Date(time.Now()),
		Amount:       23000,
		Detail:       "ayam",
		DebtorID:     1,
	}

	debtRepository.Mock.On("UpdateDebtController", &dataDebt, 1, 2).Return(nil)

	e := echo.New()

	bDataDebt, _ := json.Marshal(dataDebt)
	req := httptest.NewRequest(http.MethodPut, "/debts:id", bytes.NewReader(bDataDebt))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	user := &middleware.JwtCustomClaims{}
	c.Set("user", &jwt.Token{
		Claims: user,
	})

	UpdateDebtController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteDebtController(t *testing.T) {
	debtRepository := &service.DebtRepositoryMock{Mock: mock.Mock{}}
	service.SetDebtRepository(debtRepository)

	debtRepository.Mock.On("DeleteDebtController", 1, 2).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/debts/:id", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	user := &middleware.JwtCustomClaims{}
	c.Set("user", &jwt.Token{
		Claims: user,
	})

	DeleteDebtController(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	expectResult := map[string]interface{}{
		"message": "success delete debt by id",
	}
	var resultJSON map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resultJSON)
	assert.Equal(t, expectResult, resultJSON)
}

func TestGetDebtByCreditorController(t *testing.T) {
	debtRepository := &service.DebtRepositoryMock{Mock: mock.Mock{}}
	service.SetDebtRepository(debtRepository)

	var result = make(map[string]interface{})
	result["total debt"] = 23000
	result["mayla"] = []model.Debt{
		{
			Date:   datatypes.Date(time.Now()),
			Amount: 23000,
			Detail: "cilok",
		},
		{
			Date:   datatypes.Date(time.Now()),
			Amount: 23000,
			Detail: "cilok",
		},
	}

	debtRepository.Mock.On("GetDebtByCreditorController", "mayla", 1).Return(nil)

	e := echo.New()

	bDataDebt, _ := json.Marshal(result)
	req := httptest.NewRequest(http.MethodGet, "/debts/creditor", bytes.NewReader(bDataDebt))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.QueryParam("creditor_name")
	c.SetParamValues("mayla")

	user := &middleware.JwtCustomClaims{}
	c.Set("user", &jwt.Token{
		Claims: user,
	})

	GetDebtByCreditorController(c)
	assert.Equal(t, http.StatusOK, rec.Code)

	// expectResult := map[string]interface{}{
	// 	"message": "success delete debt by id",
	// }
	// var resultJSON map[string]interface{}
	// json.Unmarshal(rec.Body.Bytes(), &resultJSON)
	// assert.Equal(t, expectResult, resultJSON)
}
