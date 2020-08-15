package main

import "fmt"

func main() {
	var s1 Stringer
	s1 = Imp{"abc"}
	fmt.Println(s1.toString())

}

type Stringer interface {
	toString() string
}

type Imp struct {
	val string
}

func (i Imp) toString() string {
	return fmt.Sprintf("%#v", i.val)
}

//To “convert” an interface to a
//string, struct or map you
// should use a type assertion or
// a type switch. A type assertion
//doesn’t really convert an interface to
// another data type, but it provides access
//to an interface’s concrete value,
//which is typically what you want.

// switch on intf val to get string val

func switch(s Stringer) (Stringer, Imp){
	s := Imp{"abc"}
	s1 := s.(Imp)
	return s,s1
	
}

type