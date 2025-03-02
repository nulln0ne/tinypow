package core

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256 returns the SHA-256 hash of the input data
func Sha256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// Sha256Hash returns the SHA-256 hash of the input data as a hex string
func Sha256Hash(data string) string {
	return hex.EncodeToString(Sha256([]byte(data)))
}

// Sha256HashBytes returns the SHA-256 hash of the input data as a hex string
func Sha256HashBytes(data []byte) string {
	return hex.EncodeToString(Sha256(data))
}
