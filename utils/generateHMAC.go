package utils

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func GenerateHMAC(data string, secret string) string {
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(data))
	result := hex.EncodeToString(h.Sum(nil))
	return result
}
