package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if t == nil {
	   return
	}	

	Walk(t.Left, ch)
	ch <- t.Value	
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	chan1 := make(chan int)
	chan2 := make(chan int)
	
	arr1 := make([]int, 10)
	arr2 := make([]int, 10)

	go func () {
        defer close(chan1)
        Walk(t1, chan1)
    }()
	
	go func () {
        defer close(chan2)
        Walk(t2, chan2)
    }()
	
	for i:= 0; i < 10; i++ {
		arr1 = append(arr1, <- chan1)
		arr2 = append(arr2, <- chan2)
	}
	
	
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	
	return true
}

func main() {

	ch := make(chan int)
	go func () {
        defer close(ch)
        Walk(tree.New(1), ch)
    }()
    for i := range ch {
        fmt.Println(i)
    }

	result := Same(tree.New(1), tree.New(2))
	fmt.Println(result)
}
