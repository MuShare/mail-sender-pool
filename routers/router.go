package routers

import (
	v1 "github.com/MuShare/mail-sender-pool/routers/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/health", v1.HealthCheck)
	}
	return r
}
