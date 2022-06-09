package controllers

import (
	"commerce/internal/controllers/auth"
	"commerce/internal/models"

	"github.com/jinzhu/gorm"
)

type Server struct {
	DB         *gorm.DB
	R          models.Store
	TokenMaker *auth.JwtWrapper
}

// type LoginCredentials struct {
// 	Email    string `form:"email"`
// 	Password string `form:"password"`
// }

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
	tokenMaker := auth.NewJWT()

	server := &Server{
		R:          store,
		TokenMaker: tokenMaker,
	}
	// server.RouterHandler()
	return server, nil
}

// type LoginService interface {
// 	LoginUser(email string, password string) bool
// }
// type loginInformation struct {
// 	email    string
// 	password string
// }

// func StaticLoginService() LoginService {
// 	return &loginInformation{
// 		email:    "admin@wesionary.team",
// 		password: "admin",
// 	}
// }
// func (info *loginInformation) LoginUser(email string, password string) bool {
// 	return info.email == email && info.password == password
// }
