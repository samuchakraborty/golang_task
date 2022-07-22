package main

import (
	"fmt"
	"h/algorithm"
	"h/utils"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	//app := fiber.New()
	//app.Static("", "./index.html")
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Static("./style", "./styles/style")
	app.Get("/index", func(c *fiber.Ctx) error {

		input := (c.Query("number"))
		number, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		rand.Seed(time.Now().UnixNano())

		array := []int{}
		for i := 0; i < number; i++ {
			round := rand.Intn(100)
			array = append(array, round)

		}

		result, executionTime := algorithm.BubbleSort(array)
		insertionSort, executionTime2 := algorithm.InsertionSort(array)

		insertionSortData := utils.InsertionSort{
			Data:          insertionSort,
			ExecutionTime: executionTime2,
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
		}

		// return c.JSON(response)

		return c.Render("index", fiber.Map{

			"msg":                 util.Message,
			"bubble_sort_time":    util.BubbleSort.ExecutionTime,
			"insertion_sort_time": util.InsertionSort.ExecutionTime,
			"data":                util.BubbleSort.Data,
		})

	})

	// app.Get("/", func(c *fiber.Ctx) error {

	// 	return c.Render("index", fiber.Map{

	// 		"msg": "samu",
	// 	})
	// })

	app.Post("", func(c *fiber.Ctx) error {

		payLoad := utils.BodyData{}
		if err := c.BodyParser(&payLoad); err != nil {
			return err
		}
		fmt.Println(payLoad)

		return c.SendString("post work")
	})
	app.Post("/post", func(c *fiber.Ctx) error {
		payload := struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}
		fmt.Print(payload)
		return c.JSON(payload)
	})

	app.Listen(":3000")
	// fmt.Printf("%v \n", result)
	// fmt.Println("main excution end at ", time.Since(start))

}
