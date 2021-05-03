package middlewares

import (
	"api-orders/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

const ACCESS_KEY = "Access-Token"

func responseWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"message": message})
}

// Authenticate fetches user details from token
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		requiredToken := c.Request.Header[ACCESS_KEY]

		fmt.Println(c.Request.Header)

		if len(requiredToken) == 0 {
			responseWithError(c, 403, "Please login to your account")
			return
		}

		userID, _ := services.DecodeToken(requiredToken[0])

		result, err := services.GetAndValidateUser(userID)

		if result.UserId == "" {
			responseWithError(c, 404, "User account not found")
			return
		}

		if err != nil {
			responseWithError(c, 500, "Something went wrong giving you access")
			return
		}

		c.Set("User", result)

		c.Next()
	}
}
