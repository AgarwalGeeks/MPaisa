package util

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomInt generates a random integer between min and max (inclusive).
func RandomInt(min, max int64) (int64, error) {
	if min > max {
		return 0, nil
	}
	delta := max - min + 1
	n, err := rand.Int(rand.Reader, big.NewInt(delta))
	if err != nil {
		return 0, err
	}
	return n.Int64() + min, nil
}

// RandomString generates a random string of the given length.
func RandomString(length int) (string, error) {
	if length <= 0 {
		return "", nil
	}
	var sb strings.Builder
	for i := 0; i < length; i++ {
		n, err := RandomInt(0, int64(len(letterBytes)-1))
		if err != nil {
			return "", err
		}
		sb.WriteByte(letterBytes[n])
	}
	return sb.String(), nil
}

// RandomBytes generates a random byte slice of the given length.
func RandomBytes(length int) ([]byte, error) {
	if length <= 0 {
		return nil, nil
	}
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
