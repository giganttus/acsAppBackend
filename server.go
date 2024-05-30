package main

import (
	"acsApp/helpers"
	"acsApp/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(helpers.ErrEnvLoad)
	}

	logFilename := os.Getenv("LOG_FILENAME")
	if logFilename == "" {
		panic(helpers.ErrEnvLoad)
	}

	logFile, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening log file: %v\n", err)
	}
	log.SetOutput(logFile)
	defer logFile.Close()

	if os.Getenv("CON_STRING") == "" {
		panic(helpers.ErrEnvLoad)
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("CON_STRING")), &gorm.Config{})
	if err != nil {
		panic(helpers.ErrDbConn)
	}

	rlRate, rlBurst := helpers.GetRlValues()
	routes := routes.InitRoutes(db, rlRate, rlBurst)

	routes.Run("0.0.0.0:1312")
}
