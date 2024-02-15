# Emergency Room Triage
A pet project to practice priority queues

### Background

An emergency room (ER) at a hospital categorizes patients based on the urgency of their conditions. Urgent conditions require immediate attention, while less critical conditions can wait a bit longer. This system is a perfect application for a priority queue, where the priority is determined by the urgency of the patient's condition.

### Task

Implement a priority queue to manage ER patients. Your priority queue should support the following operations:

1. **Insert/Add patient**: Add a new patient to the queue with an urgency level. The urgency level is an integer where a higher number indicates a more urgent need for care.
2. **Remove patient**: Remove and return the patient with the highest urgency level from the queue. If there are multiple patients with the same urgency level, remove the one who arrived first.
3. **Peek**: Return the patient with the highest urgency level without removing them from the queue.
4. **IsEmpty**: Check if the queue is empty.

### Implementation Details

- You can choose to implement this priority queue using an array, a linked list, or a binary heap. For simplicity and efficiency, a binary heap is recommended.
- Ensure your implementation maintains the heap property after each operation.
- If implementing with an array or a linked list, you might need to sort the elements or find the maximum urgency level each time you remove a patient, which could be less efficient.

### Example

```
Add patient "John Doe" with urgency 5
Add patient "Jane Smith" with urgency 8
Add patient "Alice Johnson" with urgency 5
Peek -> Should return "Jane Smith" (highest urgency)
Remove patient -> Should remove and return "Jane Smith"
Peek -> Should now return "John Doe" or "Alice Johnson" (whichever is first in your implementation with urgency 5)
IsEmpty -> Should return false
Remove patient -> Should remove and return "John Doe" or "Alice Johnson"
Remove patient -> Should remove and return the remaining patient with urgency 5
IsEmpty -> Should return true after all patients are removed
```

### Advanced Extension (Optional)

- Implement dynamic resizing for your priority queue if using an array-based implementation.
- Add functionality to change the urgency level of a patient already in the queue, adjusting the queue accordingly.

This exercise will test your ability to implement and manipulate complex data structures, manage priorities, and ensure the integrity of the structure after each operation.
