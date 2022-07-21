package main

import (
	"fmt"

	"time"

	"github.com/gofiber/fiber/v2"
)

var start time.Time

func init2() {

	start = time.Now()
}

type Data struct {
	Message       string `json:"msg"`
	StatusCode    int    `json:"statusCode"`
	ExecutionTime string `json:"executionTime"`
	Data          []int  `json:"data"`
}

func main() {

	app := fiber.New()

	var array = []int{28, 12, 23, 7, 9, 60, 20}

	fmt.Println("main excution start at ", time.Since(start))
	
	result := bubbleSort(array)

	app.Get("/", func(c *fiber.Ctx) error {

		// enc := json.NewEncoder(os.Stdout)
		// if err := enc.Encode(user); err != nil {
		// 	fmt.Printf("error encoding struct into JSON: %v\n", err)
		// 	return c.Status(200).JSON((err))
		// } else {
		// 	return c.Status(200).JSON((err))

		// }
		// init2()

		response := &Data{
			Message:       "success",
			StatusCode:    200,
			ExecutionTime: time.Since(start).String(),
			Data:          result,
		}
		fmt.Println(response)
		// time.NewTimer()(1 )
		//const timeout = 1 * time.Second
		
		return c.JSON(response)

	})
	app.Listen(":3000")
	fmt.Printf("%v \n", result)
	fmt.Println("main excution end at ", time.Since(start))

}

func bubbleSort(arr []int) []int {
	array := arr
	fmt.Printf("%v \n", array)

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
	time.Sleep(20 * time.Millisecond)
	return array
}

// func insertionSort(arr []int) []int {

// 	for i := 0; i < len(arr)-1; i++ {
// 		if arr[i] > arr[i+1] {

// 			arr[i] = arr[i+1]

// 		}
// 	}
// return arr
// }
