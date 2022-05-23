package models

import "github.com/jinzhu/gorm"

type Store interface {
	Create(name string, Price int, image string, description string, email string) error
	FindAllProduct() ([]Product, error)
	FindProductById(id int) (*Product, error)
	FindProductByStatus(status bool) (*[]Product, *int, *int, error)
	Delete(id int) error
	Update(status bool, id int) error
	Register(name string, email string, password string) error
	Login(email string, password string) error
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Store {
	return &Repository{DB: db}
}
