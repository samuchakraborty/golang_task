package random

import (
	"math/rand"
	"time"
)

func RandomNumber() []int {

	rand.Seed(time.Now().UnixNano())

	array := []int{}
	for i := 0; i < 20; i++ {
		round := rand.Intn(100)
		array = append(array, round)

	}

	return array
}
