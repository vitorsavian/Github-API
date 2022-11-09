package health

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	log.Println("Health check called")
	c.Writer.WriteHeader(http.StatusOK)
}
