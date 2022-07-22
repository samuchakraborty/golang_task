package main

import (
	"fmt"
	"h/algorithm"
	random "h/random_number"
	"h/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

var start time.Time

func init() {

	start = time.Now()
}

func main() {

	app := fiber.New()

	fmt.Printf("random number %v\n", random.RandomNumber())

	result, executionTime := algorithm.BubbleSort(random.RandomNumber())
	insertionSort, executionTime2 := algorithm.InsertionSort(random.RandomNumber())

	app.Get("", func(c *fiber.Ctx) error {

		insertionSortData := utils.InsertionSort{
			Data:          insertionSort,
			ExecutionTime: executionTime2,
		}
		bubbleSortData := utils.BubbleSort{
			Data:          result,
			ExecutionTime: executionTime,
		}

		response := &utils.Data{
			Message:       "success",
			StatusCode:    200,
			BubbleSort:    bubbleSortData,
			InsertionSort: insertionSortData,
		}
		fmt.Println(response)

		return c.JSON(response)

	})
	app.Get("/index", func(c *fiber.Ctx) error {

		return c.Response().SendFile("./index.html")
	})

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
	fmt.Printf("%v \n", result)
	fmt.Println("main excution end at ", time.Since(start))

}
