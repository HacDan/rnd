package generator

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
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
	letters := "0123456789abcdefghijklmnopqrstuvwxyz"
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

func GeneratePassword(length int) (string, error) {
	letters := `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz~!@#$%^&*()_-+={[}]|\:;"'<,>.?/`

	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes), nil
}

func GenerateRandomNum(length int) (int, error) {
	fmt.Println("Length: ", length)
	b := make([]int, length)

	for i := 0; i < length; i++ {
		x, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			return -1, err
		}
		b = append(b, int(x.Int64()))
	}

	i := sliceToInt(b)
	fmt.Println(b)
	return i, nil
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func GenerateRandomHex(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
