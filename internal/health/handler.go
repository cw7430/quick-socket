package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// RegisterRoutes health check godoc
// @Summary Health check
// @Description 서버 상태를 확인합니다.
// @Tags Health
// @Success 204
// @Router /health [get]
func RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/health", Check)
}
