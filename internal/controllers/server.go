package controllers

import (
	"commerce/internal/models"

	"github.com/jinzhu/gorm"
)

type Server struct {
	DB *gorm.DB
	R  models.Store
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
