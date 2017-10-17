package main

// #include <stdio.h>
//
// static void djavul_cinit() {
//    void (**djavul_cpp_init_funcs)(void) = (void *) 0x483000;
//    for (int i = 1; i < 34; i++) {
//       void (*f)(void) = djavul_cpp_init_funcs[i];
//       if (f == NULL) {
//          break;
//       }
//       printf("cinit: %p\n", f);
//       f();
//    }
// }
import "C"

import (
	"os"

	"github.com/AllenDang/w32"
	"github.com/sanctuary/djavul/diablo"
)

//export Start
func Start() {
	cinit()
	inst := w32.GetModuleHandle("")
	// TODO: Parse arguments from command line.
	args := ""
	show := w32.SW_SHOWDEFAULT
	status := diablo.WinMain(inst, 0, args, show)
	os.Exit(status)
}

func cinit() {
	C.djavul_cinit()
}
