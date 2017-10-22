package init

// static void init_archives() {
//    void (*f)() = (void *)0x41ACA1;
//    f();
// }
import "C"

// Archives initializes the MPQ archives.
//
// ref: 0x41ACA1
func Archives() {
	C.init_archives()
}
