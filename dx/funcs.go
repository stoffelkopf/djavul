// Function wrappers for dx.cpp

package dx

// #include <ddraw.h>
//
// void dx_init(HWND hWnd) {
//    void (*f)(HWND) = (void*)0x4153A0;
//    f(hWnd);
// }
//
// void dx_create_back_buffer(void) {
//    void (*f)(void) = (void*)0x4154B5;
//    f();
// }
//
// void dx_create_primary_surface(void) {
//    void (*f)(void) = (void*)0x4155C2;
//    f();
// }
//
// HRESULT dx_DirectDrawCreate(GUID *guid, IDirectDraw **dd, void *unknown) {
//    HRESULT (*f)(GUID *, IDirectDraw **, void *) = (void*)0x41561A;
//    return f(guid, dd, unknown);
// }
//
// void dx_lock_mutex(void) {
//    void (*f)(void) = (void*)0x41569A;
//    f();
// }
//
// void dx_unlock_mutex(void) {
//    void (*f)(void) = (void*)0x415725;
//    f();
// }
//
// void dx_cleanup(void) {
//    void (*f)(void) = (void*)0x4157A0;
//    f();
// }
//
// void dx_reinit(void) {
//    void (*f)(void) = (void*)0x415848;
//    f();
// }
import "C"
import "unsafe"

// Init initializes the DirectX rendering system.
//
// ref: 0x4153A0
func Init(hwnd C.HWND) {
	C.dx_init(hwnd)
}

// CreateBackBuffer creates the DirectDraw back buffer.
//
// ref: 0x4154B5
func CreateBackBuffer() {
	C.dx_create_back_buffer()
}

// CreatePrimarySurface creates the primary DirectDraw surface.
//
// ref: 0x4155C2
func CreatePrimarySurface() {
	C.dx_create_primary_surface()
}

// DirectDrawCreate creates the DirectDraw system.
//
// ref: 0x41561A
func DirectDrawCreate(guid *C.GUID, dd **C.IDirectDraw, unknown unsafe.Pointer) C.HRESULT {
	return C.dx_DirectDrawCreate(guid, dd, unknown)
}

// LockMutex locks the DirectX mutex.
//
// ref: 0x41569A
func LockMutex() {
	C.dx_lock_mutex()
}

// UnlockMutex unlocks the DirectX mutex.
//
// ref: 0x415725
func UnlockMutex() {
	C.dx_unlock_mutex()
}

// Cleanup terminates the DirectX rendering system.
//
// ref: 0x4157A0
func Cleanup() {
	C.dx_cleanup()
}

// ReInit reinitializes the DirectX rendering system.
//
// ref: 0x415848
func ReInit() {
	C.dx_reinit()
}
