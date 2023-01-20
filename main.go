package main

import (
	"log"
	"os"

	"github.com/pramodshenkar/restaurant-app/modules/commen"
	"github.com/pramodshenkar/restaurant-app/modules/logreport"
)

func main() {
	err := commen.Init()
	if err != nil {
		log.Println("Error while initializing app : ", err)
		return
	}
	ShowLogReport()
}

func ShowLogReport() {
	var menuCount int = 3

	output := log.New(os.Stdout, "", 0)
	errs := log.New(os.Stderr, "Err : ", 0)

	menuReport, err := logreport.GetTopFoodmenuConsumed(menuCount)
	if err != nil {
		errs.Fatal("Error while getting top ", menuCount, " foodmenus : ", err)
		return
	}

	if menuCount > len(menuReport) {
		output.Println("Only ", len(menuReport), " foodmenus log found for Restaurant")
	}

	output.Println("Top", len(menuReport), "foodmenus in Restaurant are : ")
	output.Println("-----------------------------------------------------------------")
	output.Println("| FoodMenuId \t| Name  \t| Price \t| OrderCount \t|")
	output.Println("-----------------------------------------------------------------")

	for _, foodmenu := range menuReport {
		output.Println("| ", foodmenu.FoodmenuId, "\t\t|", foodmenu.LogReport.Name, "\t|", foodmenu.LogReport.Price, "\t\t|", foodmenu.OrderCount, "\t\t|")
	}
	output.Println("-----------------------------------------------------------------")
}
