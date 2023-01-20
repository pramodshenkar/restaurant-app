package logreport

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/pramodshenkar/restaurant-app/models"
)

func GetLogFileDAO() (models.LogReport, error) {

	var logReport models.LogReport

	logfile, err := ioutil.ReadDir(models.LogDir)
	if err != nil {
		log.Println("error while getting file  from: ", models.LogDir, " : ", err)
		return models.LogReport{}, err
	}

	if len(logfile) <= 0 {
		log.Println("no logfile present in : ", models.LogDir)
		return models.LogReport{}, os.ErrNotExist

	}
	for _, filename := range logfile {

		path := models.LogDir + filename.Name()
		if err != nil {
			log.Println("error while getting file : ", path, " : ", err)
			return models.LogReport{}, err
		}

		logBytes, err := os.ReadFile(path)
		if err != nil {
			log.Println("GetLogFileDAO: error while reading file : ", path, " : ", err)
			return models.LogReport{}, err
		}

		var logReportFile models.LogReport
		err = json.Unmarshal(logBytes, &logReportFile)
		if err != nil {
			log.Println("GetLogFileDAO: error while unmarshalling file : ", path, " : ", err)
			return models.LogReport{}, err
		}

		logReport = append(logReport, logReportFile...)
	}

	if len(logReport) <= 0 {
		log.Println("GetLogFileDAO: error : no records found in logfile")
		return models.LogReport{}, errors.New("error : no log found in logfile")
	}

	return logReport, err
}
