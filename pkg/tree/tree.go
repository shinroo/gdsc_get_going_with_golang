// Utility functions for the Tree data structure
// provided by the golang tour
package tree

import (
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
	l1 := Length(t1)
	l2 := Length(t2)

	if l1 != l2 {
		return false
	}

	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < l1; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}
