package request

import "energia/entities"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (loginRequest LoginRequest) ToEntities() entities.User {
	return entities.User{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
}

func (registerRequest RegisterRequest) ToEntities() entities.User {
	return entities.User{
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
}
