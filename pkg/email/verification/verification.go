package verification

import (
	"errors"
	"net"
	"net/smtp"
	"regexp"
	"strings"
)

func isValidEmailFormat(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func hasValidMX(domain string) bool {
	mxRecords, err := net.LookupMX(domain)
	return err == nil && len(mxRecords) > 0
}

func checkSMTP(email string) (bool, error) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false, errors.New("invalid emailVerification format")
	}
	domain := parts[1]

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return false, err
	}

	client, err := smtp.Dial(mxRecords[0].Host + ":25")
	if err != nil {
		return false, err
	}
	defer func(client *smtp.Client) {
		_ = client.Close()
	}(client)

	err = client.Hello("example.com")
	err = client.Mail("you@example.com")
	if err := client.Rcpt(email); err != nil {
		return false, err
	}

	return true, nil
}

func IsVerifiedEmail(email string) (bool, error) {
	if !isValidEmailFormat(email) {
		return false, errors.New("invalid emailVerification format")
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false, errors.New("invalid emailVerification format")
	}
	domain := parts[1]

	if !hasValidMX(domain) {
		return false, errors.New("invalid emailVerification domain")
	}

	smtpCheck, err := checkSMTP(email)
	if err != nil {
		return false, err
	}
	return smtpCheck, nil
}
