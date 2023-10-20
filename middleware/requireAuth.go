package middlewre

import (
	"fmt"
	"money-management/initializer"
	"money-management/internals/user"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRETKEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		claim := claims["MapClaims"]
		exp := claim.(map[string]interface{})

		//check the exp
		if float64(time.Now().Unix()) > exp["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user wirh token sub
		var user user.User
		initializer.DB.Find(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//attach to req
		c.Set("user_id", int(user.ID))
		//continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
