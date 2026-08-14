package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// walkRecursive performs an in-order traversal of the tree
func walkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkRecursive(t.Left, ch)
	ch <- t.Value
	walkRecursive(t.Right, ch)
}

// Walk walks the tree t sending all values from the tree to the channel ch, closing it when done
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

// Same determines whether the trees t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// If channel status (open/closed) differs, trees aren't equivalent
		if ok1 != ok2 {
			return false
		}

		// Both channels closed at the same time: trees match
		if !ok1 {
			return true
		}

		// Values don't match
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Print("Walk test: ")
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for val := range ch {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	fmt.Println("Same(tree.New(1), tree.New(1)):", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2)):", Same(tree.New(1), tree.New(2)))
}
