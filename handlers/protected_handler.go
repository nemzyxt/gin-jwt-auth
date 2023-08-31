package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ProtectedRouteHandler(c *gin.Context) {
	username := c.GetString("username")
	c.Writer.WriteString(fmt.Sprintf("Access Granted for: %s", username))
}
