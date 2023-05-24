package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// APIStatus returns an HTTP handler to perform health checks.
func APIStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "everything is ok!")
	}
}
