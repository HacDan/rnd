package generator

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"strings"
)

// GenerateRandomASCIIString returns a securely generated random ASCII string.
// It reads random numbers from crypto/rand and searches for printable characters.
// It will return an error if the system's secure random number generator fails to
// function correctly, in which case the caller must not continue.

func GenerateRandomStringBase62(length int, prefix string) (string, error) {
	letters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if len(prefix) > 0 {
		length = length - len(prefix)
	}
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return prefix + string(bytes), nil
}

func GenerateRandomStringBase62Lower(length int, prefix string) (string, error) {
	s, err := GenerateRandomStringBase62(length, prefix)
	if err != nil {
		return "", err
	}
	return strings.ToLower(s), nil
}

func GenerateRandomNum(length int) (int, error) {
	numbers := "0123456789"

	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return -1, err
	}
	for i, b := range bytes {
		bytes[i] = numbers[b%byte(len(numbers))]
	}
	i, err := strconv.Atoi(string(bytes))
	if err != nil {
		return -1, nil
	}
	return i, nil
}

func GenerateRandomHex(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
