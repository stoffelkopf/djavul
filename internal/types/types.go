// Package types describes the data types used in Diablo 1.
package types

// Bool8 is an 8-bit boolean.
type Bool8 int8

// Bool32 is a 32-bit boolean.
type Bool32 int32

// A Point is an X, Y coordinate pair. The axes increase right and down.
type Point struct {
	X int32
	Y int32
}
