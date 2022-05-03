package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Name        string
	Price       int
	Image       string
	Status      bool
	Description string
}

func Create(db *gorm.DB, name string, price int, image string, description string) error {

	product := Product{Name: name, Price: price, Image: image, Description: description, Status: false}
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

func FindProductByStatus(db *gorm.DB, status bool) (*[]Product, *int, *int, error) {
	// fmt.Println("Heloo.............................")
	data := []Product{}
	err := db.Model(&data).Where("status = ?", status).Find(&data).Error
	count := 0
	db.Model(&data).Where("status = ?", status).Count(&count)
	if err != nil {
		return nil, nil, nil, err
	}
	sum := 0
	for i, p := range data {
		fmt.Println(i)
		fmt.Println(p.Price)
		sum = sum + p.Price

	}
	fmt.Println(sum)
	return &data, &count, &sum, err

}

func Delete(db *gorm.DB, id int) error {
	Products := Product{}
	err := db.Model(&Products).Where("id=?", id).Delete(&Products).Error
	if err != nil {
		return err
	}
	return nil

}

func Update(db *gorm.DB, status bool, id int) error {
	fmt.Println("Helllo I am inside Update")
	data := &Product{}
	data, err := FindProductById(db, id)
	// fmt.Println(id)
	fmt.Println(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	data.Status = status
	err = db.Model(&data).Save(&data).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
