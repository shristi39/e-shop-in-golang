package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Name   string
	Price  int
	Image  string
	Status bool
}

func Create(db *gorm.DB, name string, price int, image string) error {

	product := Product{Name: name, Price: price, Image: image, Status: false}
	err := db.Model(&product).Create(&product).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func FindAllProduct(db *gorm.DB) (*[]Product, error) {
	data := []Product{}
	err := db.Model(&Product{}).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func FindProductById(db *gorm.DB, id int) (*Product, error) {
	data := Product{}
	err := db.Model(&data).Where("id = ?", id).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func Delete(db *gorm.DB, id int) error {
	Products := Product{}
	err := db.Model(&Products).Where("id=?", id).Delete(&Products).Error
	if err != nil {
		return err
	}
	return nil

}

func Update(db *gorm.DB, name string, price int, image string) error {

	product := Product{Name: name, Price: price, Image: image, Status: true}
	err := db.Model(&product).Update(&product).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
