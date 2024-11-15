package service

import (
	"BetterPC_2.0/pkg/email/smtpServer"
)

type NotificationService struct {
	SmtpServer *smtpServer.SMTPServer
}

func NewNotificationService(smtpServer *smtpServer.SMTPServer) *NotificationService {
	return &NotificationService{SmtpServer: smtpServer}
}

func (n *NotificationService) SendEmailToUser(userEmail, subject, body string) error {
	err := n.SmtpServer.SendEmail([]string{userEmail}, subject, body)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotificationService) SendEmailToGroup(userEmailList []string, subject, body string) error {
	err := n.SmtpServer.SendEmail(userEmailList, subject, body)
	if err != nil {
		return err
	}
	return nil
}
