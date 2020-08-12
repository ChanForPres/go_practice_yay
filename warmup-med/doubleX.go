// Given a string, return true if the first instance of "x" in the string is immediately followed by another "x".

// doubleX("axxbb") → true
// doubleX("axaxax") → false
// doubleX("xxxxx") → true

package main

import "fmt"

type ExStruct struct {
}

func checkNilIntf() bool {
	// convert to nil ptr
	var empStruct *ExStruct // ptr to struct
	var inf interface{}

	inf = empStruct // assigning nil ptr to intf
	// check intf is nil ptr

	// switch on nil
	res, ok := inf.(nil)
	res, ok = inf.((*ExStruct)(nil)) // ptr to nil intf

	// If T is not an interface, it asserts that the dynamic type of x is identical to T.
	// If T is an interface, it asserts that the dynamic type of x implements T.

	if (*ExStruct)(nil) == inf {
		return true
	}
	return false
}

func main() {
	// using regex
	// split stringby x
	//fmt.Printf("%#v\n",strings.Split("abc", ""))

	// string is slice of bytes in ASCII, not necessarily UTF8/Unicode
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// iterate using runes
	//fmt.Printf(" \"Hello\"[1]: %c", "HELLO"[1], "\n")
	fmt.Println("HELLO[1:2] ->", "HELLO"[1:2])
	fmt.Println("HELLO[1:2] ->", "HELLO"[1:2])
	fmt.Println("HELLO[1] -> ", "HELLO"[1])

	fmt.Println("string(HELLO[1]) -> ", string("Hello"[1]))

	fmt.Printf("[]rune(HELLO)[1] -> %c", []rune("HELLO")[1])
	// var x = 'c' // rune
	// var str = "st"
	// check using type assertion
	var intf1 interface{} = "foo"
	
	fmt.Println(intf1.(string))

	// recur
	// loop

}

// //func  {
// 	if == ''
// }

// string is slice of bytes

//
