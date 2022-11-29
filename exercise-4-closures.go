package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	
	prevPrev, prev := 0,1
	
	return func() int {
		num:= prevPrev + prev
		
		if prev + prevPrev == 1 {
			prev = 1
		}
		
		prevPrev = prev
		prev = num
		
		return num
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
