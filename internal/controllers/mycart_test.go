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

func TestMyCarts(t *testing.T) {
	Item := []models.Product{
		{
			Description: "Good",
		},
	}
	dummyError := errors.New("dummy errors")
	testCases := []struct {
		name          string
		setCreateItem func(store *mockdb.MockStore)
	}{
		{
			name: "Cart_test",
			setCreateItem: func(cstore *mockdb.MockStore) {
				cstore.EXPECT().FindProductByStatus(gomock.Any()).Return(&Item, intptr(1), intptr(1), nil).Times(1)
			},
		},
		{
			name: "Cart_test_Error",
			setCreateItem: func(cstore *mockdb.MockStore) {

				cstore.EXPECT().FindProductByStatus(gomock.Any()).Return(&Item, intptr(1), intptr(1), dummyError).Times(1)
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
			//shop = cstore
			server := NewTestServer(t, cstore)
			fmt.Println(server)
			req, err := http.NewRequest("GET", "/mycart", nil)
			assert.NoError(t, err)
			res := httptest.NewRecorder()
			server.RouterHandler().ServeHTTP(res, req)

		})
	}
}

func TestCheckout(t *testing.T) {
	Item := []models.Product{
		{
			Description: "Good",
		},
	}
	dummyError := errors.New("dummy errors")
	testCases := []struct {
		name          string
		setCreateItem func(store *mockdb.MockStore)
	}{
		{
			name: "Checkout_test",
			setCreateItem: func(cstore *mockdb.MockStore) {

				cstore.EXPECT().FindProductByStatus(gomock.Any()).Return(&Item, intptr(1), intptr(1), nil).Times(1)
				cstore.EXPECT().Update(gomock.Any(), gomock.Any()).Return(dummyError).Times(1)
			},
		},
		{
			name: "Checkout_test_Error",
			setCreateItem: func(cstore *mockdb.MockStore) {

				cstore.EXPECT().FindProductByStatus(gomock.Any()).Return(&Item, intptr(1), intptr(1), dummyError).Times(1)
				cstore.EXPECT().Update(gomock.Any(), gomock.Any()).Return(dummyError).Times(1)
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
			//shop = cstore
			server := NewTestServer(t, cstore)
			fmt.Println(server)
			req, err := http.NewRequest("GET", "/checkout", nil)
			assert.NoError(t, err)
			res := httptest.NewRecorder()
			server.RouterHandler().ServeHTTP(res, req)

		})
	}
}

func TestUpdateCart(t *testing.T) {
	server := NewTestServer(t, nil)
	fmt.Println(server)
	req, err := http.NewRequest("GET", "/product/:1/cart", nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	server.RouterHandler().ServeHTTP(res, req)
}
