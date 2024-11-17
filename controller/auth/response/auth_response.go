package response

import "energia/entities"

// AuthResponse is the response for the auth controller
// @Description AuthResponse is the response for the auth controller
// @Param ID int true "ID of the user"
// @Param Email string true "Email of the user"
// @Param Token string true "Token of the user"
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
