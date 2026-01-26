package main

import (
	"fmt"
)

func main() {
	// define a slice of size 3 and cap 5
	slice := make([]int, 3, 5)
	fmt.Printf("slice: %p\n", slice)

	// Let's add 2 values
	slice = append(slice, 1)
	slice = append(slice, 2) // []int{0,0,0,1,2}
	fmt.Printf("slice: %p\n", slice)

	// This append will go beyond capacity
	// and will cause reallocation and copy
	slice = append(slice, 3)
	fmt.Printf("slice: %p\n", slice)

	// []int{0,0,0,1,2,3}

}
