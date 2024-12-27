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
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		token, err := c.Cookie("Authorization")
		if err != nil {
			helper.SendError(c, http.StatusBadRequest, "Authorization Token is missing in the header/cookie", err.Error())
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		tokenString = token
	}

	token, err := verifyToken(tokenString)
	if err != nil {
		helper.SendError(c, http.StatusBadRequest, "Authorization Token coud not be verified", err.Error())
		c.Abort()
		return
	}

	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)
	c.Next()
}
