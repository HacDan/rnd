package generator

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

func GenerateBase32(length int) (string, error) {
	letters := "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes), nil
}

func GenerateRandomHex(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateRandomStringBase62(length int, prefix string, lower bool) (string, error) {
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

	if lower {
		return strings.ToLower(prefix + string(bytes)), nil
	}

	return prefix + string(bytes), nil
}

func GeneratePassword(length int, numbers, symbols bool) (string, error) {
	letters := `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`
	if numbers {
		letters += "0123456789"
	}
	if symbols {
		letters += `~!@#$%^&*()_-+={[}]|\:;"'<,>.?/`
	}

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
