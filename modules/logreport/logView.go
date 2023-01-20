package logreport

import "github.com/pramodshenkar/restaurant-app/models"

type FoodmenuReportModel struct {
	FoodmenuId string
	LogReport  models.FoodmenuDetails
	OrderCount int
}

func GetTopFoodmenuConsumed(menuCount int) ([]FoodmenuReportModel, error) {
	return GetTopFoodMenuConsumedService(menuCount)
}
