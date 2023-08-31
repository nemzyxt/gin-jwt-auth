package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/unprotected", unprotectedRouteHandler);
	router.GET("/protected", protectedRouteHandler);

	router.Run(":1234");
	fmt.Println("Server up and running ...");
}

func unprotectedRouteHandler(c *gin.Context) {
	c.Writer.WriteString("This is the unprotected route");
}

func protectedRouteHandler(c *gin.Context) {

}
