//+build djavul

package scrollrt

// void scrollrt_draw_main_W(void) {
//    void (*f)(void) = (void*)0x4564F9;
//    f();
// }
import "C"

// DrawMainW renders the UI and the game world on screen.
//
// ref: 0x4564F9
func DrawMainW() {
	C.scrollrt_draw_main_W()
}
