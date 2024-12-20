package main

import (
	"fmt"

	"github.com/nicolasjhampton/hellogo/deferPanicRecover"
	"github.com/nicolasjhampton/hellogo/functions"
	"github.com/nicolasjhampton/hellogo/goroutines"
	"github.com/nicolasjhampton/hellogo/interfaces"
	"github.com/nicolasjhampton/hellogo/pointers"
	"github.com/nicolasjhampton/hellogo/channels"
)

func main() {
	// previousChapters()
	deferPanicRecover.DeferLessons()
	deferPanicRecover.PanicLessons()
	deferPanicRecover.RecoverLessons()

	pointers.PointerLessons()

	functions.FunctionLessons()

	interfaces.InterfaceLessons()

	goroutines.GoroutineLessons()

	channels.ChannelLessons()

}

// type Doctor struct {
// 	number     int
// 	actorName  string
// 	episodes []string
// 	companions []string
// }

// type Animal struct {
// 	Name   string `required max:"100"`
// 	Origin string
// }

// type Bird struct {
// 	Animal
// 	SpeedKPH float32
// 	CanFly   bool
// }

func previousChapters() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i)
		fmt.Println(j)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)

	// while loop
	for i < 5 {
		fmt.Println(i)
		i++
	}
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
	i = 0
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println(name)
	s := []string{"top", "middle", "bottom"}
	for k, v := range s {
		fmt.Printf("%v: %v", k, v)
		fmt.Println("")
	}
	// bird := map[string]string{
	// 	"name":   "Emu",
	// 	"origin": "Australia",
	// }
	// if pop, ok := bird["origin"]; ok {
	// 	fmt.Println(pop)
	// }
	// t := reflect.TypeOf(Animal{})
	// field, ok := t.FieldByName("Name")
	// fmt.Println(field, field.Tag, ok)
	// bird := Animal{
	// 	Name:   "Emu",
	// 	Origin: "Australia",
	// }
	// fmt.Println(bird)
	// b := Bird{
	// 	Animal: Animal{
	// 		Name:   "Emu",
	// 		Origin: "Australia",
	// 	},
	// 	SpeedKPH: 48,
	// 	CanFly:   false,
	// }
	// b := Bird{}
	// b.Name = "Emu"
	// b.Origin = "Australia"
	// b.SpeedKPH = 48
	// b.CanFly = false
	// fmt.Println(b)
	// aDoctor := struct{ name string }{name: "John Perwee"}
	// anotherDoctor := aDoctor
	// anotherDoctor.name = "Tom Baker"
	// aDoctor := Doctor{
	// 	number:    3,
	// 	actorName: "Jon Pertwee",
	// 	companions: []string{
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 		"Sarah Jane Smith",
	// 	},
	// }
	// fmt.Println(aDoctor)
	// fmt.Println(anotherDoctor)
}
