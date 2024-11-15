package user_usage

import (
	"energia/entities"
	device_usage "energia/repository/device-usage"
	user_usage "energia/repository/user-usage"
	"time"
)

func NewUserUsageService(ur user_usage.UserUsageRepoInterface, dur device_usage.DeviceUsageRepoInterface) *UserUsageService {
	return &UserUsageService{
		userUsageRepoInterface:   ur,
		deviceUsageRepoInterface: dur,
	}
}

type UserUsageService struct {
	userUsageRepoInterface   user_usage.UserUsageRepoInterface
	deviceUsageRepoInterface device_usage.DeviceUsageRepoInterface
}

func (u *UserUsageService) GetUserUsage(userID int) ([]entities.UserUsage, error) {
	userUsage, err := u.userUsageRepoInterface.FindAll(userID)
	if err != nil {
		return nil, err
	}
	return userUsage, nil
}

func (u *UserUsageService) Create(userID int, date time.Time) (entities.UserUsage, error) {

	deviceUsages, err := u.deviceUsageRepoInterface.GetDeviceUsageByDate(userID, date)
	if err != nil {
		return entities.UserUsage{}, err
	}

	var totalEnergy float32
	for _, deviceUsage := range deviceUsages {
		totalEnergy += deviceUsage.EnergyConsumed
	}

	totalCost := totalEnergy * 1352

	userUsage := entities.UserUsage{
		UserID:      userID,
		Date:        date,
		TotalEnergy: totalEnergy,
		TotalCost:   totalCost,
	}

	savedUserUsage, err := u.userUsageRepoInterface.Create(userUsage)
	if err != nil {
		return entities.UserUsage{}, err
	}

	return savedUserUsage, nil
}
