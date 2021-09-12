package rnd

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func StringWithLength(resultLength int) string {
	r := make([]rune, resultLength)

	for i := range r {
		r[i] = letters[rand.Intn(len(letters))]
	}

	return string(r)
}
