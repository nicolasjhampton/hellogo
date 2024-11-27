package deferPanicRecover

import (
	"fmt"
	"log"
)

var recoverLessons = []func(){
	recoverUse,
	recoverPanicker,
}

func RecoverLessons() {
	fmt.Println("////////////////////////////*RECOVER*////////////////////////////")
	for _, lesson := range recoverLessons {
		lesson()
		fmt.Println("------------------------------------------------------------------")
	}
}

func recoverUse() {
	// Because a deferred function executes at
	// at the end of a function call but before
    // the return value or final panic, we can
	// use the recover keyword to "catch" a panic		
	fmt.Println("start")                   // this displays
	defer func() {						   // this anon func is deferred
		if err := recover(); err != nil {  // ...the panic is recovered here
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")        // we hit the panic, but...
	fmt.Println("end")                     // never runs
}

func fakeMain() {
	fmt.Println("start")
	panicker()
	fmt.Println("end") // Even though we panicked, we recovered, and were able to continue in the call stack
}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking") // This will never run, as it's inside a function that panicked
}

func recoverPanicker() {
	fakeMain()
}