package response

// EmailResponse is the response for the email endpoint
// @Description EmailResponse is the response for the email endpoint
// @Param Message string true "Message of the email"
type EmailResponse struct {
	Message string `json:"message"`
}

func FromEntities(to string) EmailResponse {
	return EmailResponse{
		Message: "Laporan telah dikirimkan ke " + to + " melalui email.",
	}
}
