package main

import (
	"fmt"

	"auth/middleware"
	"auth/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if godotenv.L

	router := gin.Default()

	router.GET("/unprotected", unprotectedRouteHandler);
	router.POST("/auth", authRouteHandler);
	router.GET("/protected", middleware.AuthMiddleware(), protectedRouteHandler);

	router.Run(":1234");
	fmt.Println("Server up and running ...");
}

func unprotectedRouteHandler(c *gin.Context) {
	c.Writer.WriteString("This is an unprotected route");
}

func authRouteHandler(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	fmt.Println(user.username)
}

func protectedRouteHandler(c *gin.Context) {
	c.Writer.WriteString("This is a protected route")
}
