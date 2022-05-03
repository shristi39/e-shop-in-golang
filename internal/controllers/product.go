package controllers

import (
	"commerce/internal/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) Product(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	fmt.Println(id)
	product, err := models.FindProductById(s.DB, Id)

	log.Println(product)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "product.html", gin.H{
		"products": product,
	})
}
func (s *Server) Create(c *gin.Context) {

	err := models.Create(s.DB, "Mouse", 500, "https://images.unsplash.com/photo-1551515300-2d3b7bb80920?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Y29tcHV0ZXIlMjBtb3VzZXxlbnwwfHwwfHw%3D&w=1000&q=80", "controls the motion of a pointer in two dimensions in a graphical user interface")

	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}

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
	err = models.Delete(s.DB, id)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
}
