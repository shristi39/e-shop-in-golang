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
	"github.com/stretchr/testify/require"
)

func intptr(i int) *int {
	return &i
}

func NewTestServer(t *testing.T, store models.Store) *Server {
	server, err := NewServer(store)
	require.NoError(t, err)
	return server
}

func TestHome(t *testing.T) {
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
			name: "Home_test",
			setCreateItem: func(cstore *mockdb.MockStore) {
				cstore.EXPECT().FindAllProduct().Return(Item, nil).Times(1)
				cstore.EXPECT().FindProductByStatus(gomock.Any()).Return(&Item, intptr(1), intptr(1), nil).Times(1)
			},
		},
		{
			name: "Home_test_Error",
			setCreateItem: func(cstore *mockdb.MockStore) {
				cstore.EXPECT().FindAllProduct().Return(Item, dummyError).Times(1)
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
			req, err := http.NewRequest("GET", "/", nil)
			assert.NoError(t, err)
			res := httptest.NewRecorder()
			server.RouterHandler().ServeHTTP(res, req)

		})
	}
}
