package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

// var ecommerce = models.InitializeRegisterModels()

// when the user login

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
	// Email       emailResponse `json:"email"`
}

func (s *Server) loginUser(c *gin.Context) {
	// to insert the the login user request in the var
	var req loginUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, (err))
		return
	}
	// fmt.Println("paylaod ", req)
	err := s.R.Login(req.Email, req.Password)
	if err != nil {

		c.JSON(http.StatusInternalServerError, (err))

		return

	}
	fmt.Println("error ", err)
	token, err := s.TokenMaker.GenerateToken(req.Email)
	if err != nil {
		c.JSON(http.StatusForbidden, (err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: token,
		// User:        newUserResponse(user),
	}
	c.JSON(http.StatusOK, rsp)
}

func (s *Server) Login(c *gin.Context) {
	//email := c.PostForm("email")
	// password := c.PostForm("password")
	// token, err := s.TokenMaker.CreateToken(email, 15*time.Minute)
	// if err != nil {
	// 	c.HTML(http.StatusOK, "error.html", gin.H{
	// 		"message": "token create error",
	// 	})
	// }
	// fmt.Println("TOKEN :: ", token)
	// err = s.R.Login(email, password)
	// if err != nil {
	// 	c.HTML(http.StatusOK, "error.html", gin.H{
	// 		"message": "unsuccessful",
	// 	})
	// 	// fmt.Println("I am inside login")

	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (s *Server) RestLogin(c *gin.Context) {
	// fmt.Println("I am inside login")

	c.JSON(http.StatusOK, gin.H{})

}

func (s *Server) LoginAuth(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")
	// token, err := s.TokenMaker.CreateToken(email, 15*time.Minute)
	// if err != nil {
	// 	c.HTML(http.StatusOK, "error.html", gin.H{
	// 		"message": "token create error",
	// 	})
	// }
	// fmt.Println("TOKEN :: ", token)
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
