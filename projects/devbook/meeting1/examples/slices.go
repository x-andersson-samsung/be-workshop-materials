package main

import (
	"fmt"
)

func main() {
	arr := [7]int{0, 1, 2, 3, 4, 5, 6}

	sl := arr[3:6]
	fmt.Println(sl, len(sl), cap(sl)) // What will be printed here?

	sl = append(sl, 7)
	fmt.Println(sl, len(sl), cap(sl)) // What will be printed here?
	fmt.Println(arr)                  // What happened with underlying array?
}
