package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			fmt.Println(err.Err)

			c.AbortWithStatusJSON(400, gin.H{
				"Error": err.Error(),
			})
		}
	}
}
