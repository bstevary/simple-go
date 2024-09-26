package test

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVW"

// RandomInt generates a cryptographically secure random integer between min and max
func RandomInt(min, max int64) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		panic(err)
	}
	return min + n.Int64()
}

// RandomString generates a cryptographically secure random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := big.NewInt(int64(len(alphabet)))

	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, k)
		if err != nil {
			panic(err)
		}
		sb.WriteByte(alphabet[num.Int64()])
	}

	return sb.String()
}

// RandomName generates a cryptographically secure random name
func RandomName() string {
	return RandomString(6)
}

// RandomEmail generates a cryptographically secure random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
