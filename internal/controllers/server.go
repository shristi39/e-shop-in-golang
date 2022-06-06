package controllers

import (
	"commerce/internal/models"

	"github.com/jinzhu/gorm"
)

type Server struct {
	DB *gorm.DB
	R  models.Store
}

type user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       int    `json:"price"`

	Email string `json:"email"`
}

func NewServer(store models.Store) (*Server, error) {
	server := &Server{
		R: store,
	}
	// server.RouterHandler()
	return server, nil
}

// type Store interface {
// 	models.ProductModels
// }

// func NewBaseHandler(pm models.ProductModels, um models.UserModels) *Server {
// 	return &Server{
// 		shop:      pm,
// 		ecommerce: um,
// 	}
// // }
