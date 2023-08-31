package main

import (
	"fmt"

	"auth/middleware"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/unprotected", unprotectedRouteHandler);
	router.POST("/auth", authRouteHandler);
	router.GET("/protected", middleware.AuthMiddleware(), protectedRouteHandler);

	router.Run(":1234");
	fmt.Println("Server up and running ...");
}

func unprotectedRouteHandler(c *gin.Context) {
	c.Writer.WriteString("This is the unprotected route");
}

func authRouteHandler(c *gin.Context) {
	
}

func protectedRouteHandler(c *gin.Context) {

}
