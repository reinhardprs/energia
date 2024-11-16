package response

// EmailResponse mendefinisikan format response untuk pengiriman email
type EmailResponse struct {
	Message string `json:"message"`
}

// FromEntities mengonversi entitas email menjadi response
func FromEntities(to string) EmailResponse {
	return EmailResponse{
		Message: "Laporan telah dikirimkan ke " + to + " melalui email.",
	}
}
