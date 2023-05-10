package controller

import (
	"encoding/json"
	"mini-project/middleware"
	"mini-project/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
