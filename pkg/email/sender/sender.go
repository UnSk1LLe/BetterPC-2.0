package sender

import (
	"BetterPC_2.0/pkg/logging"
	"net/smtp"
)

func SendEmail(email string, subject string, body string) error {
	logger := logging.GetLogger()

	to := []string{email}

	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}
	logger.Infof("Message sent to %s", email)
	return nil
}
