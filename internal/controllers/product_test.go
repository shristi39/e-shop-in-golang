package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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
	// Item := models.Product{

	// 	Description: "Good",
	// }
	// dummyError := errors.New("dummy errors")
	// testCases := []struct {
	// 	name          string
	// 	setCreateItem func(store *mockdb.MockStore)
	// }{
	// 	{
	// 		name: "Cart_test",
	// 		setCreateItem: func(cstore *mockdb.MockStore) {
	// 			cstore.EXPECT().FindProductById(gomock.Any()).Return(&Item, nil).Times(1)
	// 		},
	// 	},
	// 	{
	// 		name: "Cart_test_Error",
	// 		setCreateItem: func(cstore *mockdb.MockStore) {

	// 			cstore.EXPECT().FindProductById(gomock.Any()).Return(&Item, dummyError).Times(1)
	// 		},
	// 	},
	// }
	// for i := range testCases {
	// 	tc := testCases[i]
	// 	t.Run(tc.name, func(t *testing.T) {
	// 		mockCtrl := gomock.NewController(t)
	// 		defer mockCtrl.Finish()
	// 		cstore := mockdb.NewMockStore(mockCtrl)
	// 		tc.setCreateItem(cstore)

	server := NewTestServer(t, nil)
	fmt.Println(server)
	req, err := http.NewRequest("GET", "/product/:1", nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	server.RouterHandler().ServeHTTP(res, req)

}

func TestDelete(t *testing.T) {

	server := NewTestServer(t, nil)
	fmt.Println(server)
	req, err := http.NewRequest("DELETE", "/product/:1", nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	server.RouterHandler().ServeHTTP(res, req)

}
func TestCreate(t *testing.T) {

	server := NewTestServer(t, nil)
	fmt.Println(server)
	req, err := http.NewRequest("POST", "/create", nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	server.RouterHandler().ServeHTTP(res, req)

}
