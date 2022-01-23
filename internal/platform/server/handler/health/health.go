package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckHandler returns an HTTP handler to perform the health checks
func CheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "everything is 0k")
	}
}
