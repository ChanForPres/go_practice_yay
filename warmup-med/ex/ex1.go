package main

import "fmt"

// func changeArr() {
// 	// Such an example is looking up a value from a
// 	// map where the key type is string,
// 	// and you index the map with a []byte,
// 	// converted to string of course (source):
// 	s1 := []byte("one")
// 	m := map[string]int{
// 		"one":   1,
// 		"two":   2,
// 		"three": 3,
// 	}

// 	i, j := m[string(s1)]
// 	// i = one, j = 1
// 	fmt.Println("i: %#v\t j:%#v", i, j)
// 	// 2nd one is
// 	s := "abc"

// 	// for i1, j1 := range []byte(s) {

// 	// }

// 	for k, v := range m {
// 		fmt.Println(k, v)
// 		// one,1 ..

// 	}
// }

func arr1(arr *[]int) *[]int {
	//edit array
	lenArr := len(*arr)
	var addrA *[]int
	println(addrA)
	if lenArr <= 0 {

		fmt.Println("i: %#v\t j:", arr)
		return arr
	} else {
		a := *arr
		addrA = &a
		fmt.Println(len(a), cap(a))
		a[0] = 1
		arr = &a

		*arr = append(*arr, 3)
		fmt.Println(len(*arr), cap(*arr))
	}

	fmt.Println(len(*arr), cap(*arr))
	addrA = arr
	fmt.Printf("Addr before append : %p\n", &arr)
	*arr = append(*arr, 3)

	//	fmt.Printf("Addr after append : %p\n", &arr)
	//	fmt.Printf("Addr of after affend arr[0]: %p\n", &arr[0])
	fmt.Printf("after append before return i: %#v\t j:\n", arr)
	return arr
}
func arr2(arr []int) []int {
	//edit array
	lenArr := len(arr)
	if lenArr <= 0 {

		fmt.Println("i: %#v\t j:", arr)
		return arr
	} else {
		arr[0] = 1
	}

	fmt.Println(len(arr), cap(arr))

	fmt.Printf("Addr before append : %p\n", &arr)
	arr = append(arr, 3)

	fmt.Printf("Addr after append : %p\n", &arr)
	fmt.Printf("Addr of after affend arr[0]: %p\n", &arr[0])
	fmt.Printf("after append before return i: %#v\t j:\n", arr)
	return arr
}

func addrOf() {
	intarr := [5]int{12, 34, 55, 66, 43}
	slice := intarr[:]
	fmt.Printf("the len is %d and cap is %d \n", len(slice), cap(slice))
	fmt.Printf("address of slice 0x%x add of Arr 0x%x \n", &slice, &intarr)
	fmt.Printf("address of slice %p add of Arr %p \n", &slice, &intarr)

	s := make([]int, 10)
	fmt.Printf("Addr of first element: %p\n", &s[0])
	fmt.Printf("Addr of slice itself:  %p\n", &s)
}
func main() {
	var arr *[]int
	arr = &[]int{1, 2, 3}

	fmt.Printf("address of slice %p add of Arr  \n", &arr)

	arr2 := arr1(arr)

	fmt.Printf("\n\n\nafter return i: %#v\t j:\n", arr, arr2)

	fmt.Printf("address of slice %p add of Arr  \n", &arr)
	//looking up a value from a map where
	// the key type is string, and you index the map with a []byte

	// key := []byte("some key")

	// //	var m map[string]
	// //	// ...
	// //	v, ok := m[string(key)] // Copying key here is optimized away
	// //
	// // range over string to byte slice
	// var s1 string
	// // []byte(s1) CREATES A COPY

	// for i, j := range []byte(s1) {
	// 	// i index
	// 	// j val
	// 	// []byte array doesnt create copy of string

	// }

	// change array content

}
