package models

import "time"

type LogReport []LogEntry

type LogEntry struct {
	EaterId    string    `json:"eater_id"`
	FoodmenuId string    `json:"foodmenu_id"`
	ServedAt   time.Time `json:"served_at"`
}
