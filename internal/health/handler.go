package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/health", Check)
}
