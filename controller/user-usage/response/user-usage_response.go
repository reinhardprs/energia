package response

import "energia/entities"

type UserUsageResponse struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Date        string  `json:"date"`
	TotalEnergy float32 `json:"total_energy"`
	TotalCost   float32 `json:"total_cost"`
}

func FromEntities(userUsage entities.UserUsage) UserUsageResponse {
	return UserUsageResponse{
		ID:          userUsage.UserID,
		UserID:      userUsage.UserID,
		Date:        userUsage.Date.String(),
		TotalEnergy: userUsage.TotalEnergy,
		TotalCost:   userUsage.TotalCost,
	}
}

func FromEntitiesArray(userUsages []entities.UserUsage) []UserUsageResponse {
	var userUsageResponses []UserUsageResponse
	for _, userUsage := range userUsages {
		userUsageResponses = append(userUsageResponses, FromEntities(userUsage))
	}
	return userUsageResponses
}
