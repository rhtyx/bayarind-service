package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"

	"github.com/labstack/gommon/log"
)

type HMAC struct {
	SecretKey []byte
}

var Hmac *HMAC

func InitHMAC() {
	secretKey, err := os.ReadFile("./cert/id_rsa.pub")
	if err != nil {
		log.Fatal(err)
	}

	Hmac = &HMAC{SecretKey: secretKey}
}

func (h HMAC) ValidMAC(message, messageMAC []byte) bool {
	mac := hmac.New(sha256.New, h.SecretKey)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)

	return hmac.Equal(messageMAC, expectedMAC)
}

func (h HMAC) GenerateHMAC(message []byte) string {
	mac := hmac.New(sha256.New, h.SecretKey)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)

	return hex.EncodeToString(expectedMAC)
}
