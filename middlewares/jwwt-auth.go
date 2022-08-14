package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/DMonkey83/go-gin-gorm/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, BearerSchema) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error message": "Not authorized"})
		} else {

			tokenString := authHeader[len(BearerSchema):]
			token, err := service.NewJWTService().ValidateToken(tokenString)

			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				log.Println("Claims[Name]: ", claims["name"])
				log.Println("Claims[Admin]: ", claims["admin"])
				log.Println("Claims[Issuer]: ", claims["iss"])
				log.Println("Claims[IssuedAt]: ", claims["iat"])
				log.Println("Claims[ExpiresAt]: ", claims["exp"])
			} else {
				log.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}
}
