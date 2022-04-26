package controllers

import (
	"github.com/gin-gonic/gin"
)

func RouterHandler(s *Server) {
	r := gin.Default()
	r.LoadHTMLGlob("static/*.html") //we can load from anywhere
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "WELCOME",
		})
	})
	r.GET("/remove/cart/:id", s.DeleteCart)
	r.GET("/update/cart/:id", s.UpdateCart)
	r.GET("/", s.Home)
	r.GET("/login", s.Login)
	r.GET("/login/Auth", s.LoginAuth)
	r.GET("/register", s.Registerhtml)
	r.POST("/register/create", s.Register)
	// r.GET("/signup", s.Signup)
	// r.POST("/signuppost", s.SignupPost)
	r.GET("/product/:id", s.Product)
	r.GET("/cart", s.MyCarts)
	r.POST("/create", s.Create)
	r.DELETE("/delete", s.Delete)

	r.Run(":7002")
}
