// Package main runs the tree Same function on two trees
package main

import (
	"fmt"

	treeFns "github.com/shinroo/gdsc_get_going_with_golang/pkg/tree"
	"golang.org/x/tour/tree"
)

func main() {
	fmt.Println(treeFns.Same(tree.New(1), tree.New(1)))
}
