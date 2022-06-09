package controllers

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// var r *gin.Engine

// r.LoadHTMLGlob("static/*.html")
//var static = r.Static("/static", "./static")

func (s *Server) RouterHandler() *gin.Engine {
	r := gin.Default()
	//	r.SetTrustedProxies(nil)
	//gin.SetMode(gin.ReleaseMode)
	store := cookie.NewStore([]byte("secret"))
	sessionNames := []string{"a"}
	fmt.Println(sessionNames)
	r.Use(sessions.SessionsMany(sessionNames, store))

	r.LoadHTMLGlob("static/*.html") //we can load from anywhere
	r.Static("/static", "./static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "WELCOME",
		})
	})
	// r.GET("/rest/", s.RestHome)
	r.GET("/", s.Home)
	r.POST("/user/login", s.loginUser)
	r.GET("/logout", s.Logout)
	//r.GET("/login", s.Login)
	//r.POST("/rest/login", s.RestLogin)
	r.POST("/loginauth", s.LoginAuth)
	// r.POST("/rest/loginauth/", s.RestLoginAuth)

	r.GET("/register", s.Registerhtml)
	r.POST("/register", s.Register)
	r.GET("/product/:id", s.Product)
	r.GET("/product/:id/cart", s.UpdateCart)
	r.GET("/checkout", s.Checkout)
	r.GET("/mycart", s.MyCarts)
	r.GET("/addproduct", s.AddProduct)
	r.POST("/create", s.Create)
	r.DELETE("/product/:id", s.Delete)
	// r.POST("/create", s.CreateToken)
	// rest API

	r.GET("/rest/product/:id", s.RestProduct)
	r.POST("/rest/create", s.RestCreate)
	r.PUT("/rest/product/:id", s.RestDelete)
	r.GET("/rest/checkout", s.RestCheckout)
	// r.POST("/rest/product/:id", s.RestLoginAuth)

	//  sabaie rest api ko yeha hunxa ,,,,,,,,login ko register product id ko get post ra put matra hunxa and checkout matra ho chaiyene ab

	r.Run(":8001")

	return r
}
