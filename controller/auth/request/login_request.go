package request

import "energia/entities"

// LoginRequest is the request for the login endpoint
// @Description LoginRequest is the request for the login endpoint
// @Param Email string true "Email of the user"
// @Param Password string true "Password of the user"
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest is the request for the register endpoint
// @Description RegisterRequest is the request for the register endpoint
// @Param Email string true "Email of the user"
// @Param Password string true "Password of the user"
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
