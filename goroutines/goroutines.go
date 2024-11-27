package goroutines

import (
	"fmt"
	"time"
	"sync"
	"runtime"
)

var goroutineLessons = []func(){
	goroutineCreation,
	goroutineWaitGroups,
	goroutineMutexes,
}

func GoroutineLessons() {
	fmt.Println("////////////////////////////*GOROUTINES*////////////////////////////")
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
	// or go routine.

	go sayHello()

	// The "go" keyword moved sayHello to a separate thread than this function
	// runs on, so normally this goroutineCreation function finishes before
	// sayHello gets to print. This is a hacky way of preventing that.
	time.Sleep(100 * time.Millisecond)

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

	var msg = "Hello"
	go func() { // (msg string) {
		// The go scheduler won't interupt the main thread until
		// it hits the time.Sleep call. By then msg will be changed
		// in the outer scope, so this will print "goodbye". To avoid
		// this race condition, we can pass the value of msg as
		// an argument to the goroutine so the thread has it's own
		// state separate from the outer scope on a different thread
		fmt.Println(msg)
	}()// (msg)
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)

	// Try running this with `go run -race .` to let the compiler detect this race condition
	// ==================
	// WARNING: DATA RACE
	// Read at 0x00c000114500 by goroutine 9:
	//   github.com/nicolasjhampton/hellogo/goroutines.goroutineCreation.func2()
	//       ~/source/golang/hellogo/goroutines/goroutines.go:61 +0x30

	// Previous write at 0x00c000114500 by main goroutine:
	// ==================
	// Found 1 data race(s)
	// exit status 66
}

var wg = sync.WaitGroup{}

// Instead of setting sleep timers for synchronization, we can set a
// wait group. A wait group sets an expectation for the amount of threads
// that need to resolve before moving on in the main thread. Here, we're
// using the Add method to say that we have one thread that will run
// parallel to the main thread. When the main thread hits the Wait method,
// it will start to check the count of threads that have called Done.
// Once it gets a count equal to the expected number of finished threads
// execution can continue.
func goroutineWaitGroups() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	wg.Wait()

	// time.Sleep(100 * time.Millisecond)
}

// What happens when goroutines in a wait group all access the same shared
// data?

// WaitGroup is designed to work in a global scope
var wgt = sync.WaitGroup{}
var counter = 0

// Here we create a Read Write mutex. A mutex creates rules for accessing
// a shared state. A read write mutex creates different rules for reading
// state and writing to state
var m = sync.RWMutex{}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wgt.Done()
}

func increment() {
	counter++
	m.Unlock()
	wgt.Done()
}

func goroutineMutexes() {
	// The GOMAXPROCS variable limits the number of os threads a Go program can
	// simultaneously use. This can be useful for testing concurrent code under
	// different conditions. Setting it to -1 will return the number of threads
	// the system has made available to the program
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		// Here, we add two goroutines to the wait group each time
		// But without a mutex, there's no rules on when the goroutines
		// access and write to the counter variable, and there's nothing
		// to guarantee that the goroutines will execute in the order
		// we called them in
		wgt.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
		// The new problem we've created here with the mutex is that we've
		// essentially taken the parallelism out of this program again.
		// The locks stop all actions until each thread is executed 
	}
	wgt.Wait()
}

// Goroutine Best Practices
//////////////////////////////////////////////////////////////
// * Don't create goroutines in libraries for consumers to use
// 		* Let the consumers of the library decide how to use concurrency
// * Know how a goroutine is going to end when you create one
// 		* This avoids subtle memory leaks
// * Check for race conditions at compile time
// 		* `go run -race .`


