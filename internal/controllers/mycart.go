package controllers

import (
	"commerce/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) MyCarts(c *gin.Context) {
	products, err := models.FindAllProduct(s.DB)
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
