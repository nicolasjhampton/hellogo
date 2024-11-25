package functions

import (
	"fmt"
)

var functionLessons = []func(){
	functionSyntax,
	functionParameters,
	functionVariadicParameters,
	functionReturn,
	functionReturns,
}

func FunctionLessons() {
	for _, lesson := range functionLessons {
		lesson()
		fmt.Println("------------------------------------------------------------------")
	}
}

// A function declaration has to have the func keyword and the paranthesis
func functionSyntax() {
	fmt.Println("Hello, playground")
}

// msg is the parameter for the function, and it is typed
// idx is the second parameter, separated by a comma
func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

// If both parameters are the same type, you can save space
// by using the same type declaration for both parameters
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}

func sayGreetingTwo(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Peter"
	fmt.Println(name)
}

func sayGreetingThree(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Peter"
	fmt.Println(*name)
}

func functionParameters() {
	for i := 0; i < 5; i++ {
		// using the parameter of the function, we can pass
		// values into the function for it to use at runtime
		sayMessage("Hello Go!", i)
	}

	sayGreeting("Hello", "Gwen")

	// Note here that we're passing in values, not pointers
	// and the parameters get copies of those values
	greeting := "hello"
	name := "Stacey"
	sayGreetingTwo(greeting, name)
	// So any changes to the values of the parameters
	// inside the function doesn't affect the variables
	// outside the function
	fmt.Println(name)

	// But if I pass POINTERS as arguments, then the function
	// is given the address of the original data, and can modify
	// that data outside of the function
	sayGreetingThree(&greeting, &name)
	// This will now print "Peter" like it did inside the function
	fmt.Println(name)

	// This also means that less memory is used in this program,
	// as we didn't make copies of the values in memory
}

// The result of the elipses is a slice of all arguments given to the
// function as a parameter
func sum(msg string, values ...int) { // can only be one variadic param, and has to be last
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

func functionVariadicParameters() {
	sum("The sum is", 1, 2, 3, 4, 5) // similar to sum("The sum is", []int{1,2,3,4,5})
}

// In this version, instead of printing the value within
// the function, we'll return it
func sumTwo(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// In this version, we dont return the value,
// instead we return a pointer to the value, thus
// preventing the return from making what could be
// an expensive copy of a value
//
// The instructor make a great point in that
// in other languages, this pointer would refer
// to memory that no longer exists after the
// function was finished executing, resulting in
// a reference to freed memory and a bug
//
// Go's garbage collector recognizes that you need
// this return value after the stack execution, and
// moves the memory from the stack to the heap for you
func sumThree(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result // note we return with addressof syntax, but the type is pointer int
}

// Go also has named return values. These are
// instansiated with the default zero value for their 
// type and are returned at the return statement
func sumFour(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

func functionReturn() {
	// we'll store the returned value in a variable to use it
	s := sumTwo(1, 2, 3, 4, 5)
	fmt.Println("The sum returned as", s)

	// t is inferred to be a pointer
	t := sumThree(1, 2, 3, 4, 5)
	// in order to get the value, we have to dereference it
	fmt.Println("The sum returned as a pointer to the value", *t)

	r := sumFour(1, 2, 3, 4, 5)
	fmt.Println("The named return value is", r)
}

// Division can result in an error if the denominator
// is zero. Not the end of the world or worth panicking over,
// but needs to be handled. 
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func functionReturns() {
	// In go, errors are often handled by
	// returning a tuple of an expected result and a possible 
	// error, thus returning two values from the function instead
	// of one.
	c, err := divide(1, 0)
	// we can check for the presence of the error value
    // to inform the direction of execution
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c)
}

func functionAnon() {

}
func functionTypes() {

}
func functionMethods() {

}

