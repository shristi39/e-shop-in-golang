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
	User        bool
	Email       string
}

//  here we make the interface inside which all the functions are defined
// type ProductModels interface {
// 	Create(db *gorm.DB, name string, price int, image string, description string, email string) error
// 	FindAllProduct(db *gorm.DB) ([]Product, error)
// 	FindProductById(db *gorm.DB, id int) (*Product, error)
// 	FindProductByStatus(db *gorm.DB, status bool) (*[]Product, *int, *int, error)
// 	Delete(db *gorm.DB, id int) error
// 	Update(db *gorm.DB, status bool, id int) error
// }

//  here we initialize the interface which has been defined above  there
// func InitializeProductModels() ProductModels {
// 	return &databases{}
// }

func (d *Repository) Create(name string, price int, image string, description string, email string) error {

	product := Product{Name: name, Price: price, Image: image, Description: description, Status: false, Email: email}
	err := d.DB.Model(&product).Create(&product).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (d *Repository) FindAllProduct() ([]Product, error) {
	data := []Product{}
	err := d.DB.Model(&Product{}).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (d *Repository) FindProductById(id int) (*Product, error) {
	data := Product{}
	err := d.DB.Model(&data).Where("id = ?", id).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (d *Repository) FindProductByStatus(status bool) (*[]Product, *int, *int, error) {
	// fmt.Println("Heloo.............................")
	data := []Product{}
	err := d.DB.Model(&data).Where("status = ?", status).Find(&data).Error
	count := 0
	d.DB.Model(&data).Where("status = ?", status).Count(&count)
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

func (d *Repository) Delete(id int) error {
	Products := Product{}
	err := d.DB.Model(&Products).Where("id=?", id).Delete(&Products).Error
	if err != nil {
		return err
	}
	return nil

}

func (d *Repository) Update(status bool, id int) error {
	fmt.Println("Helllo I am inside Update")
	data := &Product{}
	data, err := d.FindProductById(id)
	// fmt.Println(id)
	fmt.Println(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	data.Status = status
	err = d.DB.Model(&data).Save(&data).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
