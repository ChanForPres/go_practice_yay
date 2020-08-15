// copies elements into a destination slice dst from a source slice src.

// func copy(dst, src []Type) int

package main

import (
	"codingbat/geeks"
	"fmt"
)

func main() {
	// fmt.Println("returned -------->", allStar("abc"))
	// copy(dst []byte, src string) int
	// i := []byte{}
	// x := copy(i, "abc")
	// geeks.f1()
	// fmt.Println(x, i)
	// geeks.F2()
	geeks.F3()
}

func copyFun() {

	// Copy from one slice to another
	var s = make([]int, 3)
	n := copy(s, []int{0, 1, 2, 3}) // n == 3, s == []int{0, 1, 2}
	// Copy from a slice to itself
	fmt.Println(n, s)

}
