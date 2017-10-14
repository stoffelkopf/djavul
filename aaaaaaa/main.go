package main

import "C"
import (
	"fmt"

	"github.com/sanctuary/djavul/automap"
)

//export InitAAAAAAAAAAAAAAAA
func InitAAAAAAAAAAAAAAAA() {
	// Called from WinMain.
	fmt.Println("hello from Go :)")
}

//export Cccccccccccccccccccccccccc
func Cccccccccccccccccccccccccc() int {
	return 42
}

//export OnKeyPressAAA
func OnKeyPressAAA(key int) {
	// Called from on_key_press.
	fmt.Println("key press:", key)
	switch key {
	case 'X':
		// eXplore
		automap.Explore()
	}
}

func main() {}
