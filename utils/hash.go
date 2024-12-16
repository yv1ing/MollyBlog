package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Hash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
