package controllers

import (
	"commerce/internal/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) MyCarts(c *gin.Context) {

	products, err := models.FindProductByStatus(s.DB, true)
	// pro := models.Product{}
	// for i, product := range *products {
	// 	if i == 1 {
	// 		pro = product

	// 	}
	// }
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "mycart.html", gin.H{
		"products": products,
	})
}

func (s *Server) UpdateCart(c *gin.Context) {
	// fmt.Println("Hellooooo i am inside update cart")
	id := c.Param("id")
	fmt.Println("........", id)
	Id, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	err = models.Update(s.DB, true, Id)
	if err != nil {
		return
	}
	c.Redirect(301, "/cart")

}

func (s *Server) DeleteCart(c *gin.Context) {
	fmt.Println("Hellooooo i am inside update cart")
	id := c.Param("id")
	fmt.Println("........", id)
	Id, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	err = models.Update(s.DB, false, Id)
	if err != nil {
		return
	}
	c.Redirect(301, "/cart")

}
