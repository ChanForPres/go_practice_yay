// Given a string and a non-negative int n, return a larger string
// that is n copies of the original string.

// stringTimes("Hi", 2) → "HiHi"
// stringTimes("Hi", 3) → "HiHiHi"
// stringTimes("Hi", 1) → "Hi"

package main

import (
	"bytes"
	"fmt"
)

func stringTimes(str string, n int) (string, bytes.Buffer) {
	if str == "" {
		buf := new(bytes.Buffer)

		return "", *buf
	}
	//var retStr
	retStr := str

	for j := 0; j < n; j++ {
		//fmt.Println(j)
		retstr += str
	}

	// skips repeated generation of intermediary string
	var retStr2 bytes.Buffer

	for j = 0; j < n; j++ {

		retStr2.WriteString(str)
	}

	return retStr, retStr2
}

func main() {
	fmt.Println("String: ", b.String())

}

func stringTimes2(str string, n int) string {

}

// 61b projs
// 61a projs
