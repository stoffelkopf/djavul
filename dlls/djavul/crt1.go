//+build djavul

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
//       f();
//    }
// }
import "C"

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AllenDang/w32"
	"github.com/sanctuary/djavul/diablo"
)

//export Start
func Start() {
	fmt.Println("djavul.Start: entry point in Go")
	cinit()
	if err := compareL1(); err != nil {
		log.Fatalf("+%v", err)
	}
	//if err := checkL1Regular(); err != nil {
	//	log.Fatalf("%+v", err)
	//}
	//if err := checkL1Quest(); err != nil {
	//	log.Fatalf("%+v", err)
	//}
	os.Exit(0)
	return
	inst := w32.GetModuleHandle("")
	// Parse arguments from command line.
	flag.Parse()
	args := strings.Join(flag.Args(), " ")
	fmt.Println("args:", args)
	show := w32.SW_SHOWDEFAULT
	diablo.WinMain(inst, 0, args, show)
}

// cinit invokes cpp initialiation functions.
func cinit() {
	C.djavul_cinit()
}
