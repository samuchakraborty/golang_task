package main

import (
	"h/algorithm"
	"h/utils"
	"math/rand"
	"strconv"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	//arr := []int{15, 5, 24, 8, 1, 3, 6, 90}

	//fmt.Printf("%v, %v", mergeSort, int64((endTime*time.Millisecond)/time.Millisecond))
	app.Static("/dist", "./dist/")
	
	app.Post("", func(c *fiber.Ctx) error {

		payload := struct {
			Number string `json:"number"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		// fmt.Println(payload)

		// input := (c.Query("number"))
		// n := payload.Number(int)
		if payload.Number != "" {
			number, err := strconv.Atoi(payload.Number)
			if err != nil {
				panic(err)
			}
			rand.Seed(time.Now().UnixNano())

			array := []int{}
			for i := 0; i < number; i++ {
				round := rand.Intn(100)
				array = append(array, round)

			}
			startTime := time.Now()

			mergeSort := algorithm.MergeSort(array)
			endTime := time.Since(startTime)

			result, executionTime := algorithm.BubbleSort(array)
			insertionSort, executionTime2 := algorithm.InsertionSort(array)

			insertionSortData := utils.InsertionSort{
				Data:          insertionSort,
				ExecutionTime: executionTime2 * time.Millisecond,
			}
			bubbleSortData := utils.BubbleSort{
				Data:          result,
				ExecutionTime: executionTime,
			}

			util := &utils.Data{
				Message:       "success",
				StatusCode:    200,
				BubbleSort:    bubbleSortData,
				InsertionSort: insertionSortData,
				MargeSortTime:  endTime * time.Millisecond,
			}

			// return c.JSON(response)

			return c.Render("index", fiber.Map{
				"msg":                 util.Message,
				"bubble_sort_time": 
					util.BubbleSort.ExecutionTime ,
				"insertion_sort_time": int64(util.InsertionSort.ExecutionTime / time.Millisecond),
				"data":                mergeSort,
				"merge_sort_time":     int64(util.MargeSortTime / time.Millisecond),
			})
		} else {
			return c.Render("index", fiber.Map{})

		}

	})

	app.Get("", func(c *fiber.Ctx) error {

		return c.Render("index", fiber.Map{})
	})

	// app.Post("", func(c *fiber.Ctx) error {

	// 	payLoad := utils.BodyData{}
	// 	if err := c.BodyParser(&payLoad); err != nil {
	// 		return err
	// 	}
	// 	fmt.Println(payLoad)

	// 	return c.SendString("post work")
	// })
	// app.Post("", func(c *fiber.Ctx) error {
	// 	payload := struct {
	// 		Number string `json:"number"`
	// 	}{}

	// 	if err := c.BodyParser(&payload); err != nil {
	// 		return err
	// 	}
	// 	fmt.Print(payload)
	// 	// return c.JSON(payload)
	// 	return c.Redirect("index")
	// })

	app.Listen(":3000")
	// fmt.Printf("%v \n", result)
	// fmt.Println("main excution end at ", time.Since(start))

}
