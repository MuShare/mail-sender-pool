package scheduler

import (
	"github.com/MuShare/mail-sender-pool/models"
	"github.com/jasonlvhit/gocron"
)

//Setup xxx
func Setup()  {
	gocron.Every(1).Day().At("00:00").Do(models.RestSMTPQuota)
	gocron.Start()
}
