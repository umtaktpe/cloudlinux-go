package utils

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func GenerateToken(username, secretKey string) string {
	timestamp := time.Now().Unix()

	// Generate hash
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s%d", secretKey, timestamp)))
	bs := h.Sum(nil)
	hash := fmt.Sprintf("%x", bs)

	return fmt.Sprintf("%s|%d|%s", username, timestamp, hash)
}
