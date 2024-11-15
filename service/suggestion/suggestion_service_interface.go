package suggestion

import (
	"context"
	"energia/entities"
)

type SuggestionServiceInterface interface {
	GetSuggestion(ctx context.Context, userID int, city string) (entities.Suggestion, error)
}
