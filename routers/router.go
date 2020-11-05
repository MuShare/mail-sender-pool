package routers

import (
	v1 "github.com/MuShare/mail-sender-pool/routers/v1"
	"github.com/gin-gonic/gin"
)

//InitRouter xxx
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/health", v1.HealthCheck)
		apiv1.POST("/add-smtp-account", v1.AddSMTPAccount)
		apiv1.POST("/send-mail", v1.SendMail)
		apiv1.GET("/all-smtp", v1.GetAllSMTPAccount)
	}
	return r
}
