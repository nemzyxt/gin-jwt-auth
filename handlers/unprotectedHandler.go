package handlers

import (
	"github.com/gin-gonic/gin"
)

func UnprotectedRouteHandler(c *gin.Context) {
	c.Writer.WriteString("This is an unprotected route");
}
