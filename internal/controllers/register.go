package controllers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func (s *Server) Registerhtml(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", gin.H{})

}

func (s *Server) Register(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	namepattern := `^[a-zA-Z]{4,}(?: [a-zA-Z]+){0,2}$`
	emailpattern := `^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`

	match1, _ := regexp.MatchString(namepattern, name)
	match2, _ := regexp.MatchString(emailpattern, email)

	secure := true
	tests := []string{".{7,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
	for _, test := range tests {
		t, _ := regexp.MatchString(test, password)
		if !t {
			secure = false
			break
		}
	}

	fmt.Println(match1)
	if !match1 {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"namemessage": "invalid name",
		})
		return
	}

	fmt.Println(match2)
	if !match2 {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"emailmessage": "Invalid email",
		})
		return
	}

	fmt.Println(secure)
	if !secure {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"passwordmessage": "at least 8 characters at least one upper case character at least one lower case  at least one number at least one unique character",
		})
		return
	}

	err := ecommerce.Register(s.DB, name, email, password)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": "unsuccessful",
		})
		return
	}
	c.Redirect(301, "/")
	// c.HTML(http.StatusOK, "/", gin.H{
	// "message": "successfully register",
	// })
}
