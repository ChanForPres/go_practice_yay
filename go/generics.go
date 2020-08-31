package main

type Container []interface{}

// Put adds an element to the container.	
func (c *Container) Put(elem interface{}) {
	*c = append(*c, elem)
}

// Get gets an element from the container.	
func (c *Container) Get() interface{} {
	elem := (*c)[0]
	*c = (*c)[1:]
	return elem
}

// The calling code does the type assertion when retrieving an element.	
func assertExample() {
	intContainer := &Container{}
	intContainer.Put(7)
	intContainer.Put(42)
	elem, ok := intContainer.Get().(int) // assert that the actual type is int
	if !ok {
		fmt.Println("Unable to read an int from intContainer")
	}
	fmt.Printf("assertExample: %d (%T)\n", elem, elem)
}