package tool

import "math/rand"

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func RandStr(n int) string {
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(bytes)
}
