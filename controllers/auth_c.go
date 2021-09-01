package controllers

import (
	"fmt"
	"net/http"

	"example.com/api-example/models"
	"github.com/gin-gonic/gin"
)

func (auth *Auth) AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader != "" {
			tokenString := authHeader[len(BEARER_SCHEMA):]
			token, err := auth.AuthService.ValidateToken(tokenString)

			if !token.Valid {
				fmt.Println(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		}
	}
}

func (auth *Auth) Login(c *gin.Context) {
	var loginCredential = struct {
		Username string
		Password string
	}{}

	c.ShouldBind(&loginCredential)

	var user models.Users
	result := auth.Database.Where(
		"username = ? AND password = ?",
		loginCredential.Username,
		loginCredential.Password).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "unauthorized"})
	} else {
		token := auth.AuthService.GenerateToken(user.Username, true)
		c.JSON(http.StatusAccepted, gin.H{"message": "Success", "token": token})
	}

}
