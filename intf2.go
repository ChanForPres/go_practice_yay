package main

import (
	"bytes"
	"fmt"
)

// func buf() {
// 	// buffer bytes
// 	a := new(bytes.Buffer)
// 	b := new(strings.Builder)
// 	a.Grow(2)
// 	a.Write()
// }

func main() {
	// greetingStr, ok := greeting.(string)
	// 2nd val is bool success/fail

	// p1 := person{"a", 1}
	// p2 := child{"b", 2}
	// p3 := pet{"c", 3, "z"}

	// p4 := person(p2)
	// p5 := child(p1)

	var i1 interface{} = person{"!", 1}
	var suc bool
	val, suc := i1.(person)
	fmt.Printf("%#v \t %#v", val, suc)

	// check struct type implements an intf

	type S struct{}
	type I interface{}
	var i2 I1 = S{}
	fmt.Printf("\n%#v \t \n", i2)

	var i3 I = S{}
	fmt.Printf("\n%#v \t \n", i3)
	// val, suc := person.age.(int)
	// p6, suc := p2.(person)
	// p7, suc := p1.(child)

	// verify ptr to struct implements intf
	type S2 struct{}
	type I3 interface{}
	var s2 *S2
	s2 = new(S2)
	// To be concrete, a struct containing 3 int32 fields has alignment 4 but width 12.
	fmt.Printf("\n%p \t \n", s2)
	s2 = nil
	s2 = &S2{}
	fmt.Printf("\n%p \t \n", s2)

	v := make(S2{}, 1)
	v1 := make(S2{}, 1)
	// b := make([]struct{}, 20)
	fmt.Println(&v == &v1)

	// Allocate enough memory to store a bytes.Buffer value
	// and return a pointer to the value's address.
	var buf bytes.Buffer
	p := &buf
	
	// the alignment of uint16 is 2 bytes, 
	//the alignment of uint32 is 4 bytes,
	// so the overall alignment of the struct
	// is 4 bytes, which gives you width 8 to fit 6 bytes
	

	// Use a composite literal to perform allocation and
	// return a pointer to the value's address.
	p = &bytes.Buffer{}

	// Use the new function to perform allocation, which will
	// return a pointer to the value's address.
	p = new(bytes.Buffer)
	//var i3 I3 = *s2{}
	fmt.Printf("\n%p \t \n", p)

}

type I1 interface {
	//

}

type person struct {
	name string
	age  int
}

type child struct {
	name string
	age  int
}

type pet struct {
	name string
	age  int
	a    string
}

// myInt is a new type who's base type is `int`
type myInt int

// The AddOne method works on `myInt` types, but not regular `int`s
func (i myInt) AddOne() myInt { return i + 1 }

// func main() {
// 	bob := person{
// 		name: "bob",:7

// 		age: 15,
// 		}

//   babyBob := child(bob)
//   // "babyBob := pet(bob)" would result in a compilation error
// 	fmt.Println(bob, babyBob)

// 		var i myInt = 4
// 	fmt.Println(i.AddOne())
// }
