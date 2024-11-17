package response

import "energia/entities"

// SuggestionResponse is the response for the suggestion endpoint
// @Description SuggestionResponse is the response for the suggestion endpoint
// @Param Message string true "Message of the suggestion"
type SuggestionResponse struct {
	Message string `json:"message"`
}

func FromEntities(suggestion entities.Suggestion) SuggestionResponse {
	return SuggestionResponse{
		Message: suggestion.Message,
	}
}
