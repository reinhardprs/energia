package response

import "energia/entities"

type SuggestionResponse struct {
	Message string `json:"message"`
}

func FromEntities(suggestion entities.Suggestion) SuggestionResponse {
	return SuggestionResponse{
		Message: suggestion.Message,
	}
}
