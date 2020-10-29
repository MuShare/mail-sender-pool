package v1

import (
	"net/http"

	"github.com/MuShare/mail-sender-pool/pkg/app"
	"github.com/MuShare/mail-sender-pool/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Summary return health check result
// @Produce json
func HealthCheck(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, "ok")
}
