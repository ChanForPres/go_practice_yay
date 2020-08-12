package main

import "fmt"

func main() {

	const n = 9876543210 * 9876543210
	fmt.Println(float64(n))

	fmt.Println(complex128(n))
	//fmt.Println(uint64(n))

	var x uint64 = 18446744073709551615
	var y int64 = int64(x)

	fmt.Printf("y val -> %#v\n", y)

	// repr in decimal = refpr in hex,
	fmt.Printf("uint64: %v = %#[1]x, int64: %v = %#x\n", x, y, uint64(y))

	// fmt.Printf(" \n", x, y, uint64(y))

	// string is array of bytes ie 32 bit ascii repr per char
	str := "\x41\x42\x43\x44\x45"

	fmt.Printf("y val -> %#v\n", str)

	// strings are immutable in go
	// must convert to a byte string first

	// s1 := []byte(str)
	s1 := []rune(str)
	// iterating substring

	for i, s := range s1 {
		// range over string as slice not array

		fmt.Printf("y val -> %#v\n", str)

		fmt.Printf("y val -> %#v\n", s1)
		s1[i] = ('\x46') //

		fmt.Printf("%#v\n", s1[i])

		fmt.Printf("%#v\n", s)
	}

	// unicode is superset of ascii table

	// 0xFF == 15*16+ 15*1 = 255
	// 16*16 = 256

	// 0x46 == 64+6 = 70 in dec

	//y val -> []byte{0x41, 0x42, 0x43, 0x44, 0x45}
	// y val -> []int32{65, 66, 67, 68, 69}
	// array to slice

	fmt.Printf("y val -> %#v\n", str)
	fmt.Printf("y val -> %#v\n", s1)

}
