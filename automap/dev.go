package automap

import "fmt"

// Explore marks the whole automap as discovered.
func Explore() {
	fmt.Println("dev: automap.Explore")
	for x := 0; x < 40; x++ {
		for y := 0; y < 40; y++ {
			Discovered[x][y] = true
		}
	}
}
