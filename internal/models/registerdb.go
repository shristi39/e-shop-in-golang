package models

import (
	"crypto/sha1"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type userbase struct {
}

type UserModels interface {
	Register(db *gorm.DB, name string, email string, password string) error
	Login(db *gorm.DB, email string, password string) error
}

func InitializeRegisterModels() UserModels {
	return &userbase{}
}

func Hash(pwd string) string {
	h := sha1.New()
	h.Write([]byte(pwd))
	bs := h.Sum(nil)
	pass := string(bs[:])
	return pass
}

func (d *userbase) Register(db *gorm.DB, name string, email string, password string) error {
	Hashvalue := Hash(password)
	fmt.Println("......login create .......", Hashvalue)
	// Base64(password)
	user := User{Name: name, Email: email, Password: Hashvalue}
	err := db.Model(&user).Where("email = ?", email).Find(&user).Error
	fmt.Println(err)
	if err != nil {
		db.Model(&user).Create(&user)
		return nil
	}
	return errors.New("email not exist")
}

func (d *userbase) Login(db *gorm.DB, email string, password string) error {
	Hashvalue := Hash(password)
	fmt.Println("......................Auth............", Hashvalue)
	user := User{}
	err := db.Model(&user).Where("email = ?", email).Where("password=?", Hashvalue).Find(&user).Error

	if err != nil {
		return err
	}
	return nil
}
