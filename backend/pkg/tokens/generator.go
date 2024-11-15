package tokens

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
)

func GenerateNewUUID(email string) (string, error) {
	u := uuid.New()

	combined := append(u[:], []byte(email)...)

	hash := sha256.Sum256(combined)

	return hex.EncodeToString(hash[:]), nil
}
