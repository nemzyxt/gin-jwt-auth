package handlers

import (
	"net/http"
	"os"
	"time"

	"auth/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthRouteHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		return
	}

	if user.Username != "admin" && user.Password != "secret" {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid credentials"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour*1).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error generating token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token":tokenString})
}
