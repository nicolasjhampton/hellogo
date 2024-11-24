package functions

import (
	"fmt"
)

var functionLessons = []func(){
	
}

func FunctionLessons() {
	for _, lesson := range functionLessons {
		lesson()
		fmt.Println("------------------------------------------------------------------")
	}
}