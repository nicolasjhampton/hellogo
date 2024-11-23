package deferPanicRecover

import (
	"fmt"
	"net/http"
)

var panicLessons = []func(){
	// panicDivision,
	// panicWebHandler,
	// panicWithDefer,
}

func PanicLessons() {
	for _, lesson := range panicLessons {
		lesson()
		fmt.Println("------------------------------------------------------------------")
	}
}

func panicDivision() {
	// runtime generates panic here
	// panic: runtime error: integer divide by zero
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
}

func panicWebHandler() {
	// this will panic if you run it twice
	// Go libraries are rarely going to panic
	// they will give you an error and let you decide
	// if that error is fatal in your use case.
	// if it is, then you can decide to panic
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func panicWithDefer() {
	// panics happen AFTER defer statements run
	// that way, deferred resources are closed even
	// in the event of a panic
	fmt.Println("start")                   // this displays
	defer fmt.Println("this was deferred") // this is deferred
	panic("something bad happened")        // we hit the panic
	fmt.Println("end")                     // never runs
	// The defer is ran
	// Then the panic returns and bubbles up
}
