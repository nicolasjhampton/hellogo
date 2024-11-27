package channels

import (
	"fmt"
)

var channelLessons = []func(){

}

func ChannelLessons() {
	fmt.Println("////////////////////////////*CHANNELS*////////////////////////////")
	for _, lesson := range channelLessons {
		lesson()
		fmt.Println("------------------------------------------------------------------")
	}
}

func channel() {
	
}