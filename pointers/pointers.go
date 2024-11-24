package pointers

import (
	"fmt"
	"unsafe"
)

var pointerLessons = []func(){
	pointerBasics,
	pointerCreation,
	pointerDereferencing,
	pointerArithimetic,
	pointerUnsafe,
	pointerStructs,
	pointerNew,
	pointerAccessingFields,
	pointerSlices,
	pointerMaps,
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

func pointerArithimetic() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	// the memory addresses of b and c are 8 apart
	fmt.Println("%v %p %p\n", a, b, c)
	// however, go doesnt allow you to just add to get there
	// c = &a[0] + 8
	// this wont execute. as an error occurred above
	// invalid operation: &a[1] - 8 (mismatched types *int and untyped int)
	// fmt.Println("%v %p %p\n", a, b, c)
}

func pointerUnsafe() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	// we can do pointer arithimetic in unsafe mode.
	c := unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + unsafe.Sizeof(*b))
	d := unsafe.Pointer(uintptr(unsafe.Pointer(&c)) - unsafe.Sizeof(*b))
	fmt.Println("%v %p %p %p\n", a, b, c, d)
}

func pointerStructs() {
	type myStruct struct {
		foo int
	}

	var ms *myStruct // create a pointer type
	ms = &myStruct{foo: 42}
	// this prints as:
	// &{42}
	// meaning that this pointer points to a struct
	// which has 42 as a value of one of it's fields
	fmt.Println(ms)
}

func pointerNew() {
	type myStruct struct {
		foo int
	}

	var ms *myStruct
	// ms is initially 'nil'. Pointer types that aren't initialized
	// with a value start with a value of nil, so check for that before
	// using them
	fmt.Println(ms)
	// We can also initialize this pointer to a struct
	// using the new function, but all the fields are initialize 
	// to zero values
	ms = new(myStruct)
	// this prints as:
	// &{42}
	// meaning that this pointer points to a struct
	// which has 42 as a value of one of it's fields
	fmt.Println(ms)
}

func pointerAccessingFields() {
	type myStruct struct {
		foo int
	}

	var ms *myStruct // create a pointer type
	ms = new(myStruct)

	// This is big
	// in order to access fields on a struct with
	// a pointer, you have to dereference the pointer.
	// BUT the dereferencing operator is lower on the
	// order of operations than the dot operator,
	// so you have to use paraethesis to make the 
	// dereference come before the dot notation

	(*ms).foo = 42 // first dereference, then dot access

	// this prints as:
	// &{42}
	fmt.Println(ms)

	// HOWEVER, go adds syntaxic sugar to struct pointers
	// that allows you to access fields on struct pointers
	// directly without the dereference operator or paraenthesis,
	// ever though that's what's happening under the hood
	ms.foo = 19
	fmt.Println(ms)
	fmt.Println(ms.foo)
}

func pointerSlices() {
	a := [3]int{1, 2, 3}
	// When we assign an array to another variable, we're 
	// creating a new array in memory and copying the values
	b := a
	fmt.Println(a, b)
	a[1] = 42
	// b is a new array with the values copied from a
	// when we change a, the b array is still the same
	fmt.Println(a, b)

	c := []int{1, 2, 3}
	// but when we assign a slice to another variable, we're
	// creating a copy of the slice, but the slice was a collection
	// of pointers to an underlying array, so both the old and
	// new slice have pointers that point to the same underlying array
	d := c
	fmt.Println(c, d)
	c[1] = 42
	fmt.Println(c, d)
}

func pointerMaps() {
	// Maps are similar to slices in that they are
	// collections of pointers to underlying data
	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux" // this assigns this value to a["foo"], which is a POINTER
	// thus, when we check both maps, they have both changed, because 
	// both have the same pointers to the same underlying data
	fmt.Println(a, b)
}