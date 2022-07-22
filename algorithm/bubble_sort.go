package algorithm

import (
	"fmt"
	"time"
)

func BubbleSort(arr []int) ([]int, time.Duration) {
	array := arr
	fmt.Printf("%v \n", array)
	startTime := time.Since(start)
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				num := array[j]
				array[j] = array[j+1]
				array[j+1] = num

				//fmt.Printf("%v %v \n", array[j], array[j+1])

			}

		}
	}

	endTime := time.Since(start)
	fmt.Printf("start time %v , end Time %v\n", startTime, endTime)
	executionTime := endTime - startTime
	fmt.Print(executionTime)

	return array, executionTime
}

