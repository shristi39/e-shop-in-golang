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

func TestRegister(t *testing.T) {
	// Item := []models.Product{
	// 	{
	// 		Description: "Good",
	// 	},
	// }
	// dummyError := errors.New("dummy errors")
	// testCases := []struct {
	// 	name          string
	// 	setCreateItem func(store *mockdb.MockStore)
	// }{
	// 	{
	// 		name: "Register_test",
	// 		setCreateItem: func(cstore *mockdb.MockStore) {
	// 			cstore.EXPECT().Register(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
	// 		},
	// 	},
	// 	{
	// 		name: "Register_test_Error",
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
	// 		tc.setCreateItem(cstore)
	//shop = cstore
	server := NewTestServer(t, nil)
	fmt.Println(server)
	req, err := http.NewRequest("POST", "/register", nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	server.RouterHandler().ServeHTTP(res, req)

}
