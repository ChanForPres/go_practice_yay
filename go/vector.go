package main

import (
	"fmt"
	// "unsafe"
)
 
type Vector []interface{}

func init(size int) *Vector {
	
	
	
	if 	size >=0 {
		var arr1 [size]interface
		
		return &arr1
		
	} else {
		return nil
	}
	
	// return nil
}
		
	// func size() int
	// func bool add(elem interface{}) 
	// func addIndex(int, )
	// func clear()
	// func clone()
	// func contains()
	// func get()
	// func indexOf()
	// func isEmpty()
	// func remove()
	// func removeRange()
	// func toArray()


// type Vector struct {
// 	 size() int
// 	bool add() 
// 	add(int, )
// 	clear()
// 	clone()
// 	contains()
// 	get()
// 	indexOf()
// 	isEmpty()
// 	remove()
// 	removeRange()
// 	toArray()
	
// }

  