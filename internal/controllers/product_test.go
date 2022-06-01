package controllers

import (
	"commerce/internal/controllers/mockdb"
	"commerce/internal/models"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {

	server := NewTestServer(t, nil)
	fmt.Println(server)
	req, err := http.NewRequest("GET", "/addproduct", nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	server.RouterHandler().ServeHTTP(res, req)
}

func TestProduct(t *testing.T) {
	Item := models.Product{

		Description: "Good",
	}
	dummyError := errors.New("dummy errors")
	testCases := []struct {
		name          string
		setCreateItem func(store *mockdb.MockStore)
	}{
		{
			name: "Product_test",
			setCreateItem: func(cstore *mockdb.MockStore) {
				cstore.EXPECT().FindProductById(gomock.Any()).Return(&Item, nil).Times(1)
			},
		},
		{
			name: "Product_test_Error",
			setCreateItem: func(cstore *mockdb.MockStore) {

				cstore.EXPECT().FindProductById(gomock.Any()).Return(&Item, dummyError).Times(1)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			cstore := mockdb.NewMockStore(mockCtrl)
			tc.setCreateItem(cstore)

			server := NewTestServer(t, cstore)
			fmt.Println(server)
			req, err := http.NewRequest("GET", "/product/1", nil)
			assert.NoError(t, err)
			res := httptest.NewRecorder()
			server.RouterHandler().ServeHTTP(res, req)

		})
	}
}

func TestDelete(t *testing.T) {

	server := NewTestServer(t, nil)
	fmt.Println(server)
	req, err := http.NewRequest("DELETE", "/product/1", nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	server.RouterHandler().ServeHTTP(res, req)

}

// func TestDelete(t *testing.T) {

// 	server := NewTestServer(t, nil)
// 	fmt.Println(server)
// 	req, err := http.NewRequest("PUT", "delete", nil)
// 	assert.NoError(t, err)
// 	res := httptest.NewRecorder()
// 	server.RouterHandler().ServeHTTP(res, req)

// }

// func TestCreate(t *testing.T) {
// 	Item := models.Product{

// 		Description: "Good",
// 	}
// 	dummyError := errors.New("dummy errors")
// 	testCases := []struct {
// 		name string
// 		// body          gin.H
// 		setCreateItem func(store *mockdb.MockStore)
// 	}{
// 		{
// 			name: "Cart_test",
// 			// body: gin.H{
// 			// 	"name":     "test",
// 			// 	"email":    "email",
// 			// 	"password": "password",
// 			// },
// 			setCreateItem: func(cstore *mockdb.MockStore) {
// 				cstore.EXPECT().FindProductById(gomock.Any()).Return(&Item, nil).Times(1)
// 			},
// 		},
// 		{
// 			name: "Cart_test_Error",
// 			// body: gin.H{
// 			// 	"name":     "test",
// 			// 	"email":    "email",
// 			// 	"password": "password",
// 			// },
// 			setCreateItem: func(cstore *mockdb.MockStore) {

// 				cstore.EXPECT().FindProductById(gomock.Any()).Return(&Item, dummyError).Times(1)
// 			},
// 		},
// 	}
// 	for i := range testCases {
// 		tc := testCases[i]
// 		t.Run(tc.name, func(t *testing.T) {
// 			mockCtrl := gomock.NewController(t)
// 			defer mockCtrl.Finish()
// 			// data, err := json.Marshal(tc.body)
// 			// assert.NoError(t, err)
// 			cstore := mockdb.NewMockStore(mockCtrl)
// 			tc.setCreateItem(cstore)

// 			server := NewTestServer(t, nil)
// 			fmt.Println(server)
// 			req, err := http.NewRequest("POST", "/create", nil)
// 			assert.NoError(t, err)
// 			res := httptest.NewRecorder()
// 			server.RouterHandler().ServeHTTP(res, req)

// 		})
// 	}
// }
// func TestCreate(t *testing.T) {

// 	server := NewTestServer(t, nil)
// 	fmt.Println(server)
// 	req, err := http.NewRequest("POST", "/product", nil)
// 	assert.NoError(t, err)
// 	res := httptest.NewRecorder()
// 	server.RouterHandler().ServeHTTP(res, req)
// }

func TestCreate(t *testing.T) {

	// dummyError := errors.New("dummy errors")
	// testCases := []struct {
	// 	name          string
	// 	body          gin.H
	// 	setCreateItem func(store *mockdb.MockStore)
	// }{

	// 	{
	// 		name: "Login_test",
	// 		body: gin.H{

	// 			"name":        "abcde",
	// 			"price":       "1234",
	// 			"image":       "image",
	// 			"Description": "yhjsdoiuj",
	// 		},

	// 		setCreateItem: func(cstore *mockdb.MockStore) {
	// 			cstore.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// 		},
	// 	},
	// 	{
	// 		name: "Login_test_Error",
	// 		body: gin.H{

	// 			"name":        "abcde",
	// 			"price":       "1234",
	// 			"image":       "image",
	// 			"Description": "yhjsdoiuj",
	// 		},

	// 		setCreateItem: func(cstore *mockdb.MockStore) {
	// 			cstore.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(dummyError).Times(1)

	// 		},
	// 	},
	// }
	// for i := range testCases {
	// 	tc := testCases[i]
	// 	t.Run(tc.name, func(t *testing.T) {
	// 		mockCtrl := gomock.NewController(t)
	// 		defer mockCtrl.Finish()
	// 		cstore := mockdb.NewMockStore(mockCtrl)
	// 		fmt.Println("abcde12345fffffffffff")
	// 		data, err := json.Marshal(tc.body)
	// 		assert.NoError(t, err)
	// 		tc.setCreateItem(cstore)
	// 		//shop = cstore
	server := NewTestServer(t, nil)
	fmt.Println("56789000000000000000")
	fmt.Println(server)
	req, err := http.NewRequest("POST", "/create", nil)
	assert.NoError(t, err)
	fmt.Println("helooooooooooooooooooooooooooooooooooooooooooooooooo")
	res := httptest.NewRecorder()
	fmt.Println("uuuhyrrttttttttttttttttttttttttttttt")
	server.RouterHandler().ServeHTTP(res, req)
	// fmt.Println("weeeeeeeeeeeeeeeeeewewewwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwe")

}
