package models

//SMTPAccount aoocunt details
type SMTPAccount struct {
	Model

	Host           string `json:"host" gorm:"index:idx_host_username"`
	Username       string `json:"username" gorm:"index:idx_host_username"`
	Password       string `json:"password"`
	TodayUsedQuota int    `json:"today_used_quota"`
	QuotaPerDay    int    `json:"quote_per_day"`
}

//GetAvailabeSMTPAccount get available smtp account which still has quota today
func GetAvailabeSMTPAccount() (SMTPAccount, error) {
	var account SMTPAccount
	return account, nil
}
