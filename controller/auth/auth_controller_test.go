package auth_test

import (
	"bytes"
	"encoding/json"
	"energia/controller/auth"
	"energia/controller/auth/request"
	"energia/entities"
	"energia/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLoginController(t *testing.T) {
	e := echo.New()
	mockAuthService := new(mocks.AuthInterface)
	authController := auth.NewAuthController(mockAuthService)

	loginReq := request.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	loginReqBody, _ := json.Marshal(loginReq)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(loginReqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	user := entities.User{
		ID:    1,
		Email: "test@example.com",
		Token: "dummyToken",
	}

	mockAuthService.On("Login", loginReq.ToEntities()).Return(user, nil)

	err := authController.LoginController(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var resp map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, true, resp["status"])
		assert.Equal(t, user.Email, resp["data"].(map[string]interface{})["email"])
		assert.Equal(t, user.Token, resp["data"].(map[string]interface{})["token"])
	}
	mockAuthService.AssertExpectations(t)
}

func TestRegisterController(t *testing.T) {
	e := echo.New()
	mockAuthService := new(mocks.AuthInterface)
	authController := auth.NewAuthController(mockAuthService)

	registerReq := request.RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	registerReqBody, _ := json.Marshal(registerReq)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(registerReqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	user := entities.User{
		ID:    1,
		Email: "test@example.com",
		Token: "dummyToken",
	}

	mockAuthService.On("Register", registerReq.ToEntities()).Return(user, nil)

	err := authController.RegisterController(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var resp map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, true, resp["status"])
		assert.Equal(t, user.Email, resp["data"].(map[string]interface{})["email"])
		assert.Equal(t, user.Token, resp["data"].(map[string]interface{})["token"])
	}
	mockAuthService.AssertExpectations(t)
}

func TestUserController(t *testing.T) {
	e := echo.New()
	authController := auth.NewAuthController(nil)

	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": 1,
	})
	c.Set("user", token)

	err := authController.UserController(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var resp map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "User authenticated", resp["message"])
		assert.Equal(t, float64(1), resp["userID"])
	}
}
