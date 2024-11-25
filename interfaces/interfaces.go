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