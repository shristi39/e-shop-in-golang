package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogout(t *testing.T) {
	server := NewTestServer(t, nil)
	fmt.Println(server)
	req, err := http.NewRequest("GET", "/logout", nil)
	assert.NoError(t, err)
	res := httptest.NewRecorder()
	server.RouterHandler().ServeHTTP(res, req)
}
