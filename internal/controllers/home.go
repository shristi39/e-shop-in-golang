package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// var shop = models.InitializeProductModels()

func (s *Server) Home(c *gin.Context) {
	s.R.FindProductByStatus(true)
	fmt.Println("i am inside home")
	_, count, _, _ := s.R.FindProductByStatus(true)
	products, err := s.R.FindAllProduct()

	//p := append(*products, models.Product{Name: "test", Price: 235}
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	var status bool
	sessionA := sessions.DefaultMany(c, "a")
	user := sessionA.Get("user")
	// fmt.Println(user)
	if user == "Guest" {
		status = false
	} else {
		status = true
	}
	for _, u := range products {
		u.User = status
	}
	//email := sessionA.Get("email")
	// fmt.Println(user)
	c.HTML(http.StatusOK, "home.html", gin.H{
		"products": products,
		"user":     user,
		"status":   status,
		"count":    count,
	})
}
