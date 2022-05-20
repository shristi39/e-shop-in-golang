package controllers

import (
	// "fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (s *Server) Logout(c *gin.Context) {
	// fmt.Println("I am inside logout")
	sessionA := sessions.DefaultMany(c, "a")
	sessionA.Clear()
	sessionA.Set("user", "Guest")
	sessionA.Save()
	// fmt.Println(sessionA.Get("user"))
	c.Redirect(307, "/")

}
