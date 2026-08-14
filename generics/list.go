package main

import "fmt"

// List represents a singly-linked list that holds values of any type
type List[T any] struct {
	next *List[T]
	val  T
}

// Push adds a new node to the front of the list and returns the new head
func (l *List[T]) Push(v T) *List[T] {
	return &List[T]{
		next: l,
		val:  v,
	}
}

// Print displays all values in the list
func (l *List[T]) Print() {
	for curr := l; curr != nil; curr = curr.next {
		fmt.Printf("%v -> ", curr.val)
	}
	fmt.Println("nil")
}

func main() {
	// Start with an empty list
	var numbers *List[int]

	// Add elements to the list
	numbers = numbers.Push(30)
	numbers = numbers.Push(20)
	numbers = numbers.Push(10)

	// Output: 10 -> 20 -> 30 -> nil
	numbers.Print()

	// Works with strings just as easily
	var words *List[string]
	words = words.Push("World")
	words = words.Push("Hello")

	// Output: Hello -> World -> nil
	words.Print()
}
