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
	"strings"

	"github.com/AllenDang/w32"
	"github.com/faiface/pixel/pixelgl"
	"github.com/sanctuary/djavul/diablo"
	"github.com/sanctuary/djavul/engine"
)

//export Start
func Start() {
	fmt.Println("djavul.Start: entry point in Go")
	cinit()
	if err := initFrontConn(); err != nil {
		log.Fatalf("%+v", err)
	}
	if engine.UseGUI {
		pixelgl.Run(run)
	} else {
		winGUI()
	}
	//l1.UseGo = false
	//dumpL1Maps()
	//if err := compareL1(); err != nil {
	//	log.Fatalf("+%v", err)
	//}
	//if err := checkL1Regular(); err != nil {
	//	log.Fatalf("%+v", err)
	//}
	//if err := checkL1Quest(); err != nil {
	//	log.Fatalf("%+v", err)
	//}
	//os.Exit(0)
	//return
}

// winGUI initializes the Windows GUI.
func winGUI() {
	inst := w32.GetModuleHandle("")
	// Parse arguments from command line.
	var s int64
	flag.Int64Var(&s, "r", 0, "initial signed 32-bit seed for dungeon generation")
	flag.Parse()
	switch {
	case s >= -2147483648 && s <= 2147483647:
		*diablo.FlagRSeed = int32(s)
	default:
		panic(fmt.Errorf("invalid seed; expected >= -2147483648 and <= 2147483647; got %d", s))
	}
	args := strings.Join(flag.Args(), " ")
	fmt.Println("args:", args)
	show := w32.SW_SHOWDEFAULT
	diablo.WinMain(inst, 0, args, show)
}

// cinit invokes cpp initialiation functions.
func cinit() {
	C.djavul_cinit()
}
