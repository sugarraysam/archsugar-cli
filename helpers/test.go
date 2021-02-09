package helpers

import "math/rand"

func GetRandomDigit() int {
	min, max := 10000, 20000
	return rand.Intn(min) + max - min
}
