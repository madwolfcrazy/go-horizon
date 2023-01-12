package helper

import (
	"math/rand"
	"time"
)

func RandInt(maxVal int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxVal)
}
