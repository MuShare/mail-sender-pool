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
	if result := db.Where("today_used_quota = (?)", db.Table("smtp_account").Select("MIN(today_used_quota)")).First(&account); result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

//AddSMTPAccount xxx
func AddSMTPAccount(smtpAccount SMTPAccount) (int, error) {
	result := db.Create(&smtpAccount)
	return smtpAccount.ID, result.Error
}

//GetSMTPAccountByID xx
func GetSMTPAccountByID(id int) (*SMTPAccount, error) {
	var account SMTPAccount
	if result := db.First(&account, id); result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

//IncreaseTodayUsedQouta xxx
func (smtpAccount *SMTPAccount) IncreaseTodayUsedQouta() error {
	smtpAccount.TodayUsedQuota++
	if result := db.Save(smtpAccount); result.Error != nil {
		return result.Error
	}
	return nil
}

//GetAllSMTPAccount xxx
func GetAllSMTPAccount() (*[]SMTPAccount, error) {
	var accounts []SMTPAccount
	result := db.Find(&accounts)
	return &accounts, result.Error
}
