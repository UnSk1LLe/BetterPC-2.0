package service

import (
	"BetterPC_2.0/pkg/email/smtpServer"
	"BetterPC_2.0/pkg/logging"
)

type NotificationService struct {
	SmtpServer *smtpServer.SMTPServer
	logger     *logging.Logger
}

func NewNotificationService(smtpServer *smtpServer.SMTPServer, logger *logging.Logger) *NotificationService {
	return &NotificationService{
		SmtpServer: smtpServer,
		logger:     logger,
	}
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
