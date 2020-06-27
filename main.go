package main

import (
	"fmt"
	"rest-api-go/config"
	Models "rest-api-go/models"
	"rest-api-go/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&Models.User{})
	r := routes.SetupRouter()
	//running
	r.Run()
}
