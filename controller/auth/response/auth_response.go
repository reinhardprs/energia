package response

import "energia/entities"

type AuthResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FromEntities(user entities.User) AuthResponse {
	return AuthResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: user.Token,
	}
}
