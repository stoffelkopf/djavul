package main

import "C"
import "fmt"

//export Dddddddddddddddddddd
func Dddddddddddddddddddd() {
	fmt.Println("hello from Go :)")
}

//export Cccccccccccccccccccccccccc
func Cccccccccccccccccccccccccc() int {
	return 42
}

//export Bbbbbbbbbbbbb
func Bbbbbbbbbbbbb(key int) {
	fmt.Println("key press:", key)
}

func main() {}
