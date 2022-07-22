package utils

import "time"

type Data struct {
	Message       string        `json:"msg"`
	StatusCode    int           `json:"statusCode"`
	BubbleSort    BubbleSort    `json:"bubble_sort_data"`
	InsertionSort InsertionSort `json:"insertion_sort_data"`
}

type BubbleSort struct {
	ExecutionTime time.Duration `json:"executionTime"`

	Data []int `json:"result"`
}

type InsertionSort struct {
	ExecutionTime time.Duration `json:"executionTime"`
	Data          []int         `json:"result"`
}

type BodyData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
