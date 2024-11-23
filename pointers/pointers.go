package pointers

import (
	"fmt"
)

var pointerLessons = []func(){
	pointerBasics,
	pointerCreation,
	pointerDereferencing,
}

func PointerLessons() {
	for _, lesson := range pointerLessons {
		lesson()
		fmt.Println("------------------------------------------------------------------")
	}
}

func pointerBasics() {
	// When we assign a to b, its a copy
	a := 42
	b := a
	fmt.Println(a, b)
	// see how changing a doesnt change b?
	// b was a copy of the value a had
	// a and b never referred to the same memory
	a = 27
	fmt.Println(a, b)
}

func pointerCreation() {
	var a int = 42
	// b is a pointer to an integer memory block
	// &a is an address for an integer memory block
	var b *int = &a
	// the value of b is the memory address of a
	fmt.Println(a, b)
}

func pointerDereferencing() {
	var a int = 42
	var b *int = &a
	// dereferencing the b pointer will give you the value at that address
	fmt.Println(a, *b)
	// Now, when a is changed, the value pointed to by b also changes
	a = 27
	fmt.Println(a, *b)
	// and when the value pointed to by b changes
	*b = 13
	// the value of a changes as well
	fmt.Println(a, *b)
	// the address of both a and *b are the same
	fmt.Println(&a, b)
	////////////////////////////////////////
	// *int is a pointer type
	// *b is a dereference of the b pointer
	// &a is the address of a
	////////////////////////////////////////
	// *int = pointer type
	// *b = dereferencing operator
	// &a = addressof operator
	////////////////////////////////////////
}

func pointerNew() {

}