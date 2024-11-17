package request

// SuggestionRequest is the request for the suggestion endpoint
// @Description SuggestionRequest is the request for the suggestion endpoint
// @Param City string true "City of the suggestion"
type SuggestionRequest struct {
	City string `json:"city"`
}
