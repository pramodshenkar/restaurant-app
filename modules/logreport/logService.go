package logreport

import (
	"errors"
	"log"

	"github.com/pramodshenkar/restaurant-app/modules/commen"
)

func GetTopFoodMenuConsumedService(menuCount int) ([]FoodmenuReportModel, error) {

	//Get logs

	logReport, err := GetLogFileDAO()
	if err != nil {
		log.Println("GetTopFoodMenuConsumedService : error while getting LogFile : ", err)
		return []FoodmenuReportModel{}, err
	}

	// Process logs

	var FoodMenuReportMap = make(map[string][]string)

	for _, logEntry := range logReport {
		if len(FoodMenuReportMap[logEntry.FoodmenuId]) > 0 {
			for _, eaterId := range FoodMenuReportMap[logEntry.FoodmenuId] {
				if logEntry.EaterId == eaterId {
					log.Println("GetTopFoodMenuConsumedService : error eaterId with same foodmenuId found ")
					return []FoodmenuReportModel{}, errors.New("error : eaterId with same foodmenuId found")
				}
			}
		}
		FoodMenuReportMap[logEntry.FoodmenuId] = append(FoodMenuReportMap[logEntry.FoodmenuId], logEntry.EaterId)
	}

	// Get foodmenu details
	FoodmenuList, err := commen.GetFoodmenuFileDAO()
	if err != nil {
		log.Println("GetTopFoodMenuConsumedService : error while getting Foodmenu file : ", err)
		return []FoodmenuReportModel{}, err
	}

	// Generate LogReport

	var FoodmenuReport []FoodmenuReportModel
	for foodmenuId, eaterDetails := range FoodMenuReportMap {
		report := FoodmenuReportModel{foodmenuId, FoodmenuList[foodmenuId], len(eaterDetails)}

		if len(FoodmenuReport) <= 0 {
			FoodmenuReport = append(FoodmenuReport, report)
		} else if len(FoodmenuReport) > 0 && len(FoodmenuReport) < menuCount {
			var isElementAdded bool
			for i, foodmenuReport := range FoodmenuReport {
				if foodmenuReport.OrderCount <= len(eaterDetails) {
					FoodmenuReport = append(FoodmenuReport[:i+1], FoodmenuReport[i:]...)
					FoodmenuReport[i] = report
					isElementAdded = true
					break
				}
			}
			if !isElementAdded {
				FoodmenuReport = append(FoodmenuReport, report)
			}

		} else {
			for i, foodmenuReport := range FoodmenuReport {
				if foodmenuReport.OrderCount <= len(eaterDetails) {
					FoodmenuShiftedItems := make([]FoodmenuReportModel, len(FoodmenuReport[i:len(FoodmenuReport)-1]))
					_ = copy(FoodmenuShiftedItems, FoodmenuReport[i:len(FoodmenuReport)-1])

					FoodmenuReport = append(FoodmenuReport[:i], report)
					FoodmenuReport = append(FoodmenuReport, FoodmenuShiftedItems...)

					break
				}
			}
		}

	}
	return FoodmenuReport, nil
}
