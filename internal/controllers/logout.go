package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (s *Server) Logout(c *gin.Context) {
	sessionA := sessions.DefaultMany(c, "a")
	sessionA.Clear()
	sessionA.Set("user", "Guest")
	sessionA.Save()
	c.Redirect(307, "/")

}
