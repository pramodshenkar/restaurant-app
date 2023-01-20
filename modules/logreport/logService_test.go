package logreport

import (
	"errors"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/pramodshenkar/restaurant-app/models"
)

func TestGetTopFoodMenuConsumedService(t *testing.T) {
	menuCount := 3

	type TestCase struct {
		TestCaseId  string
		Description string
		Type        string
		MockInit    func()
		Want        []FoodmenuReportModel
		Err         error
	}

	var testCases = []TestCase{
		{
			TestCaseId:  "TC_show_3_item",
			Description: "Getting 3 menus successfully",
			Type:        "Positive",
			MockInit: func() {
				SetTestCasePath("TestDB/TC_show_3_item")
			},
			Want: []FoodmenuReportModel{{FoodmenuId: "f3", LogReport: models.FoodmenuDetails{"Fries", 39}, OrderCount: 5}, {FoodmenuId: "f7", LogReport: models.FoodmenuDetails{"Salad", 69}, OrderCount: 4}, {FoodmenuId: "f6", LogReport: models.FoodmenuDetails{"Pasta", 49}, OrderCount: 3}},
			Err:  nil,
		},
		{
			TestCaseId:  "TC_show_2_item",
			Description: "Logfile contains logs for less than 3 elements",
			Type:        "Positive",
			MockInit: func() {
				SetTestCasePath("TestDB/TC_show_2_item")
			},
			Want: []FoodmenuReportModel{{FoodmenuId: "f2", LogReport: models.FoodmenuDetails{"Burger", 79}, OrderCount: 6}, {FoodmenuId: "f1", LogReport: models.FoodmenuDetails{"Coffee", 99}, OrderCount: 5}},
			Err:  nil,
		},
		{
			TestCaseId:  "TC_no_item_found",
			Description: "No items in presnet in logfile",
			Type:        "Negative",
			MockInit: func() {
				SetTestCasePath("TestDB/TC_no_item_found")
			},
			Want: []FoodmenuReportModel{},
			Err:  errors.New("error : no log found in logfile"),
		},
		{
			TestCaseId:  "TC_error_same_log_found",
			Description: "EaterId with same FoodmenuId",
			Type:        "Negative",

			MockInit: func() {
				SetTestCasePath("TestDB/TC_error_same_log_found")
			},
			Want: []FoodmenuReportModel{},
			Err:  errors.New("error : eaterId with same foodmenuId found"),
		},
		{
			TestCaseId:  "TC_file_absent",
			Description: "Logfile or Foodmenufile absent",
			Type:        "Negative",
			MockInit: func() {
				SetTestCasePath("TestDB/TC_file_absent")
			},
			Want: []FoodmenuReportModel{},
			Err:  os.ErrNotExist,
		},
	}

	for _, testcase := range testCases {

		testcase.MockInit()
		got, err := GetTopFoodMenuConsumedService(menuCount)
		if err != nil && testcase.Type == "Positive" {
			t.Error(err)
		}

		if !reflect.DeepEqual(got, testcase.Want) {
			if testcase.Type == "Positive" {
				t.Errorf("TestCaseId: %v, Description: %v, Type: %v, Result got: %v, Result want:%v ", testcase.TestCaseId, testcase.Description, testcase.Type, got, testcase.Want)
			} else if testcase.Type == "Negative" {
				t.Errorf("TestCaseId: %v, Description: %v, Type: %v, Error got: %v, Error want:%v ", testcase.TestCaseId, testcase.Description, testcase.Type, err.Error(), testcase.Err.Error())
			}
		}

	}
}

func SetTestCasePath(relativepath string) {
	workDir, _ := os.Getwd()
	models.SetDatabaseDir(strings.ReplaceAll(workDir, "modules/logreport", ""))
	models.LogDir = models.DatabaseDir + "/" + relativepath + "/"
}

func BenchmarkGetTopFoodMenuConsumedService(b *testing.B) {
	menuCount := 3
	SetTestCasePath("TestDB/TC_show_3_item")

	for i := 0; i < b.N; i++ {
		_, err := GetTopFoodMenuConsumedService(menuCount)
		if err != nil {
			b.Errorf("error %v", err)
		}
	}
}
