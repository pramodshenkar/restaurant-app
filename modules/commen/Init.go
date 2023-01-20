package commen

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/pramodshenkar/restaurant-app/models"
)

func Init() error {
	//set log

	workDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}

	logfilePath := workDir + "/logs/" + time.Now().Format("2006-02-01") + ".log"

	logfile, err := os.OpenFile(logfilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	log.SetOutput(io.MultiWriter(logfile))

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//set directory/DB paths & other vars
	models.SetDatabaseDir()
	models.SetLogDir()
	return nil
}
