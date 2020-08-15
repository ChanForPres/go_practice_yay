package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

======================================================
// Comment from https://dave.cheney.net/2014/03/25/the-empty-struct#comment-2815
It’s not true that “a value must be aligned in memory to a multiple of its width.” Each type has another property, its alignment. Alignments are always powers of two. The alignment of a basic type is usually equal to its width, but the alignment of a struct is the maximum alignment of any field, and the alignment of an array is the alignment of the array element. The maximum alignment of any value is therefore the maximum alignment of any basic type. Even on 32-bit systems this is often 8 bytes, because atomic operations on 64-bit values typically require 64-bit alignment.

To be concrete, a struct containing 3 int32 fields has alignment 4 but width 12.

It is true that a value’s width is always a multiple of its alignment. One implication is that there is no padding between array elements.
======================================================


// Their struct looks like a a pad pad b b b b

//basically makes a struct{a uint16, b uint32} 
//addressable like a [2]uint32
// or whatever the largest field is


//Notice that every field starts on a
// multiple of 4 bytes from the start of
// the struct, because the alignment is 4


func main() {
	fmt.Println(0, runtime.GOOS)
	fmt.Println(1, runtime.GOARCH)
	
	fmt.Println(2, unsafe.Sizeof(struct{}{})
	
	fmt.Println(3, unsafe.Sizeof(struct {
		x uint16
	}{}))

	fmt.Println(4, unsafe.Sizeof(struct {
		x uint32
	}{}))

	fmt.Println(5, unsafe.Sizeof(struct {
		x uint64
	}{}))

	fmt.Println(6, unsafe.Sizeof(struct {
		x uint16
		y uint32
	}{}))

	fmt.Println(7, unsafe.Sizeof(struct {
		x uint16
		y uint16
		z uint32
	}{}))

	fmt.Println(8, unsafe.Sizeof(struct {
		x uint16
		y uint64
	}{}))d

	fmt.Println(9, unsafe.Sizeof(struct {
		x uint16 //2
		y uint32 //4 
		z uint64 //8 
		// 14 bytes,  aftern align = 16  bytes
		//so it appears to be packing like x y y pad z z z z or x pad y y z z z z
		//er, double each of those if we're talking bytes
		//that means the largest field size is determining the struct width, not the field alignment
			

		// 9 above looks like x x pad pad y y y y z z z z z z z z
	// Notice that all the constraints are satisfied—y starts 
	//on a multiple of 4, and z starts on a multiple of 8
 
		}{}))

	fmt.Println(10, unsafe.Sizeof(struct {
		x uint16 
		y uint64
		z uint64
		// 2 4 8 16 32
	}{}))
	
	
}	

//Well, actually what I said wasn't quite accurate, it looks like it's guaranteed to be word-aligned
// word size is 4-byte on 32-bit architectures and 8-byte on 64-bit architectures.
// This means, for the current version of the standard Go compiler, the alignment guarantees of other
//  types may be either 4 or 8, depends on different build target architectures. This is also true for gccgo.

// But my original point stands: fields are aligned, not structs
// (well, struct are aligned, but that means the entire struct is aligned at an 8-byte alignment, not the individual fields)
// The alignment of a struct only affects the struct—it doesn't change the alignment of every field.




type S struct {
        a uint16 // 2 bytes 
        b uint32 // 4 bytes
}
var s S
fmt.Println(unsafe.Sizeof(s)) // prints 8, not 6
// a a pad pad b b b b, so this doesn't change, right?
//The padding is there because b needs to be 4-byte aligned
// So it needs to start on a multiple of 4

// 

type S struct {
			a uint16 //starts on 0, pads 2
        b uint32 // starts on 4, pads 0
        c uint16 // starts on 8, pads 4
        d uint64 // starts on 16, pads 0
}
var s S
fmt.Println(unsafe.Sizeof(s)) // prints 24 as expected
// 2 4 2 8 
// a a pad pad b b b b c c pad pad pad pad pad pad d d d d d d d d

Basic types 32 bits and smaller are aligned to their size; bigger than that is aligned to word size
A 2 byte type has an alignment of 2.
Because d is 8 bytes, on 64-bit architectures it will be have an alignment of 8,
 and the struct has the alignment of its largest-aligned field

I think you still might be conflating the alignment of the fields with the alignment of the struct


Add e byte after d, and youll see the size jump all the way up to 32
// a a pad pad b b b b c c pad pad pad pad pad pad d d d d d d d d e pad pad pad pad pad pad pad


But if you add a byte after a, the size wont change at all
Because if you add it after d, the struct is no longer aligned to 8, so it needs to pad to a multiple of it, but if you add a byte after a, it doesnt affect the alignment of b, so everything is fine
   

Ther wordpadding is done for faster access to fields at runtime
 because the field offset calculation is done using words. pointer
  values are usually at a word resulution. And in compiled code the
   set of fields is a list of workd alighed pointers.  
