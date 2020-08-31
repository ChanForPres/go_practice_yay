package main

import (
	"fmt"
	"unsafe"
)


func main() {

	s1 := make([]int, 3, 5) // len cap

	copy(s1, []int{1, 2, 3}) //

	fmt.Println(len(s1), cap(s1), &s1[0])

	s1 = append(s1, 4) // len 4

	fmt.Println(len(s1), cap(s1), &s1[0])

	s2 := s1[1:] // len 3  points to backing array

	fmt.Println(len(s2), cap(s2), &s2[0])
	fmt.Printf("\n%d %d\n", &s1[0], &s2[0])
	var  x int
	fmt.Println(unsafe.Sizeof(x))

}
