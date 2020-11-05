package v1

import (
	"net/http"

	"github.com/MuShare/mail-sender-pool/pkg/logging"

	"github.com/MuShare/mail-sender-pool/service/mail"

	"github.com/MuShare/mail-sender-pool/pkg/e"

	"github.com/MuShare/mail-sender-pool/pkg/app"
	"github.com/gin-gonic/gin"
)

//SMTPAccountRequest xxx
type SMTPAccountRequest struct {
	Host        string `json:"host" binding:"required"`
	Username    string `json:"user_name" binding:"required"`
	Password    string `json:"password" bingding:"required"`
	QuotaPerDay int    `json:"quota_per_day" binding:"required"`
}

//SendMailRequest xxx
type SendMailRequest struct {
	To          string `json:"to" binding:"required" validate:"email"`
	Subject     string `json:"subject" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
	Body        string `json:"body" binding:"required"`
}

//AddSMTPAccount xxx
func AddSMTPAccount(c *gin.Context) {
	appG := app.Gin{C: c}
	var request SMTPAccountRequest
	if err := c.ShouldBind(&request); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	id, err := mail.AddSMTPAccount(request.Host, request.Username, request.Password, request.QuotaPerDay)
	if err != nil {
		logging.Error(err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"inserted_id": id,
	})
}

//SendMail xxx
func SendMail(c *gin.Context) {
	var appG = app.Gin{C: c}
	var request SendMailRequest
	if err := c.ShouldBind(&request); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	err := mail.SendMailWithAutoSelectSMTP(request.To, request.Subject, request.ContentType, request.Body)
	if err != nil {
		logging.Error(err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"result": "ok",
	})
}

//GetAllSMTPAccount xxx
func GetAllSMTPAccount(c *gin.Context) {
	var appG = app.Gin{C: c}
	accounts, err := mail.GetAllSMTPAccount()
	if err != nil {
		logging.Error(err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"accounts": *accounts,
	})
}

//SendMailWithSMTP xxxx
func SendMailWithSMTP() error {
	return nil
}
