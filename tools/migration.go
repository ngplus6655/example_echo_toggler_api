package main

import (
	"hello/database"
	"hello/models"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		logrus.Fatal(err)
	}
	db.Debug().AutoMigrate(&models.BoolToggler{})
}