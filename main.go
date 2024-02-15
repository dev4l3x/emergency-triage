package main

import (
	"emergency-triage/priorityqueue"
	"fmt"
)

func main() {
	queue := priorityqueue.PriorityQueue{}

	queue.Add(5, "John Doe")
	queue.Add(8, "Jane Smith")
	queue.Add(5, "Alice Johnson")

	_, patient := queue.PopMax()
	queue.PopMax()
	queue.PopMax()
	fmt.Println(patient)
	fmt.Println(queue.IsEmpty())
}