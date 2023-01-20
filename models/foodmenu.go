package models

type FoodMenuList map[string]FoodmenuDetails

type FoodmenuDetails struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
