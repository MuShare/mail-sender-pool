package mail

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"

	"github.com/MuShare/mail-sender-pool/pkg/logging"

	"github.com/MuShare/mail-sender-pool/models"
)

//SendMailWithSMTP send mail
func SendMailWithSMTP(smtpID int, recv string, subj string, contentType string, body string) error {
	var smtpAccount *models.SMTPAccount
	var err error
	smtpAccount, err = models.GetSMTPAccountByID(smtpID)
	if err != nil {
		return err
	}
	return SendMail(smtpAccount, recv, subj, contentType, body)
}

//SendMailWithAutoSelectSMTP xxx
func SendMailWithAutoSelectSMTP(recv string, subj string, contentType string, body string) error {
	var (
		smtpAccount *models.SMTPAccount
		err         error
	)
	smtpAccount, err = models.GetAvailabeSMTPAccount()
	if err != nil {
		return err
	}
	return SendMail(smtpAccount, recv, subj, contentType, body)
}

//SendMail send mail
func SendMail(smtpAccount *models.SMTPAccount, recv string, subj string, contentType string, body string) error {
	from := mail.Address{Name: "", Address: "no_reply@kaboocha.com"}
	to := mail.Address{Name: "", Address: recv}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(subj)) + "?="

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "Content-Type: " + contentType + "; charset=UTF-8\r\n\n" + body

	host, port, err := net.SplitHostPort(smtpAccount.Host)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", smtpAccount.Username, smtpAccount.Password, host)

	if port == "465" {
		tlsconfig := &tls.Config{
			ServerName: host,
		}

		// Here is the key, you need to call tls.Dial instead of smtp.Dial
		// for smtp servers running on 465 that require an ssl connection
		// from the very beginning (no starttls)
		conn, err := tls.Dial("tcp", smtpAccount.Host, tlsconfig)
		if err != nil {
			return err
		}

		c, err := smtp.NewClient(conn, host)
		if err != nil {
			return err
		}

		// Auth
		if err = c.Auth(auth); err != nil {
			return err
		}
		// To && From
		if err = c.Mail(from.Address); err != nil {
			return err
		}

		if err = c.Rcpt(to.Address); err != nil {
			return err
		}

		// Data
		w, err := c.Data()
		if err != nil {
			return err
		}

		_, err = w.Write([]byte(message))
		if err != nil {
			return err
		}

		err = w.Close()
		if err != nil {
			return err
		}

		err = c.Quit()
		if err != nil {
			return err
		}
	} else {
		toAddresses := []string{to.Address}
		err := smtp.SendMail(smtpAccount.Host, auth, from.Address, toAddresses, []byte(message))
		if err != nil {
			return err
		}
	}

	if err := smtpAccount.IncreaseTodayUsedQouta(); err != nil {
		return err
	}
	logging.Info(fmt.Sprintf("succeed to send mail: %s to %s", subj, recv))
	logging.Info(fmt.Sprintf("used host: %s username: %s used quota: %d today's quota: %d", smtpAccount.Host, smtpAccount.Username, smtpAccount.TodayUsedQuota, smtpAccount.QuotaPerDay))
	return nil
}
