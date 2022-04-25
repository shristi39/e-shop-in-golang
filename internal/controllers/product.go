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

	err := models.Create(s.DB, "speaker", 11000, "https://media.istockphoto.com/photos/digital-background-smart-assistant-picture-id1320780003?s=612x612")

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
