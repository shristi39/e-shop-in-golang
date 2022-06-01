package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

// var ecommerce = models.InitializeRegisterModels()

func (s *Server) Login(c *gin.Context) {
	// fmt.Println("I am inside login")

	c.HTML(http.StatusOK, "login.html", gin.H{})

}

func (s *Server) LoginAuth(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")
	err := s.R.Login(email, password)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": "unsuccessful",
		})

	} else {
		sessionA := sessions.DefaultMany(c, "a")
		sessionA.Set("user", email)
		sessionA.Save()
		sessionA.Clear()
		c.Redirect(302, "/")
	}

}

// func (s *Server) RestLoginAuth(c *gin.Context) {

// 	User := &user{}
// 	err := c.ShouldBindJSON(&User)
// 	if err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"error": "error",
// 		})
// 		return
// 	}
// 	err = s.R.Login(User.Email, User.Password)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "error",
// 		})
// 		return

// 	} else {
// 		c.JSON(http.StatusOK, gin.H{
// 			"login": "successfull",
// 		})
// 	}

// }
