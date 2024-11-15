package passwordHasher

import "golang.org/x/crypto/bcrypt"

func GeneratePasswordHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), err
}

func CompareHashAndPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
