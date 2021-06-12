package main

import (
	"github.com/lion-devs/ilife-api/pkg/database"
	"github.com/lion-devs/ilife-api/pkg/model"
	"github.com/lion-devs/ilife-api/pkg/route"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var err error

func main() {

	// init database connection
	database.DB, err = gorm.Open(mysql.Open(database.DbURL(database.BuildDBConfig())), &gorm.Config{})
	log.Println("Database mysql connected.")

	if err != nil {
		log.Fatalf("Error when connecting to database: '%v'", err)
	}

	// Migrate the schema from code to database
	err := database.DB.AutoMigrate(&model.User{})
	if err != nil {
		return
	}

	r := route.SetupRouter()
	//running
	err = r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
	if err != nil {
		return
	}
}
