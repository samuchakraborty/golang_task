package main

import (
	"h/algorithm"
	"h/utils"
	"math/rand"
	"strconv"

	// "strconv"
	"fmt"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/template/html"
)

func main() {
	// engine := html.New("./views", ".html")

	// app := fiber.New(fiber.Config{Views: engine})
	app := fiber.New()

	// Default config

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//arr := []int{15, 5, 24, 8, 1, 3, 6, 90}

	//fmt.Printf("%v, %v", mergeSort, int64((endTime*time.Millisecond)/time.Millisecond))
	app.Static("/dist", "./dist/")

	app.Static("/sevlet", "./h")

	app.Post("/data", func(c *fiber.Ctx) error {

		payload := struct {
			Number string `json:"number"`
		}{}
		fmt.Println(payload)

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		fmt.Println(payload)

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
				round := rand.Intn(1000)
				array = append(array, round)

			}
			fmt.Println(array)
			startTime := time.Now()
			mergeSort := algorithm.MergeSort(array)
			endTime := time.Since(startTime)

			result, bubbleSortExecutionTime := algorithm.BubbleSort(array)
			insertionSort, insertionSortExecutionTime := algorithm.InsertionSort(array)

			insertionSortData := utils.InsertionSort{
				Data:          insertionSort,
				ExecutionTime: insertionSortExecutionTime,
			}
			bubbleSortData := utils.BubbleSort{
				Data:          result,
				ExecutionTime: bubbleSortExecutionTime,
			}

			util := &utils.Data{
				Message:       "success",
				StatusCode:    200,
				BubbleSort:    bubbleSortData,
				InsertionSort: insertionSortData,
				MargeSortTime: endTime,
				MargeSortData: mergeSort,
				UnSortData: array,
			}
			fmt.Println(util.UnSortData)

			return c.JSON(util)

			// return c.Render("index", fiber.Map{
			// 	"msg":                 util.Message,
			// 	"bubble_sort_time":    (util.BubbleSort.ExecutionTime),
			// 	"insertion_sort_time": (util.InsertionSort.ExecutionTime),
			// 	"data":                mergeSort,
			// 	"merge_sort_time":    ( util.MargeSortTime ),
			// })
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

	app.Listen(":8080")
	// fmt.Printf("%v \n", result)
	// fmt.Println("main excution end at ", time.Since(start))

}
