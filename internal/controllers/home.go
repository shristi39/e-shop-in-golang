package controllers

import (
	"commerce/internal/models"
	"fmt"

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (s *Server) Home(c *gin.Context) {
	products, err := models.FindAllProduct(s.DB)
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
	fmt.Println(user)
	if user == "Guest" {
		status = false
	} else {
		status = true
	}
	//email := sessionA.Get("email")
	// fmt.Println(user)
	c.HTML(http.StatusOK, "home.html", gin.H{
		"products": products,
		"user":     user,
		"status":   status,
	})
}
