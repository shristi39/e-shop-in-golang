package controllers

import (
	"commerce/internal/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

func (s *Server) Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (s *Server) LoginAuth(c *gin.Context) {

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	session := sessions.Default(c)

	if session.Get("hello") != "world" {
		session.Set("hello", "world")
		session.Save()
	}

	c.JSON(200, gin.H{"hello": session.Get("hello")})

	email := c.PostForm("email")
	password := c.PostForm("password")

	err := models.Login(s.DB, email, password)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": "unsuccessful",
		})
		return
	}

}
