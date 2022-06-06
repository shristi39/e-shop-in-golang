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
	Id, _ := strconv.Atoi(id)
	// if err != nil {
	// 	return
	// }

	// fmt.Println(id)
	Product := &product{}
	err := c.ShouldBindJSON(&Product)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		// fmt.Println("i am inside home")
		return
	}

	product, err := s.R.FindProductById(Id)

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
func (s *Server) RestProduct(c *gin.Context) {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	// Product := &product{}

	product, err := s.R.FindProductById(Id)
	fmt.Println(id)
	log.Println(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "this is error",
		})
		return

	}

	//	log.Println(Product)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "this is an error",
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"products": product,
	})
}

func (s *Server) AddProduct(c *gin.Context) {
	// fmt.Println("i am inside add")

	c.HTML(http.StatusOK, "user.html", gin.H{})
}
func (s *Server) Create(c *gin.Context) {

	// sessionA := sessions.DefaultMany(c, "a")
	// data := sessionA.Get("user")
	image := c.PostForm("image")
	name := c.PostForm("name")
	description := c.PostForm("Description")
	price := c.PostForm("price")

	Id, err := strconv.Atoi(price)
	fmt.Println("Hello from create")
	if err != nil {

		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	err = s.R.Create(name, Id, image, description)

	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
	}
	c.Redirect(302, "/")

}

func (s *Server) RestCreate(c *gin.Context) {
	Product := &product{}
	err := c.ShouldBindJSON(&Product)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})

		return
	}
	err = s.R.Create(Product.Name, Product.Price, Product.Image, Product.Description)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfull",
	})
}

func (s *Server) Delete(c *gin.Context) {
	Id := c.Param("id")
	id, err := strconv.Atoi(Id)
	if err != nil {
		return
	}
	err = s.R.Delete(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
}

func (s *Server) RestDelete(c *gin.Context) {
	Id := c.Param("id")
	fmt.Println(Id)
	id, err := strconv.Atoi(Id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = s.R.Delete(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "succesful",
	})
}
