package controllers

import (

	// "fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (s *Server) MyCarts(c *gin.Context) {
	// fmt.Println("i am inside cart")
	// id := c.Param("id")
	// Id, err := strconv.Atoi(id)
	// if err != nil {
	// 	return
	// }

	// err = models.Update(s.DB, true, Id)
	// if err != nil {
	// 	return
	// }

	products, count, sum, err := shop.FindProductByStatus(s.DB, true)

	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "mycart.html", gin.H{
		"products": products,
		"count":    count,
		"sum":      sum,
	})

}

func (s *Server) UpdateCart(c *gin.Context) {
	// fmt.Println("Hellooooo i am inside update cart")
	id := c.Param("id")
	// fmt.Println("........", id)
	Id, err := strconv.Atoi(id)
	sessionA := sessions.DefaultMany(c, "a")
	if sessionA.Get("user") == "Guest" {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	} else {
		c.Redirect(307, "/")
	}
	if err != nil {

		return
	}

	err = shop.Update(s.DB, true, Id)
	if err != nil {
		return
	}
	c.Redirect(302, "/")

}

func (s *Server) Checkout(c *gin.Context) {
	// fmt.Println("Hellooooo i am inside update cart")
	products, _, _, _ := shop.FindProductByStatus(s.DB, true)
	for _, product := range *products {
		shop.Update(s.DB, false, int(product.ID))
	}

	c.Redirect(302, "/")

}
