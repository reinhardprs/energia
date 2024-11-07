package middleware

type JwtInterface interface {
	GenerateJWT(userID int, email string) (string, error)
}
