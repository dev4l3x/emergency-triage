package main

import (
	"emergency-triage/priorityqueue"
	"fmt"
)

func main() {
	queue := priorityqueue.PriorityQueue{}

	queue.Add(1, "Test")
	queue.Add(2, "Test")
	queue.Add(3, "Test")
	queue.Add(4, "Test")
	priorityqueue.PrintQueue(queue)
	
	maxPriority, maxValue := queue.PopMax()
	fmt.Println("Max priority:", maxPriority, "Max value:", maxValue)
	priorityqueue.PrintQueue(queue)

	queue.PopMax()
	priorityqueue.PrintQueue(queue)
}