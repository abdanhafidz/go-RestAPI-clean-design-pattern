package utils

import (
	"log"
	"os"

	"go-dp.abdanhafidz.com/config"
)

func LogError(errorLogged error) {
	file, err := os.OpenFile(config.LOG_PATH+"error_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println(errorLogged)
}
