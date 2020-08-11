// check 1 word repeated after another

package main

import "fmt"

func main() {
	//	s := make([]string, 3)
	//	s = append(s, "d")
	//	s = append(s, "e", "f")
	//	s = append(s, "e", "f")
	//
	//	c := make([]string, len(s))
	//
	//	copy(c, s)
	//fmt.Println("arr: ", s, "arr copied: ", c, "\n")
	// store results in slice like [T, F,T,F,T]

	s1 := make([]bool, 0)
	res := checkRep("aba")
	s1 = append(s1, res)

	fmt.Println("RESULT: ", res)
	fmt.Printf("results arr: %#v \n", s1)

	res = checkRep("aa")

	s1 = append(s1, res)
	fmt.Printf("results arr: %#v \n", s1)
	fmt.Println("RESULT: ", res)
}

func checkRep(str string) bool {
	// check even or empty string
	strlen := len(str)
	if (str == "") || (strlen%2 == 1) {
		return false
	}
	startPt := len(str) / 2 // ie 6/2 4/2 2/2

	// func Split(s, sep string) []string
	//	fmt.Printf("%#v\n",strings.Split("abc", ""))

	for i := 0; i < startPt; i++ {

		// divide by 2
		if str[i] != str[startPt+i] {
			fmt.Println("\ni==str[+i] => ", str[i], "==", str[startPt+i])

			return false
		}

		// example strings
		//"ab a"

		// indexes: 012 345
		//"abcabc"

		//"axax"
		//"ax"
		//"aa"
		//"amc"

	}
	return true

	// return false

}

// Use a conversion to runes, for example
// works for unicode chars
// if the string is empty, the result is []rune(nil).

// 			{
//         s := "Hello, 世界"
//         for i, r := range s {
//                 fmt.Printf("i%d r %c\n", i, r)
//         }
//         fmt.Println("----")
//         a := []rune(s)
//         for i, r := range a {
//                 fmt.Printf("i%d r %c\n", i, r)
//         }
