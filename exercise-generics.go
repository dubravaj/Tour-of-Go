package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (list List[T]) printList() {
	
	fmt.Println(list.val)
	
	next := list.next
	
	if next == nil {
		return 
	}
	
	next.printList()
}

func main() {
	
	list:= List[int]{&List[int]{&List[int]{nil, 45}, 22}, 15}
	list.printList()

}
