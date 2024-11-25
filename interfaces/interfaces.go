package interfaces

import (
	"fmt"
)

var interfaceLessons = []func(){

}

func InterfaceLessons() {
	for _, lesson := range interfaceLessons {
		lesson()
		fmt.Println("------------------------------------------------------------------")
	}
}

// The interface defines BEHAVIOR
type Writer interface {
	Write([]byte) (int, error) // anything that writes fulfills the writer interface
}

// Interfaces are implicitly declared
type ConsoleWriter struct {}

// To implement the interface, just implement the listed behaviors
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func interfaceBasics() {
	// instead of a type here, we use an interface
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}