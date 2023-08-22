package utils

import "math/rand"

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomUrl(size int) string {
	str := make([]rune, size)
	rlen := len(runes)

	for i := range str {
		str[i] = runes[rand.Intn(rlen)]
	}

	return string(str)
}
