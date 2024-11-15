package helpers

import "strings"

func HideEmail(email string) string {
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		return email
	}

	localPart := email[:atIndex]
	domainPart := email[atIndex:]

	if len(localPart) <= 2 {

		return localPart + domainPart
	}

	hiddenEmail := localPart[:2] + "***" + domainPart

	return hiddenEmail
}
