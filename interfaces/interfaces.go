package interfaces

import (
	"bytes"
	"fmt"
	"io"
)

var interfaceLessons = []func(){
	interfaceBasics,
	interfaceOnOtherTypes,
	interfaceComposition,
	interfaceTypeConversion,
	interfaceConversionPanics,
	interfaceEmpty,
	interfaceSwitching,
	interfaceReferenceReceiver,
}

func InterfaceLessons() {
	for _, lesson := range interfaceLessons {
		lesson()
		fmt.Println("------------------------------------------------------------------")
	}
}

// The interface defines BEHAVIOR
// Convention for single method interfaces is to
// name the interface "{methodName}er". So here
// the method Write is the single method in the
// "Writer" interface
type Writer interface {
	Write([]byte) (int, error) // anything that writes fulfills the writer interface
}

// Interfaces are implicitly declared
type ConsoleWriter struct{}

// To implement the interface, just implement the listed behaviors
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func interfaceBasics() {
	// instead of a type here, we use an interface
	var w Writer = ConsoleWriter{}
	// So here, the type is actually unknown by the complier
	// but because the interface is declared, the complier
	// can assume w can "Write". Polymorphic behavior.
	w.Write([]byte("Hello Go!"))
	// Also note that the writer interface doesnt need
	// to know ConsoleWriter exists. ConsoleWriter, just
	// by having the behavior of the Writer interface,
	// can be used anywhere Writers are used.
}

type Incrementer interface {
	Increment() int
}

// Yhis time, we'll implement an interface on a type alias called IntCounter
// We can't add an interface to a type we don't control, so we made a
// type alias around the int type so we could make it satisfy the Incrementer
// interface
type IntCounter int

// Note here that we need to modify the integer to increment
// Incrementer, so we use a reference receiver to point to
// the original integer instead of copying the value into
// a new memory space
func (ic *IntCounter) Increment() int {
	*ic++ // ic is a pointer to the memory, so we need to deref it to modify the value
	// We dereference the pointer here to access the value
	return int(*ic)
}

func interfaceOnOtherTypes() {
	myInt := IntCounter(0)       // Cast an int to our IntCounter type alias
	var inc Incrementer = &myInt // Hmm, interesting. come back to this
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

// Declaring a new interface...
type Closer interface {
	Close() error
}

// So we can compose a new interface from
// two other interfaces. Note that Writer and
// Closer are interfaces, not methods, and that
// any WriterCloser is required to have the
// behavior of both Writers and Closers
type WriterCloser interface {
	Writer
	Closer
}

// We'll implement WriterCloser on this struct
// BufferedWriterCloser contains a pointer to
// a buffer of memory
type BufferedWriterCloser struct {
	buffer   *bytes.Buffer
	greeting string
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	// Write the byte slice into bwc
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	// Create a temporary byte slice to read chunks of the bwc into
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		// read the next 8 characters into our byte slice
		// this will write over the last 8 characters in the slice
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		// print the 8 characters on a new line
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	// Note: this for loop will cut off the last characters
	// if the length isn't a multiple of 8
	// Later note: This is in anticipation of the Close
	// implementation using buffer.Next to read the
	// last bytes. My guess is Next is less efficent than
	// read, as it has to find our how much is left in
	// the buffer, and that could be expensive to do
	// repeatedly if the buffer is large. Running Write
	// first allows the fastest operation to be done for
	// most of the print, and Close can use Next when
	// the buffer is at it's smallest

	// It's customary for writers to return the length of whatever they
	// wrote in most languages for reasons
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	// Finish writing out the last of the buffer
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func (bwc *BufferedWriterCloser) Hello() {
	fmt.Println(bwc.greeting)
}

// Initializer function (constructor function)
func NewBufferedWriterCloser() *BufferedWriterCloser {
	// note that we return a POINTER to the BufferedWriterCloser,
	// so Go is most likely taking this memory address off
	// of the stack at the return and moving it to the heap
	// so the outer context has access to the memory and
	// the function memory is freed. GC work most likely.
	return &BufferedWriterCloser{
		buffer:   bytes.NewBuffer([]byte{}),
		greeting: "Hello there!",
	}
}

func interfaceComposition() {
	// We can now use both methods specified by WriterCloser
	// Note: if there were more methods on BufferedWriterCloser,
	// we couldn't use them because we're using the WriterCloser
	// interface as the type
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello YouTube listeners, this is a test"))
	wc.Close()
}

func interfaceTypeConversion() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello YouTube listeners, this is a test"))
	wc.Close()

	fmt.Println(wc) // this prints the memory address of this data

	// wc.Hello()
	// this errors with: wc.Hello undefined (type WriterCloser has no field or method Hello)
	// This error occurs because the WriterCloser interface doesn't have a Hello method,
	// that's a method defined on the BufferedWriterCloser struct
	// In order to use the Hello method, we'll have to do a type conversion of
	// wc to a BufferedWriterCloser type
	bwc := wc.(*BufferedWriterCloser) // this is the type conversion syntax

	// type conversion syntax for interfaces can be modeled as:
	// b := a.(type)
	// where interface a is converted to b of "type" type
	// so here, wc is being type converted to a BufferedWriterCloser pointer,
	// which is what NewBufferedWriterCloser() originally returned.
	// If we were to try to convert the interface to another interface it
	// didn't actually implement, like
	// bwc := wc.(io.Reader)
	// we would get a
	// panic: interface conversion: *interfaces.BufferedWriterCloser is not io.Reader: missing method Read

	// Now, with our syntaxic sugar that allows us to access fields
	// directly from a pointer, we can run the Hello function
	bwc.Hello()
	fmt.Println(bwc) // this prints the same memory address as wc

	// If we wanted to cast our BufferedWriterCloser pointer back to an interface,
	// we could use var syntax:
	var wctwo WriterCloser = bwc
	fmt.Println(wctwo)
}

func interfaceConversionPanics() {
	// Now lets say we start exactly the same as our TypeConversion example
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello YouTube listeners, this is a test"))
	wc.Close()

	// If we do try to convert a type into an interface that the type
	// doesn't fulfill, we can decide not to panic if the program
	// can go on by returning the second item in the type conversion's
	// return's tuple, a success boolean
	r, ok := wc.(io.Reader)
	if !ok {
		fmt.Println("Conversion failed")
		return
	}
	fmt.Println(r)
}

func interfaceEmpty() {
	// An empty interface is just that, an interface with no methods
	// assigned to it. We can use it when we don't know enough about
	// the type we're receiving
	var myObj interface{} = NewBufferedWriterCloser()
	// With this syntax, we can check if the data that's under the
	// empty interface fulfills a particular interface. If it does,
	// we can use that behavior accordingly and avoid a panic
	if wc, ok := myObj.(WriterCloser); ok {
		// In this example, if myObj is a WriterCloser, we're going to
		// write to it, but if it's not we'll simply move on
		wc.Write([]byte("Hello YouTube listeners, this is a test of the emergency broadcast system"))
		wc.Close()
	}
	r, ok := myObj.(io.Reader)
	if !ok {
		// In this example, if myObj isn't a reader, that's a problem
		// big enough to stop the execution of the current function
		fmt.Println("Conversion failed")
		return
	}
	fmt.Println(r)
}

func interfaceSwitching() {
	// If you don't know what type a variable is, this is a way
	// to switch on logic for each expect4ed outcome
	var i interface{} = "0"
	switch i.(type) { // using the "type" keyword here allows us to switch on type
	case int:
		fmt.Println("i is a integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")
	}
}

type myWriterCloser struct{}

// If we use a value for myWriterCloser, like...
// var wc WriterCloser = myWriterCloser{}
// ...and any of these methods use a reference receiver,
// we'll get an error saying that myWriterCloser{} doesn't implement
// the methods, because the method has a pointer receiver
func (mwc *myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

// The method set for a value type in the context of an interface
// is all the methods with a value receiver. BUT the method
// set for a pointer type is all the value receivers AND all
// the pointer receivers
func (mwc myWriterCloser) Close() error {
	return nil
}

// So for this code to work, either all the receivers have
// to be value receivers, or we need to implement the interface
// with a pointer type in the code
func interfaceReferenceReceiver() {
	// If we used myWriterCloser{}, this code wouldn't compile
	var wc WriterCloser = &myWriterCloser{}
	fmt.Println(wc)
}

// Interface Best Practices
//////////////////////////////////////////////////////////////
// * Prefer many small interfaces as opposed to one large one
//		* io.Writer, io.Reader, even interface{}
// * Don't export interfaces for types that will be consumed by others
// 		* This allows others to define custom interfaces for types
//		  you export that fit just their needs
// * Do export interfaces for types that you import to be used by
// 	 your package
// 		* accepting an interface you've defined instead of a
//        concrete type is the best way to consume imported code
// * Where possible, functions and methods should be designed to
//	 receive interfaces
// 		* This is harder to do if you depend on the underlying fields
// 		  of a type, but if you only interact with the type's behavior
// 		  this is perfect
