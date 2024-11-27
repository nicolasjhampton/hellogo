package channels

import (
	"fmt"
	"sync"
)

var channelLessons = []func(){
	channelBasics,
	channelForAsync,
	channelRestrictions,
	channelBuffered,
	channelRange,
	channelCloseCheck,
}

func ChannelLessons() {
	fmt.Println("////////////////////////////*CHANNELS*////////////////////////////")
	for _, lesson := range channelLessons {
		lesson()
		fmt.Println("------------------------------------------------------------------")
	}
}

// Need a waitgroup for flow control of the outer scope
var wg = sync.WaitGroup{}

// Channels are designed to allow you to pass data between goroutines
// safely while avoiding race conditions
func channelBasics() {
	// This creates a channel that can ONLY pass integers from goroutine to goroutine
	ch := make(chan int)
	wg.Add(2)
	// receiving goroutine
	go func() {
		i := <- ch // receiving data from our channel
		fmt.Println(i)
		ch <- 33 // channel goes both ways
		wg.Done()
	}()
	// sending goroutine
	go func() {
		i := 42
		ch <- i // Sending data into our channel
		i = 27
		fmt.Println(i) // This i won't affect the value we sent in the channel
		fmt.Println(<- ch) // for every send, we need a matching receive IF THE CHANNEL IS UNBUFFERED
		wg.Done()
	}()
	wg.Wait()
}

func channelForAsync() {
	// We'll spawn 10 goroutines, and all of them will use this one channel
	// Also, we've specified no buffer, so no message can be stored in the
	// channel waiting for a receiver
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		// if this goroutine was outside of the for loop, we would only have one receiver
		// for the 5 senders, 4 messages wouldnt have anywhere to go, and we would have a 
		// deadlock error
		go func() {
			i := <- ch // this goroutine will wait for a value to come from this channel
			fmt.Println(i)
			wg.Done()
		}()
		// this goroutine can take as much time as it needs to execute
		go func() {
			ch <- 42 + j // this goroutine will pause here until it can send it's message to a receiver
			wg.Done()
		}()
	}
	wg.Wait()
}

// restricting data flow direction on a go channel makes it much easier
// to reason about the program
func channelRestrictions() {
	ch := make(chan int)
	wg.Add(2)
	// receiving goroutine
	go func(ch <- chan int) { // syntax for a receive only channel parameter
		i := <- ch // receiving data from our channel
		fmt.Println(i)
		// ch <- 33 // this is a receive only channel in this scope
		wg.Done()
	}(ch) // passing the channel in as an argument to restrict it's usage
	// sending goroutine
	go func(ch chan <- int) { // syntax for a send only channel parameter
		i := 42
		ch <- i // Sending data into our channel
		// fmt.Println(<- ch) // this is a send only channel in this scope
		wg.Done()
	}(ch) // passing the channel in as an argument to restrict it's usage
	wg.Wait()
}

func channelBuffered() {
	ch := make(chan int, 1) // A buffer will allow x amount of messages to be held in the channel
	wg.Add(2)
	go func(ch <- chan int) {
		fmt.Println(<- ch)
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		ch <- 42
		ch <- 27 // We send two values, one is consumed by our goroutine, and one is stored in the channel
		// this program will execute and print 42, but 27 will never print
		wg.Done()
	}(ch)
	wg.Wait()
	// if our senders were producing data faster than our receivers, the buffer
	// would allow our receviers to become x amount of messages behind
}

func channelRange() {
	ch := make(chan int, 1)
	wg.Add(2)
	go func(ch <- chan int) {
		// This range will continue to wait for values from the channel
		for i := range ch { // If we don't close the channel, this will wait forever, causing deadlock err
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		ch <- 42
		// close(ch) If we close the channel early, we'll get a panic when we try to use the channel again
		// it is not possible to check if a channel has been closed on the sending side, so be careful
		ch <- 27 
		close(ch) // We have to close the channel to let the range loop know to end
		wg.Done()
	}(ch)
	wg.Wait()
}

func channelCloseCheck() {
	ch := make(chan int, 1)
	wg.Add(2)
	go func(ch <- chan int) {
		for {
			// Although we can't check if a channel is closed on the sending end,
			// we can use an if statement to check if the channel is closed from
			// the receiving end
			if i, ok := <- ch; ok {
				fmt.Println(i)
			} else {
				// this could be useful for any final or cleanup logic
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		ch <- 42
		// close(ch) If we close the channel early, we'll get a panic when we try to use the channel again
		// it is not possible to check if a channel has been closed on the sending side, so be careful
		ch <- 27 
		close(ch) // We have to close the channel to let the range loop know to end
		wg.Done()
	}(ch)
	wg.Wait()
}
