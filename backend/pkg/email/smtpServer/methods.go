package smtpServer

import (
	"net/smtp"
)

func (s *SMTPServer) SendEmail(toEmails []string, subject string, body string) error {
	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(s.Host+":"+s.Port, s.Auth, s.Sender, toEmails, message)
	if err != nil {
		return err
	}

	return nil
}
