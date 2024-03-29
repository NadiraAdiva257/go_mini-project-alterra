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

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var dataDebtor = model.Debtor{
	Name:     "nadira",
	Email:    "nadira123@gmail.com",
	Password: "nadira123",
}

func TestCreateDebtorController(t *testing.T) {
	debtorRepository := &service.DebtorRepositoryMock{Mock: mock.Mock{}}
	service.SetDebtorRepository(debtorRepository)

	debtorRepository.Mock.On("CreateDebtorController", &dataDebtor).Return(nil)

	e := echo.New()

	bDataDebtor, _ := json.Marshal(dataDebtor)
	req := httptest.NewRequest(http.MethodPost, "/debtors", bytes.NewReader(bDataDebtor))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateDebtorController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateDebtorController(t *testing.T) {
	debtorRepository := &service.DebtorRepositoryMock{Mock: mock.Mock{}}
	service.SetDebtorRepository(debtorRepository)

	debtorRepository.Mock.On("UpdateDebtorController", &dataDebtor, 1).Return(nil)

	e := echo.New()

	bDataDebtor, _ := json.Marshal(dataDebtor)
	req := httptest.NewRequest(http.MethodPut, "/debtors", bytes.NewReader(bDataDebtor))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	user := &middleware.JwtCustomClaims{}
	c.Set("user", &jwt.Token{
		Claims: user,
	})

	UpdateDebtorController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}
