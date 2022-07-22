package random

import (
	"math/rand"
	"time"
)

func RandomNumber(number int) []int {

	rand.Seed(time.Now().UnixNano())

	array := []int{}
	for i := 0; i < number; i++ {
		round := rand.Intn(100)
		array = append(array, round)

	}

	return array
}
