package controllers

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func RouterHandler(s *Server) {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	sessionNames := []string{"a"}
	fmt.Println(sessionNames)
	r.Use(sessions.SessionsMany(sessionNames, store))

	r.LoadHTMLGlob("static/*.html") //we can load from anywhere
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "WELCOME",
		})
	})

	// r.GET("/cart/remove/:id", s.DeleteCart)
	//r.GET("/cart/update/:id", s.UpdateCart)
	r.GET("/", s.Home)
	r.GET("/logout", s.Logout)
	r.GET("/login", s.Login)
	r.POST("/loginauth", s.LoginAuth)
	// r.GET("/login/check", s.CheckLogin)
	r.GET("/register", s.Registerhtml)
	r.POST("/register", s.Register)
	// r.GET("/signup", s.Signup)
	// r.POST("/signuppost", s.SignupPost)
	r.GET("/product/:id", s.Product)
	r.GET("/product/:id/cart", s.MyCarts)
	r.GET("/checkout", s.Checkout)
	r.POST("/checkout")
	r.GET("/mycart", s.MyCarts)
	r.POST("/product", s.Create)
	r.PUT("/product/:id", s.Delete)
	r.DELETE("/product/:id", s.Delete)
	//r.GET("/hello", s.Check)

	r.Run(":8000")
}
