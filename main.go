package main

import (
	"commerce/internal/controllers"
	"commerce/internal/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	} else {
		log.Info("We are getting the env values")
	}
	db := Connectdb()
	r := models.NewRepository(db)
	server, err := controllers.NewServer(r)
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	}
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	server.RouterHandler()

}

func Connectdb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})

	return db

}
