package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type TestServer struct {
	DB    *gorm.DB
	Mock  sqlmock.Sqlmock
	Store Store
}

var server TestServer

func TestMain(m *testing.M) {
	Database()
	os.Exit(m.Run())
}

func Database() {
	var err error
	var db *sql.DB
	_ = os.Setenv("TestDbDriver", "postgres")
	TestDbDriver := os.Getenv("TestDbDriver")
	db, server.Mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	server.DB, err = gorm.Open(TestDbDriver, db)
	if err != nil {
		fmt.Printf("Cannot connect to mock %s database\n", TestDbDriver)
		log.Fatal("This is the error:", err)
	}
	server.Store = NewRepository(server.DB)
	// server.ustore = InitializeRegisterModels()
	// server.store = InitializeProductModels()
}

func TestCreateProductModel(t *testing.T) {
	Product := &Product{
		Name:        "product name ",
		Price:       66677,
		Image:       "product image",
		Email:       " user email ",
		Description: "product description",
	}
	data := server.Mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT`)).WillReturnRows(sqlmock.NewRows([]string{"Id"}).AddRow(1))
	fmt.Println(data)
	server.Mock.ExpectCommit()
	server.Mock.ExpectBegin()
	server.Mock.MatchExpectationsInOrder(false)
	err := server.Store.Create(Product.Name, Product.Price, Product.Image, Product.Email, Product.Description)
	fmt.Println()
	if err != nil {
		t.Errorf("The error is creating the item: %v\n", err)
		return
	}
	// err = server.Store.Create(Product.Name, Product.Price, Product.Image, Product.Email, Product.Description)
}
func TestFindProductById(t *testing.T) {
	Item := &Product{
		Description: "Hello",
	}
	Item.ID = 1
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "deleted_at", "description"}).
		AddRow(Item.ID, time.Now(), Item.Description))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	saved, err := server.Store.FindProductById(int(Item.ID))
	if err != nil {
		t.Errorf("This is error getting all items: %v", err)
		return
	}
	_, err = server.Store.FindProductById(int(Item.ID))
	if err != nil {
		assert.Error(t, err)
		return
	}
	assert.Equal(t, saved.Description, Item.Description)
}

func TestFindProductByStatus(t *testing.T) {
	Item := &Product{
		Description: "Hello",
		Status:      false,
	}
	Item.ID = 1
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "description", "status"}).
		AddRow(1, Item.Description, Item.Status))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	saved, _, _, err := server.Store.FindProductByStatus(bool(Item.Status))
	if err != nil {
		t.Errorf("this is the error getting the item: %v\n", err)
		return
	}
	_, _, _, err = server.Store.FindProductByStatus(bool(Item.Status))
	if err != nil {
		assert.Error(t, err)
		return
	}
	assert.Equal(t, saved, Item.Description)
}
func TestDeleteProductModel(t *testing.T) {
	Item := &Product{
		Description: "Delete",
		Status:      true,
	}
	Item.ID = 1
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "description", "deleted_at", "status"}).AddRow(Item.ID, Item.Description, time.Now(), Item.Status))
	server.Mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).WillReturnResult(sqlmock.NewResult(0, 1))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	err := server.Store.Delete(int(Item.ID))
	if err != nil {
		t.Errorf("This is error deleting the item:  %v", err)
		return
	}
	err = server.Store.Delete(int(Item.ID))
	if err != nil {
		assert.Error(t, err)
		return
	}
}

func TestUpdateProductModel(t *testing.T) {
	Item := &Product{
		Description: "Update",
		Status:      true,
	}
	Item.ID = 1
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(`UPDATE`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "description", "deleted_at", "status"}).AddRow(Item.ID, Item.Description, time.Now(), Item.Status))
	server.Mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).WillReturnResult(sqlmock.NewResult(0, 1))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	err := server.Store.Delete(int(Item.ID))
	if err != nil {
		t.Errorf("This is error update the item:  %v", err)
		return
	}
	err = server.Store.Update(bool(Item.Status), int(Item.ID))

	if err != nil {
		assert.Error(t, err)
		return
	}
}

func TestFindAllProduct(t *testing.T) {
	Item := &Product{
		Description: "Hello",
		Status:      false,
	}
	Item.ID = 1
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "description", "status"}).
		AddRow(1, Item.Description, Item.Status))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	saved, err := server.Store.FindAllProduct()
	if err != nil {
		t.Errorf("this is the error getting the item: %v\n", err)
		return
	}
	_, err = server.Store.FindAllProduct()
	if err != nil {
		assert.Error(t, err)
		return
	}
	assert.Equal(t, saved, Item.Description)
}
