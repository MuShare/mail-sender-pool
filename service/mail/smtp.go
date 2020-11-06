package mail

import (
	"github.com/MuShare/mail-sender-pool/models"
)

//AddSMTPAccount xxx
func AddSMTPAccount(host string, username string, password string, quotaPerDay int) (int, error) {
	return models.AddSMTPAccount(models.SMTPAccount{
		Host:        host,
		Username:    username,
		Password:    password,
		QuotaPerDay: quotaPerDay,
	})
}

//GetAvailableSMTPAccount xxx
func GetAvailableSMTPAccount() (*models.SMTPAccount, error) {
	var result *models.SMTPAccount
	var err error
	if result, err = models.GetAvailabeSMTPAccount(); err != nil {
		return nil, err
	}
	return result, nil
}

//GetAllSMTPAccount xxx
func GetAllSMTPAccount() (*[]models.SMTPAccount, error) {
	return models.GetAllSMTPAccount()
}
