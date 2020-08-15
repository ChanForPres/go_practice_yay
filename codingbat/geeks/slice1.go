package geeks

import "fmt"

func F4() {
	s1 := make([]int, 0, 0)
	s1 = append(s1, 1,1,1,1,1)
	fmt.Println(len(s1),cap(s1))
	
}	

func F3() {
	s1 := make([]byte, 10)
	s2 := make([]byte, 0, 10)
	s2 = append(s2, 1)
	s2 = append(s2, '\x31')
	a := s2[0]
	fmt.Println(s1, s2)
	fmt.Printf("%x == %#v", s2[0], s2[0])
	fmt.Printf("%x == %#v == %T == %d == %#U\n", a, a, a, a, a)
	
	fmt.Printf("%x == %#v", s2[0], s2[0])
	fmt.Printf("%x == %#v == %T == %d == %#U\n", s2, s2, s2, s2, s2)

	a2 := []byte(("123"))
	fmt.Printf("%x == %#v == %T == %d == %#U\n", a2, a2, a2, a2, a2)

	a1 := []byte(string("123"))
	fmt.Printf("%x == %#v == %T == %d == %#U\n", a1, a1, a1, a1, a1)

}

// Because our ReverseRunes function begins with an upper-case letter,
// it is exported, and can be used in other packages that import our morestrings package.
func F1() {
	fmt.Println("geeks")
}

func F2() {

	s := []int{0, 1, 2}
	n := copy(s, s[1:]) // n == 2, s == []int{1, 2, 2}
	// // Copy from a string to a byte slice (special case)
	var b = make([]byte, 5)
	copy(b, "Hello, world!") // b == []byte("Hello")
	fmt.Println(n)
}

func Sl1() {
	// a := []byte("\x65\x65\x65\x65")
	a := []byte("\x61")

	b := make([]byte, len(a))

	copy(b, a)
	fmt.Println(a)
	for i, _ := range a {
		fmt.Printf("%x == %#v == %T == %d == %#U\n", a[i], a[i], a[i], a[i], a[i])
		// fmt.Printf("%d\n", j)
	}
	c := a[:]
	if true {
		fmt.Println("equals")
		fmt.Println(c)
	}
}
