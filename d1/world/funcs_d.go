//+build djavul

package world

// #include <stdint.h>
//
// static void __fastcall world_draw_upper_screen(uint8_t *dst_buf) {
//    void (__fastcall *f)(uint8_t *) = (void *)0x4652C5;
//    f(dst_buf);
// }
//
// static void __fastcall world_draw_lower_screen(uint8_t *dst_buf) {
//    void (__fastcall *f)(uint8_t *) = (void *)0x46886B;
//    f(dst_buf);
// }
import "C"

import (
	"unsafe"

	"github.com/sanctuary/djavul/d1/engine"
)

const (
	// useGo specifies whether to use the Go implementation.
	useGo = true
)

// ref: 0x4652C5
func DrawUpperScreen(dstBuf unsafe.Pointer) {
	if engine.UseGUI {
		drawUpperScreen(dstBuf)
	}
	C.world_draw_upper_screen((*C.uint8_t)(dstBuf))
}

// ref: 0x46886B
func DrawLowerScreen(dstBuf unsafe.Pointer) {
	if engine.UseGUI {
		drawLowerScreen(dstBuf)
	}
	C.world_draw_lower_screen((*C.uint8_t)(dstBuf))
}
