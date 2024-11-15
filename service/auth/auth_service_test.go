package auth_test

import (
	"energia/constant"
	"energia/entities"
	"energia/mocks"
	"energia/service/auth"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_Login_Success(t *testing.T) {
	mockAuthRepo := new(mocks.AuthRepoInterface)
	mockJwt := new(mocks.JwtInterface)
	authService := auth.NewAuthService(mockAuthRepo, mockJwt)

	hashedPassword, _ := auth.HashPassword("password")
	dbUser := entities.User{Email: "test@example.com", Password: hashedPassword}

	mockAuthRepo.On("Login", mock.Anything).Return(dbUser, nil)
	mockJwt.On("GenerateJWT", dbUser.ID, dbUser.Email).Return("mockToken", nil)

	result, err := authService.Login(entities.User{Email: "test@example.com", Password: "password"})

	assert.NoError(t, err)
	assert.Equal(t, "mockToken", result.Token)
	mockAuthRepo.AssertExpectations(t)
	mockJwt.AssertExpectations(t)
}


func TestAuthService_Login_InvalidPassword(t *testing.T) {
	mockAuthRepo := new(mocks.AuthRepoInterface)
	mockJwt := new(mocks.JwtInterface)
	authService := auth.NewAuthService(mockAuthRepo, mockJwt)

	dbUser := entities.User{Email: "test@example.com", Password: "$2a$10$hashedPassword"}
	mockAuthRepo.On("Login", mock.Anything).Return(dbUser, nil)

	_, err := authService.Login(entities.User{Email: "test@example.com", Password: "wrongpassword"})

	assert.Error(t, err)
	assert.Equal(t, constant.INVALID_PASSWORD, err)
	mockAuthRepo.AssertExpectations(t)
}

func TestAuthService_Register_Success(t *testing.T) {
	mockAuthRepo := new(mocks.AuthRepoInterface)
	mockJwt := new(mocks.JwtInterface)
	authService := auth.NewAuthService(mockAuthRepo, mockJwt)

	user := entities.User{Email: "test@example.com", Password: "hashedpassword"}
	mockAuthRepo.On("Register", mock.Anything).Return(user, nil)

	result, err := authService.Register(entities.User{Email: "test@example.com", Password: "password"})

	assert.NoError(t, err)
	assert.Equal(t, user.Email, result.Email)
	mockAuthRepo.AssertExpectations(t)
}

func TestAuthService_Register_EmailEmpty(t *testing.T) {
	mockAuthRepo := new(mocks.AuthRepoInterface)
	mockJwt := new(mocks.JwtInterface)
	authService := auth.NewAuthService(mockAuthRepo, mockJwt)

	_, err := authService.Register(entities.User{Password: "password"})

	assert.Error(t, err)
	assert.Equal(t, constant.EMAIL_IS_EMPTY, err)
	mockAuthRepo.AssertNotCalled(t, "Register")
}
