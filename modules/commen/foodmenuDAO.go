package commen

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/pramodshenkar/restaurant-app/models"
)

func GetFoodmenuFileDAO() (models.FoodMenuList, error) {

	//read foodmenu file
	filemenuBytes, err := os.ReadFile(models.FoodMenuFilePath())
	if err != nil {
		log.Println("Error while reading file : ", models.FoodMenuFilePath())
		return models.FoodMenuList{}, err
	}

	//unmarshal foodmenu
	FoodMenuList := models.FoodMenuList{}
	err = json.Unmarshal(filemenuBytes, &FoodMenuList)
	if err != nil {
		log.Println("Error while unmarshalling file : ", models.FoodMenuFilePath())
		return models.FoodMenuList{}, err
	}

	if len(FoodMenuList) <= 0 {
		log.Println("Foodmenu file is empty")
		return models.FoodMenuList{}, errors.New("foodmenu file is empty")
	}

	//return
	return FoodMenuList, nil
}
