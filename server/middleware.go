package server

import (
	"all_inclusive_app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ExtractingToken(c)
		err := Validation(tokenString)
		if err != nil {
			c.JSON(http.StatusForbidden, models.ErrorResponse{ErrorDescription: ""})
			c.Abort()
			return
		}
		c.Next()
	}
}
