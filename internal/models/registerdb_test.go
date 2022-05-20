package models

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRegisterUserModel(t *testing.T) {
	user := &User{
		Name:     "user name ",
		Email:    "user email",
		Password: "user password",
	}
	server.Mock.ExpectQuery(regexp.QuoteMeta(`INSERT`)).WillReturnRows(sqlmock.NewRows([]string{"Id"}).AddRow(1))
	server.Mock.ExpectCommit()
	server.Mock.ExpectBegin()
	server.Mock.MatchExpectationsInOrder(false)
	err := server.ustore.Register(server.DB, user.Name, user.Email, user.Password)
	fmt.Println()
	if err != nil {
		t.Errorf("The error is creating the item: %v\n", err)
		return
	}
	//err = server.ustore.Register(server.DB, .Name, User.Email, User.Password)
}
func TestLoginUserModel(t *testing.T) {
	user := &User{

		Email:    "user email",
		Password: "user password",
	}
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"Id", "deleted_at", "email", "password"}).AddRow(1, time.Now(), user.Email, user.Password))
	server.Mock.ExpectCommit()
	server.Mock.ExpectBegin()
	server.Mock.MatchExpectationsInOrder(false)
	err := server.ustore.Login(server.DB, user.Email, user.Password)
	fmt.Println()
	if err != nil {
		t.Errorf("The error is creating the item: %v\n", err)
		return
	}
	//err = server.ustore.Register(server.DB, .Name, User.Email, User.Password)
}
