package middleware

import (
	"fmt"
	"go/jwt_auth/intializers"
	"go/jwt_auth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {

	// get the cookie of request

	tokenString, err := c.Cookie("Authorization")

	if err != nil {

		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// decode/validate it

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// expire xcheck

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)

		}

		// find the user token subject

		var user models.User
		query := "SELECT * FROM users WHERE id = ? "
		id := claims["sub"]
		intializers.GetDB().Raw(query, id).Scan(&user)

		// attach to the req

		c.Set("user",user)

		// continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
