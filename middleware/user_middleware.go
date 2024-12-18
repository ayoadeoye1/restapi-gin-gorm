package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ayoadeoye1/restapi-gin-gorm/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func verifyToken(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func UserAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization") //c.Cookie()
	// if err != nil {
	// 	fmt.Println("Token missing in cookie")
	// 	// c.Redirect(http.StatusSeeOther, "/login")
	// 	// c.Abort()
	// 	helper.SendError(c, http.StatusBadRequest, "Authorization Token is missing in the header/cookie")
	// 	return
	// }

	token, err := verifyToken(tokenString)
	if err != nil {
		helper.SendError(c, http.StatusBadRequest, "Authorization Token coud not be verified", err.Error())
		c.Abort()
		return
	}

	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)
	c.Next()
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
