package deferPanicRecover

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var deferLessons = []func(){
	deferOrder,
	// deferServer,
	deferVariables,
}

func DeferLessons() {
	for _, lesson := range deferLessons {
		lesson()
		fmt.Println("------------------------------------------------------------------")
	}
}

func deferOrder() {
	// start end middle
	// deferred functions are ran AFTER main
	// but BEFORE the main return
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}

func deferServer() {
	// defer solves the problem of dangling resources
	// forgetting to close resources can introduce bugs in code
	// if you can write your intention to close the resource at
	// the beginning, then you don't have to tracl down your last
	// use later

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// res.Body.Close()
	defer res.Body.Close()
	robots, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func deferVariables() {
	// defer takes the value of variables at the time of deferment,
	// not the values at the time of the function's end
	// reminds me of a closure
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

 // defer statements are executed last in first out, like a stack
