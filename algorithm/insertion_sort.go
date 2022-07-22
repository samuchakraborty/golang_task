package algorithm

import (
	"fmt"
	"time"
)
var start time.Time

func InsertionSort(arr []int) ([]int, time.Duration) {
	fmt.Printf(" arr ::: %v \n", arr)
	startTime := time.Since(start)
	for i := 1; i < len(arr); i++ {
		//fmt.Printf(" ::: %v :::: i ?? %v \n", arr[i], i)
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				//fmt.Printf("%v %v :: \n\t", arr[j], arr[j-1])
				arr[j], arr[j-1] = arr[j-1], arr[j]
				//fmt.Printf("%v =>>>> %v :: \n\t\t", arr[j], arr[j-1])
			}
			j--
		}

	}
	endTime := time.Since(start)
	fmt.Printf("start time %v , end Time %v\n", startTime, endTime)
	executionTime := endTime - startTime
	return arr, executionTime
}
