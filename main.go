/*
	Author 		  : Nemuel Wainaina
	Creation date : 31 Aug 2023
	Last updated  : 31 Aug 2023
*/

package main

import (
	"auth/handlers"
	"auth/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/unprotected", handlers.UnprotectedRouteHandler);
	router.POST("/auth", handlers.AuthRouteHandler);
	router.GET("/protected", middleware.AuthMiddleware(), handlers.ProtectedRouteHandler);

	router.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")));
}
