// Given a string, compute recursively a new
// string where all the adjacent chars are now separated by a "*".

// allStar("hello") → "h*e*l*l*o"
// allStar("abc") → "a*b*c"
// allStar("ab") → "a*b"

package main

import "fmt"

func allStar(str string) string {
	// tmp := str[0]
	if len(str) < 2 {
		return str
	}
	
	str2 := string(str[0]) + "*" + allStar(string(str[1]))
	fmt.Println(str2)
	return str2
}

// abc
// bc
// c

// *c
// b*c

func main() {
	fmt.Println("returned -------->", allStar("abc"))

}

// public String allStar(String str) {
//     + allStar(str)

// }
