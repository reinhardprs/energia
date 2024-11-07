package auth

import (
	"energia/constant"
	"energia/entities"
	"energia/middleware"
	"energia/repository/auth"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(ar auth.AuthRepoInterface, jt middleware.JwtInterface) *AuthService {
	return &AuthService{
		authRepoInterface: ar,
		jwtInterface:      jt,
	}
}

type AuthService struct {
	authRepoInterface auth.AuthRepoInterface
	jwtInterface      middleware.JwtInterface
}

func (authService AuthService) Login(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, constant.EMAIL_IS_EMPTY
	} else if user.Password == "" {
		return entities.User{}, constant.PASSWORD_IS_EMPTY
	}

	var err error
	dbUser, err := authService.authRepoInterface.Login(user)
	if err != nil {
		return entities.User{}, err
	}

	isPasswordValid := CheckPasswordHash(user.Password, dbUser.Password)
	if !isPasswordValid {
		return entities.User{}, constant.INVALID_PASSWORD
	}

	token, err := authService.jwtInterface.GenerateJWT(dbUser.ID, dbUser.Email)
	if err != nil {
		return entities.User{}, err
	}
	dbUser.Token = token

	return dbUser, nil
}

func (authService AuthService) Register(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, constant.EMAIL_IS_EMPTY
	} else if user.Password == "" {
		return entities.User{}, constant.PASSWORD_IS_EMPTY
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		return entities.User{}, err
	}
	user.Password = hash

	user, err = authService.authRepoInterface.Register(user)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
