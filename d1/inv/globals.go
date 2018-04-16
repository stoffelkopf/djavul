// Global variables of inv.cpp.

package inv

import (
	"github.com/sanctuary/djavul/internal/types"
)

// --- [ .rdata section ] ------------------------------------------------------

// Read-only global variables.
var (
	// ScreenPos maps from inventory slot to screen position. The inventory slots
	// are arranged as follows:
	//
	//                             00 01
	//                             02 03   06
	//
	//                 07 08       19 20       13 14
	//                 09 10       21 22       15 16
	//                 11 12       23 24       17 18
	//
	//                    04                   05
	//
	//                 25 26 27 28 29 30 31 32 33 34
	//                 35 36 37 38 39 40 41 42 43 44
	//                 45 46 47 48 49 50 51 52 53 54
	//                 55 56 57 58 59 60 61 62 63 64
	//
	//    65 66 67 68 69 70 71 72
	//
	// ref: 0x47AE60
	//
	// References:
	//    * https://raw.githubusercontent.com/sanctuary/graphics/master/inventory.png
	ScreenPos *[73]types.Point
)

// --- [ .data section ] -------------------------------------------------------

// Read-write global variables.
var (
	// StartSlot2x2 specifies the starting inventory slot for placement of 2x2
	// items.
	//
	// ref: 0x48E9A8
	StartSlot2x2 *[10]int32
)

// --- [ .bss section ] --------------------------------------------------------

// Uninitialized global variables.
var ()
