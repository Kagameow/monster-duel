package utils

import "math/rand"

func RandomRangeInt(min, max int) int {
	return rand.Intn(max+1-min) + min
}
