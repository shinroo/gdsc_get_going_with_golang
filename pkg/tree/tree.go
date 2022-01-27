package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Determines the number of nodes in the tree t
// By walking and incrementing a counter
func Length(t *tree.Tree) int {
	ch := make(chan int)
	go walkProducer(t, ch)
	idx := 0
	for range ch {
		idx++
	}
	return idx
}

// walkProducer walks the tree t sending all values
// from the tree to the channel ch.
// Closes channel when done
func walkProducer(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	testTree := tree.New(5)

	ch := make(chan int)
	go Walk(testTree, ch)

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
