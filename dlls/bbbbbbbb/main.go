package main

import "C"

import "fmt"

//export B1_BBBBBBBB
func B1_BBBBBBBB() {
	fmt.Println("hello from lib2")
}

//export B2_BBBBBBBBBBB
func B2_BBBBBBBBBBB() {
}

//export B3_BBBBBBBBBBBBB
func B3_BBBBBBBBBBBBB() {
}

//export B4_BBBBBBBBBB
func B4_BBBBBBBBBB() {
}

//export B5_BBBBBBBBB
func B5_BBBBBBBBB() {
}

func main() {}
