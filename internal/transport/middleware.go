package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID string `json:"id"`
	jwt.RegisteredClaims
}

func AuthMiddleware(jwtKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization required"})
				return
			}
			tokenString = authHeader[len("Bearer "):]
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		if token.Method != jwt.SigningMethodHS256 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unexpected signing method"})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
