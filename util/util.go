package util

import "math/rand"

func RandString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, n)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
