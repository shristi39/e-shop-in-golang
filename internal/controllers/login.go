package controllers

import (
	"commerce/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

func (s *Server) Login(c *gin.Context) {
	fmt.Println("I am inside login")

	c.HTML(http.StatusOK, "login.html", gin.H{})

}

func (s *Server) CheckLogin(c *gin.Context) {
	// fmt.Println("I am inside CHeck Login")
	sessionA := sessions.DefaultMany(c, "a")
	if sessionA.Get("user") == "Guest" {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	} else {
		c.Redirect(301, "/")
	}
}

func (s *Server) LoginAuth(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")
	fmt.Println("i am inside login")
	err := models.Login(s.DB, email, password)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": "unsuccessful",
		})

	} else {
		sessionA := sessions.DefaultMany(c, "a")
		sessionA.Set("user", email)
		sessionA.Save()
		c.Redirect(301, "/")
	}

}
