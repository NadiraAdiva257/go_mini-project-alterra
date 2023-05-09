package controller

import (
	"bytes"
	"encoding/json"
	"mini-project/model"
	"mini-project/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateDebtorController(t *testing.T) {
	debtorRepository := &service.DebtorRepositoryMock{Mock: mock.Mock{}}
	service.SetDebtorRepository(debtorRepository)

	dataDebtor := model.Debtor{
		Name:     "nadira",
		Email:    "nadira123@gmail.com",
		Password: "nadira123",
	}

	debtorRepository.Mock.On("CreateDebtorController", &dataDebtor).Return(nil)

	e := echo.New()

	bDataDebtor, _ := json.Marshal(dataDebtor)
	req := httptest.NewRequest(http.MethodPost, "/debtors", bytes.NewReader(bDataDebtor))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateDebtorController(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	expectResult := map[string]interface{}{
		"message": "success create new user",
	}
	var resultJSON map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resultJSON)
	assert.Equal(t, expectResult, resultJSON)
}

func TestUpdateDebtorController(t *testing.T) {
	debtorRepository := &service.DebtorRepositoryMock{Mock: mock.Mock{}}
	service.SetDebtorRepository(debtorRepository)

	dataDebtor := model.Debtor{
		Name: "nadira",
	}

	debtorRepository.Mock.On("UpdateDebtorController", &dataDebtor, 1).Return(nil)

	e := echo.New()

	bDataDebtor, _ := json.Marshal(dataDebtor)
	req := httptest.NewRequest(http.MethodPut, "/debtors", bytes.NewReader(bDataDebtor))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	UpdateDebtorController(c)

	assert.Equal(t, http.StatusOK, rec.Code)

	expectResult := map[string]interface{}{
		"message": "succes update debtor by id",
	}
	var resultJSON map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resultJSON)
	assert.Equal(t, expectResult, resultJSON)
}
