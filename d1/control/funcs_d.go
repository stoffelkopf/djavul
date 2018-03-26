//+build djavul

package control

// #include <stdint.h>
//
// static void __fastcall control_draw_panel(int panel_x, int panel_y, uint16_t width, uint16_t height, int screen_x, int screen_y) {
//    void (__fastcall *f)(int, int, uint16_t, uint16_t, int, int) = (void *)0x404259;
//    f(panel_x, panel_y, width, height, screen_x, screen_y);
// }
import "C"

import (
	"github.com/sanctuary/djavul/d1/engine"
)

const (
	// useGo specifies whether to use the Go implementation.
	useGo = true
)

// DrawPanel renders the control panel onto screen.
//
// ref: 0x404259
func DrawPanel(panelX, panelY, width, height, screenX, screenY int) {
	if engine.UseGUI {
		drawPanel(panelX, panelY, width, height, screenX, screenY)
	}
	C.control_draw_panel(C.int(panelX), C.int(panelY), C.uint16_t(width), C.uint16_t(height), C.int(screenX), C.int(screenY))
}
