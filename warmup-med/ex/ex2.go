package main

import (
	"fmt"
	"strings"
)

func f1() {
	var b strings.Builder
	// prealloc memory
	b.Grow(32) // 32 byte str
	fmt.Println("cap: len: ", b.Cap(), b.Len())
	byArr := [32]byte{}
	b.Write(byArr[:])

	fmt.Printf("resulting str: %#v\n", b.String())
	b.Reset()

	fmt.Printf("resulting str: %#v\n", b.String())
}

func f2(nums ...int) *[]int {
	//a3 := []int{nums}
	nums[0] = 0
	a := []int{1, 2}
	b := []int{11, 22}
	a = append(nums, b...) // a == [1 2 11 22]
	return &a
}

type Str struct {
	I int64
	F float64
	// j uintptr
}

func Generate(count int) (value []byte) {
	value = make([]byte, count)
	for i := 0; i < count; count++ {
		value[i] = byte(i) // just for an example
	}
	return
}

func ptrRet(ptr *int, v int) (*int, int) {
	// copy of ptr param is made

	fmt.Println(&ptr)
	v += 1
	return ptr, v
}

func mapEdit() {
	
}
// refr types : map,slice,ptr,chan
// Go types (e.g. pointers, slices, channels, maps) are reference type
//func getter( arr const []int ) {
//	for i,j := range
//}

func ptr() (*Str, *Str) {
	s1 := new(Str)
	var s2 *Str
	s2 = s1
	return s1, s2
}

func main() {
	i := 5
	ij := &i
	a, b := ptrRet(ij, *ij)

	// a, b := ptr()
	fmt.Println(&a, b)
	
	

	//f1()
	// arr := []int{1, 2, 3}
	// arrAddr := &arr

	// fmt.Println(arrAddr)
	// arr2 := f2(arr...)
	// fmt.Println(arr2)
	// var s1 []int
	// f2(s1...)
	// fmt.Println(unsafe.Sizeof(s1))
	// // ptr = 8, int = 8
	// fmt.Println(len(s1), cap(s1))
	// var s2 []int
	// copy(s2, s1)

	// var ascStr []byte
	// copy(ascStr, "string")

	// // bytes to string
	// var bArr string
	// copy(bArr, string([]byte{1, 2, 3}))
	// conversion makes new string with same bytes as byte slice

}
