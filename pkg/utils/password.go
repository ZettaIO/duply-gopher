package utils

import "math/rand"

const data = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// NewPassword generates a random password of length n
func NewPassword(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = data[rand.Intn(len(data))]
	}
	return string(b)
}
