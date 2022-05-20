package controllers

import (
	"fmt"

	// "fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (s *Server) Product(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	// fmt.Println(id)
	product, err := shop.FindProductById(s.DB, Id)

	log.Println(product)
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
	c.HTML(http.StatusOK, "product.html", gin.H{
		"products": product,
		"status":   status,
	})
}

func (s *Server) AddProduct(c *gin.Context) {
	// fmt.Println("i am inside add")

	c.HTML(http.StatusOK, "user.html", gin.H{})
}
func (s *Server) Create(c *gin.Context) {
	// fmt.Println("i am inside create product")

	//err := models.Create(s.DB, "Mouse", 500, "https://images.unsplash.com/photo-1551515300-2d3b7bb80920?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Y29tcHV0ZXIlMjBtb3VzZXxlbnwwfHwwfHw%3D&w=1000&q=80", "controls the motion of a pointer in two dimensions in a graphical user interface")
	sessionA := sessions.DefaultMany(c, "a")
	data := sessionA.Get("user")
	image := c.PostForm("image")
	name := c.PostForm("name")
	description := c.PostForm("Description")
	price := c.PostForm("price")
	email := fmt.Sprintf("%v", data)

	Id, err := strconv.Atoi(price)

	if err != nil {

		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	err = shop.Create(s.DB, name, Id, image, description, email)

	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
	}
	c.Redirect(302, "/")

}

func (s *Server) Delete(c *gin.Context) {
	Id := c.Request.URL.Query().Get("id")
	id, err := strconv.Atoi(Id)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	err = shop.Delete(s.DB, id)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
}
