package id

import (
	"crypto/rand"
	"encoding/hex"
)

func Gen() string {
	bytes := make([]byte, 16) // 16 bytes = 128 bits
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err) // Handle error as needed
	}
	return hex.EncodeToString(bytes)
}
