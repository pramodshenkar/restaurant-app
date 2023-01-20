package models

import (
	"log"
	"os"
)

// Directories
var (
	DatabaseDir  = ""
	LogDir       = "LogDB"
	RestaurantDB = "RestaurantDB"
)

// Files
var (
	FoodMenuFile = "foodmenu.json"
)

func SetDatabaseDir(path ...string) {
	var workDir = ""
	var err error
	if len(path) <= 0 {
		workDir, err = os.Getwd()
		if err != nil {
			log.Println(err)
		}
	} else {
		workDir = path[0]
	}

	DatabaseDir = workDir + "/database/"
}

func SetLogDir() {
	LogDir = DatabaseDir + LogDir + "/"
}

func FoodMenuFilePath() string {
	return DatabaseDir + "/" + RestaurantDB + "/" + FoodMenuFile
}
