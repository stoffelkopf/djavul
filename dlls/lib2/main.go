package main

import "C"

import "fmt"

//export F1
func F1() {
	fmt.Println("hello from lib2")
}

//export F2
func F2() {
}

//export F3
func F3() {
}

//export F4
func F4() {
}

//export F5
func F5() {
}

func main() {}
