package main

import (
	"fmt"
	"go-tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// func Walk(t *tree.Tree, ch chan int) {
// 	if t != nil {
// 		Walk(t.Left, ch)
// 		ch <- t.Value
// 		Walk(t.Right, ch)
// 	} else {
// 		close(ch)
// 	}

// }
var levelCheck map[int]int

func traversal(t *tree.Tree) {
	if t == nil {
		return
	}

	fmt.Println(*t)
	traversal(t.Left)
	traversal(t.Right)
}

//Walk  tree
func Walk(t *tree.Tree, ch chan int) {
	sendValue(t, ch)
	close(ch)
}

func sendValue(t *tree.Tree, ch chan int) {
	if t != nil {
		sendValue(t.Left, ch)
		ch <- t.Value
		sendValue(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true

}

func main() {
	var ch = make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("==============================")
	traversal(tree.New(1))
	fmt.Println("==============================")
	traversal(tree.New(2))
	fmt.Println("==============================")
	traversal(tree.New(3))
	fmt.Println("==============================")
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

}
