package utils

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"strings"
)

func Hash(str string) string {
	ByteHash := sha256.Sum256([]byte(str))
	HashedPassword := fmt.Sprintf("%x", ByteHash[:])
	return HashedPassword
}

func IsEmptyString(str any) bool {
	realString, ok := str.(string)
	return !ok || len(strings.TrimSpace(realString)) <= 0
}

func IsValidUrl(URL string) bool {
	_, err := url.ParseRequestURI(URL)
	return err == nil
}
