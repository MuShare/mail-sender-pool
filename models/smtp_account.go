package models

//SMTPAccount aoocunt details
type SMTPAccount struct {
	Model

	Host           string `json:"host" gorm:"uniqueIndex:idx_host_username;size:100;not null"`
	Username       string `json:"username" gorm:"uniqueIndex:idx_host_username;size:100;not null"`
	Password       string `json:"password" gorm:"size:100;not null"`
	TodayUsedQuota int    `json:"today_used_quota"`
	QuotaPerDay    int    `json:"quote_per_day"`
}

//GetAvailabeSMTPAccount get available smtp account which still has quota today
func GetAvailabeSMTPAccount() (*SMTPAccount, error) {
	var account SMTPAccount
	if result := db.First(&account); result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

//AddSMTPAccount xxx
func AddSMTPAccount(smtpAccount SMTPAccount) error {
	result := db.Create(&smtpAccount)
	return result.Error
}
