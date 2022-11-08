package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}
