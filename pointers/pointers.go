package pointers

var pointerLessons = []func(){
	deferOrder,
	// deferServer,
	deferVariables,
}

func PointerLessons() {
	for _, lesson := range pointerLessons {
		lesson()
		fmt.Println("------------------------------------------------------------------")
	}
}

func pointerCreation() {

}

func pointerDereferencing() {

}