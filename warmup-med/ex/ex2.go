package main

import (
	"fmt"
	
	"unsafe"
)

func main() {
	var s1 []int
	fmt.Println(unsafe.Sizeof(s1))
	// ptr = 8, int = 8
	fmt.Println(len(s1), cap(s1))
	
}
