package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterhtml(t *testing.T) {

	server := NewTestServer(t, nil)
	fmt.Println(server)
	req, err := http.NewRequest("GET", "/register", nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()

	server.RouterHandler().ServeHTTP(res, req)

}

// func TestRegister(t *testing.T) {
// 	server := NewTestServer(t, nil)
// 	fmt.Println(server)
// 	req, err := http.NewRequest("POST", "/register", nil)
// 	assert.NoError(t, err)
// 	res := httptest.NewRecorder()
// 	server.RouterHandler().ServeHTTP(res, req)

// }

// func TestRegister(t *testing.T) {
// 	Item := models.Product{

// 		Description: "Good",
// 	}
// 	dummyError := errors.New("dummy errors")
// 	testCases := []struct {
// 		name          string
// 		setCreateItem func(store *mockdb.MockStore)
// 	}
// 		func TestProduct(t *testing.T) {
// 		Item := models.Product{

// 			Description: "Good",
// 		}
// 		dummyError := errors.New("dummy errors")
// 		testCases := []struct {
// 			name          string
// 			setCreateItem func(store *mockdb.MockStore)
// 		}{
// 			{
// 				name: "Cart_test",
// 				setCreateItem: func(cstore *mockdb.MockStore) {
// 					cstore.EXPECT().FindProductById(gomock.Any()).Return(&Item, nil).Times(1)
// 				},
// 			},
// 			{
// 				name: "Cart_test_Error",
// 				setCreateItem: func(cstore *mockdb.MockStore) {

// 					cstore.EXPECT().FindProductById(gomock.Any()).Return(&Item, dummyError).Times(1)
// 				},
// 			},
// 		}
// 		for i := range testCases {
// 			tc := testCases[i]
// 			t.Run(tc.name, func(t *testing.T) {
// 				mockCtrl := gomock.NewController(t)
// 				defer mockCtrl.Finish()
// 				cstore := mockdb.NewMockStore(mockCtrl)
// 				tc.setCreateItem(cstore)

// 				server := NewTestServer(t, cstore)
// 				fmt.Println(server)
// 				req, err := http.NewRequest("GET", "/product/1", nil)
// 				assert.NoError(t, err)
// 				res := httptest.NewRecorder()
// 				server.RouterHandler().ServeHTTP(res, req)

// 			})
// 		}
// 			}	}

func TestRegister(t *testing.T) {

	// dummyError := errors.New("dummy errors")
	// testCases := []struct {
	// 	name          string
	// 	body          gin.H
	// 	setCreateItem func(store *mockdb.MockStore)
	// }{
	// 	{
	// 		name: "Login_test",
	// 		body: gin.H{
	// 			"name":     "name",
	// 			"email":    "email",
	// 			"password": "password",
	// 		},

	// 		setCreateItem: func(cstore *mockdb.MockStore) {
	// 			cstore.EXPECT().Register(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// 		},
	// 	},
	// 	{
	// 		name: "Login_test_Error",
	// 		body: gin.H{
	// 			"name":     "name",
	// 			"email":    "email",
	// 			"password": "password",
	// 		},

	// 		setCreateItem: func(cstore *mockdb.MockStore) {
	// 			cstore.EXPECT().Register(gomock.Any(), gomock.Any(), gomock.Any()).Return(dummyError).Times(1)

	// 		},
	// 	},
	// }
	// for i := range testCases {
	// 	tc := testCases[i]
	// 	t.Run(tc.name, func(t *testing.T) {
	// 		mockCtrl := gomock.NewController(t)
	// 		defer mockCtrl.Finish()
	// 		cstore := mockdb.NewMockStore(mockCtrl)
	// 		data, err := json.Marshal(tc.body)
	// 		assert.NoError(t, err)
	// 		tc.setCreateItem(cstore)
	// 		//shop = cstore
	server := NewTestServer(t, nil)
	fmt.Println(server)
	req, err := http.NewRequest("POST", "/register", nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	server.RouterHandler().ServeHTTP(res, req)

}

// 	for i := range testCases {
// 		tc := testCases[i]
// 		t.Run(tc.name, func(t *testing.T) {
// 			mockCtrl := gomock.NewController(t)
// 			// data, err := json.Marshal(tc.body)
// 			// assert.NoError(t, err)
// 			defer mockCtrl.Finish()
// 			cstore := mockdb.NewMockStore(mockCtrl)
// 			tc.setCreateItem(cstore)
// 			server := NewTestServer(t, cstore)
// 			fmt.Println("server", server)
// 			req, err := http.NewRequest("POST", "/register", nil)
// 			fmt.Println("req", req)
// 			r := req.URL.Query()
// 			r.Add("name", "sppp")
// 			r.Add("email", "abc@")
// 			r.Add("password", "abcde")
// 			req.URL.RawQuery = r.Encode()
// 			fmt.Println("aaareq", req)

// 			assert.NoError(t, err)
// 			res := httptest.NewRecorder()
// 			server.RouterHandler().ServeHTTP(res, req)

// 		})
// 	}
// }
