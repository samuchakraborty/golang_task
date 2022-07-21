package main

import (
	"fmt"

	"time"

	"github.com/gofiber/fiber/v2"
)

var start time.Time

func init() {

	start = time.Now()
}

type Data struct {
	Message       string        `json:"msg"`
	StatusCode    int           `json:"statusCode"`
	ExecutionTime time.Duration `json:"executionTime"`
	BubbleSort    BubbleSort    `json:"bubble_sort_data"`
	InsertionSort InsertionSort `json:"insertion_sort_data"`
}

type BubbleSort struct {
	ExecutionTime time.Duration `json:"executionTime"`

	Data          []int         `json:"result"`
}

type InsertionSort struct {
	ExecutionTime time.Duration `json:"executionTime"`
	Data          []int         `json:"result"`
}

func main() {

	app := fiber.New()

	var array = []int{28, 12, 23, 7, 9, 60, 20}


	result, executionTime := bubbleSort(array)
	insertionSort, executionTime2 := insertionSort(array)

	app.Get("", func(c *fiber.Ctx) error {

		insertionSortData := InsertionSort{
			Data:          insertionSort,
			ExecutionTime: executionTime2,
		}
		bubbleSortData := BubbleSort{
			Data:          result,
			ExecutionTime: executionTime,
		}

		response := &Data{
			Message:       "success",
			StatusCode:    200,
			ExecutionTime: executionTime,
			BubbleSort:    bubbleSortData,
			InsertionSort: insertionSortData,
		}
		fmt.Println(response)
	
		return c.JSON(response)

	})
	app.Get("/index", func(c *fiber.Ctx) error {

		return c.Response().SendFile("./index.html")
	})
	app.Listen(":3000")
	fmt.Printf("%v \n", result)
	fmt.Println("main excution end at ", time.Since(start))

}

func bubbleSort(arr []int) ([]int, time.Duration) {
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

func insertionSort(arr []int) ([]int, time.Duration) {
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
			j = j - 1
		}

	}
	endTime := time.Since(start)
	fmt.Printf("start time %v , end Time %v\n", startTime, endTime)
	executionTime := endTime - startTime
	return arr, executionTime
}
