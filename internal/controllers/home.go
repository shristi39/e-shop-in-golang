package controllers

import (
	"commerce/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Home(c *gin.Context) {
	products, err := models.FindAllProduct(s.DB)
	//p := append(*products, models.Product{Name: "test", Price: 235}
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"products": products,
	})
}
