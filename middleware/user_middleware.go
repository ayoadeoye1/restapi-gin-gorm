package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			secretKey := 
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort() // Stop further processing if unauthorized
			return
		}

		// Set the token claims to the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("claims", claims)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next() // Proceed to the next handler if authorized
	}
}

// func adminMiddleware() gin.HandlerFunc {
//  return func(c *gin.Context) {
//   claims := c.MustGet("claims").(jwt.MapClaims)
//   role := claims["role"].(string)

//   if role != "admin" {
//    c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
//    c.Abort()
//    return
//   }

//   c.Next()
//  }
// }
