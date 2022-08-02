package algorithm

import (
	"time"
)

func InsertionSort(arr []int) ([]int, time.Duration) {
	startTime := time.Now()
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			j--
		}

	}
	endTime := time.Since(startTime)
	
	return arr, endTime
}
