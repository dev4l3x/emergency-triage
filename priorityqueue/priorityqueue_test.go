package priorityqueue

import "testing"


func TestPriorityQueue(t *testing.T) {

	queue := PriorityQueue{}

	queue.Add(5, "John Doe")
	queue.Add(8, "Jane Smith")
	queue.Add(5, "Alice Johnson")

	_, patient := queue.Peek()

	if patient != "Jane Smith" {
		t.Fatalf("Expected %v but got %v", "Jane Smith", patient)
	}

	_, patient2 := queue.PopMax()
	if patient2 != "Jane Smith" {
		t.Fatalf("Expected %v but got %v", "Jane Smith", patient2)
	}
	
	_, patient4 := queue.PopMax()
	if patient4 != "Alice Johnson" {
		t.Fatalf("Expected %v but got %v", "Alice Johnson", patient4)
	}
	
	if queue.IsEmpty() {
		t.Fatal("Expected queue to not to be empty")
	}

	_, patient3 := queue.PopMax()
	if patient3 != "John Doe" {
		t.Fatalf("Expected %v but got %v", "John Doe", patient3)
	}

	if !queue.IsEmpty() {
		t.Fatal("Expected queue to be empty")
	}

}