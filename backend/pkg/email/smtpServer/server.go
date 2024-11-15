package smtpServer

import (
	"BetterPC_2.0/configs"
	"net/smtp"
)

var SmtpServer *SMTPServer

type SMTPServer struct {
	Sender   string
	Password string
	Host     string
	Port     string
	Auth     smtp.Auth
}

func InitWithConfig(cfg *configs.Config) {
	SmtpServer = &SMTPServer{
		Sender:   cfg.Notifications.Email,
		Password: cfg.Notifications.Password,
		Host:     cfg.Notifications.SmtpHost,
		Port:     cfg.Notifications.SmtpPort,
		Auth: smtp.PlainAuth(
			"",
			cfg.Notifications.Email,
			cfg.Notifications.Password,
			cfg.Notifications.SmtpHost,
		)}
}

func MustGet() *SMTPServer {
	if SmtpServer == nil {
		panic("SMTPServer not initialized")
	}
	return SmtpServer
}
