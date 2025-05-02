package main

import (
	"fmt"
)

func main() {
	a := "string"
	// the & for address of the pointer
	testPointer(&a)
	fmt.Printf("a: %s\n", a)
}

// This starts as a copy and need to pass the pointer
func testPointer(a *string) {
	// Need to assign the a to the pointer memory location 
	*a = "another string"
}