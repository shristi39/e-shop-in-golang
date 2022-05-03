package main

import (
	"commerce/internal/controllers"
	"commerce/internal/models"

	"github.com/jinzhu/gorm"
)

func main() {
	db := Connectdb()
	server := &controllers.Server{DB: db}
	controllers.RouterHandler(server)
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
