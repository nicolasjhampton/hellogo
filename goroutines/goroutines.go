package goroutines

import (
	"fmt"
)

var goroutineLessons = []func(){
	goroutineCreation,
}

func GoroutineLessons() {
	for _, lesson := range goroutineLessons {
		lesson()
		fmt.Println("------------------------------------------------------------------")
	}
}

func goroutineCreation() {
	sayHello := func() {
		fmt.Println("Hello")
	}

	// here we're invoking the function as a function on the main thread
	sayHello()

	// Here we're invoking the function as a goroutine on it's own "green" thread,
	// or go routine
	go sayHello()
	// Most languages split off threads using os threads with their own call
	// stack and resources, which can be expensive (think of thread pools in ruby)
	//
	// Go's goroutines are not os threads, but a high level abstraction of threads
	// using a scheduler to give each goroutine time on os threads without
	// managing those os threads ourselves
	//
	// Because the scheduler is managing os thread separate from the goroutines,
	// the goroutines can share os thread time, making the goroutines very
	// cheap to use
}

func goroutineSynchronization() {

}

func goroutineWaitGroups() {

}

func goroutineMutexes() {

}

func goroutineParallelism() {

}
