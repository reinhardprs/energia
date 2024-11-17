package response

import "energia/entities"

// UserUsageResponse is the response for the user usage endpoint
// @Description UserUsageResponse is the response for the user usage endpoint
// @Param ID int true "ID of the user usage"
// @Param UserID int true "ID of the user"
// @Param Date string true "Date of the user usage"
// @Param TotalEnergy float32 true "Total energy of the user usage"
// @Param TotalCost float32 true "Total cost of the user usage"
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
