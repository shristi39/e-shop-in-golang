package main

import (
	"commerce/internal/controllers"
	"commerce/internal/models"

	"github.com/jinzhu/gorm"
)

func main() {
	db := Connectdb()
	r := models.NewRepository(db)
	server := &controllers.Server{DB: db, R: r}
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	server.RouterHandler()

	// mux := http.NewServeMux()

	// fs := http.FileServer(http.Dir("./static"))
	// mux.Handle("/static/", http.StripPrefix("/static/", fs))
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
