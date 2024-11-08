package sender

import (
	"BetterPC_2.0/configs"
	"net/smtp"
)

var (
	from,
	password,
	smtpHost,
	smtpPort string
	auth smtp.Auth
)

func Init(cfg *configs.Config) {
	from = cfg.Notifications.Email
	password = cfg.Notifications.Password
	smtpHost = cfg.Notifications.SmtpHost
	smtpPort = cfg.Notifications.SmtpPort
	auth = smtp.PlainAuth("", from, password, smtpHost)
}
