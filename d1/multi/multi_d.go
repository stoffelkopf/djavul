//+build djavul

package multi

import (
	"fmt"
	"image"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/sanctuary/djavul/d1/diablo"
	"github.com/sanctuary/djavul/internal/proto"
	// TODO: update to github.com/AllenDang/w32 when
	// https://github.com/AllenDang/w32/pull/81 is merged.
	"github.com/TheTitanrain/w32"
)

// processNetworkPackets processes network incoming packets.
//
// ref: 0x440153
func processNetworkPackets() {
	// Handle incoming network packets from front-end.
	timer := time.NewTimer(time.Second / 100)
	for {
		select {
		case <-timer.C:
			//dbg.Println("timeout")
			return
		case action := <-proto.Actions:
			switch action := action.(type) {
			case proto.ButtonPressedAction:
				button := pixelgl.Button(action.Button)
				pos := image.Pt(int(action.X), int(action.Y))
				dbg.Printf("button pressed: %v at %v", button, pos)
				if isMouseButton(button) {
					// handle mouse button click.
					*diablo.MouseX = action.X
					*diablo.MouseY = action.Y
					w32.SetCursorPos(int(action.X), int(action.Y))
					switch button {
					case pixelgl.MouseButtonLeft:
						diablo.OnLeftMouseClick()
					case pixelgl.MouseButtonRight:
						diablo.OnRightMouseClick()
					}
				} else if keysym, ok := keyboardKey(button); ok {
					// handle key press.
					diablo.OnKeyPress(keysym)
					diablo.OnCharPress(keysym)
				}
			case proto.ButtonReleasedAction:
				button := pixelgl.Button(action.Button)
				pos := image.Pt(int(action.X), int(action.Y))
				dbg.Printf("button released: %v at %v", button, pos)
				if isMouseButton(button) {
					// handle mouse button click.
					*diablo.MouseX = action.X
					*diablo.MouseY = action.Y
					w32.SetCursorPos(int(action.X), int(action.Y))
					switch button {
					case pixelgl.MouseButtonLeft:
						diablo.OnLeftMouseRelease()
					}
				}
			default:
				panic(fmt.Errorf("support for action %T not yet implemented", action))
			}
		}
	}
}

// isMouseButton reports whether the given Pixel button is a mouse button.
func isMouseButton(button pixelgl.Button) bool {
	switch button {
	case pixelgl.MouseButton1, pixelgl.MouseButton2, pixelgl.MouseButton3,
		pixelgl.MouseButton4, pixelgl.MouseButton5, pixelgl.MouseButton6,
		pixelgl.MouseButton7, pixelgl.MouseButton8:
		return true
	}
	return false
}

// keyboardKey returns the corresponding keyboard key of the given Pixel button.
// The boolean return value indicate if the given button was a keyboard key.
func keyboardKey(button pixelgl.Button) (keysym int, ok bool) {
	switch button {
	case pixelgl.KeySpace:
		return w32.VK_SPACE, true
	case pixelgl.KeyApostrophe:
		return '\'', true
	case pixelgl.KeyComma:
		return ',', true
	case pixelgl.KeyMinus:
		return '-', true
	case pixelgl.KeyPeriod:
		return '.', true
	case pixelgl.KeySlash:
		return '/', true
	case pixelgl.Key0:
		return '0', true
	case pixelgl.Key1:
		return '1', true
	case pixelgl.Key2:
		return '2', true
	case pixelgl.Key3:
		return '3', true
	case pixelgl.Key4:
		return '4', true
	case pixelgl.Key5:
		return '5', true
	case pixelgl.Key6:
		return '6', true
	case pixelgl.Key7:
		return '7', true
	case pixelgl.Key8:
		return '8', true
	case pixelgl.Key9:
		return '9', true
	case pixelgl.KeySemicolon:
		return ';', true
	case pixelgl.KeyEqual:
		return '=', true
	case pixelgl.KeyA:
		return 'A', true
	case pixelgl.KeyB:
		return 'B', true
	case pixelgl.KeyC:
		return 'C', true
	case pixelgl.KeyD:
		return 'D', true
	case pixelgl.KeyE:
		return 'E', true
	case pixelgl.KeyF:
		return 'F', true
	case pixelgl.KeyG:
		return 'G', true
	case pixelgl.KeyH:
		return 'H', true
	case pixelgl.KeyI:
		return 'I', true
	case pixelgl.KeyJ:
		return 'J', true
	case pixelgl.KeyK:
		return 'K', true
	case pixelgl.KeyL:
		return 'L', true
	case pixelgl.KeyM:
		return 'M', true
	case pixelgl.KeyN:
		return 'N', true
	case pixelgl.KeyO:
		return 'O', true
	case pixelgl.KeyP:
		return 'P', true
	case pixelgl.KeyQ:
		return 'Q', true
	case pixelgl.KeyR:
		return 'R', true
	case pixelgl.KeyS:
		return 'S', true
	case pixelgl.KeyT:
		return 'T', true
	case pixelgl.KeyU:
		return 'U', true
	case pixelgl.KeyV:
		return 'V', true
	case pixelgl.KeyW:
		return 'W', true
	case pixelgl.KeyX:
		return 'X', true
	case pixelgl.KeyY:
		return 'Y', true
	case pixelgl.KeyZ:
		return 'Z', true
	case pixelgl.KeyLeftBracket:
		return '[', true
	case pixelgl.KeyBackslash:
		return '\\', true
	case pixelgl.KeyRightBracket:
		return ']', true
	case pixelgl.KeyGraveAccent:
		return '`', true
	case pixelgl.KeyWorld1:
		panic(fmt.Errorf("support for key %v not yet implemented", button))
	case pixelgl.KeyWorld2:
		panic(fmt.Errorf("support for key %v not yet implemented", button))
	case pixelgl.KeyEscape:
		return w32.VK_ESCAPE, true
	case pixelgl.KeyEnter:
		return w32.VK_RETURN, true
	case pixelgl.KeyTab:
		return w32.VK_TAB, true
	case pixelgl.KeyBackspace:
		return w32.VK_BACK, true
	case pixelgl.KeyInsert:
		return w32.VK_INSERT, true
	case pixelgl.KeyDelete:
		return w32.VK_DELETE, true
	case pixelgl.KeyRight:
		return w32.VK_RIGHT, true
	case pixelgl.KeyLeft:
		return w32.VK_LEFT, true
	case pixelgl.KeyDown:
		return w32.VK_DOWN, true
	case pixelgl.KeyUp:
		return w32.VK_UP, true
	case pixelgl.KeyPageUp:
		panic(fmt.Errorf("support for key %v not yet implemented", button))
	case pixelgl.KeyPageDown:
		panic(fmt.Errorf("support for key %v not yet implemented", button))
	case pixelgl.KeyHome:
		return w32.VK_HOME, true
	case pixelgl.KeyEnd:
		return w32.VK_END, true
	case pixelgl.KeyCapsLock:
		return w32.VK_CAPITAL, true
	case pixelgl.KeyScrollLock:
		return w32.VK_SCROLL, true
	case pixelgl.KeyNumLock:
		return w32.VK_NUMLOCK, true
	case pixelgl.KeyPrintScreen:
		return w32.VK_PRINT, true
	case pixelgl.KeyPause:
		return w32.VK_PAUSE, true
	case pixelgl.KeyF1:
		return w32.VK_F1, true
	case pixelgl.KeyF2:
		return w32.VK_F2, true
	case pixelgl.KeyF3:
		return w32.VK_F3, true
	case pixelgl.KeyF4:
		return w32.VK_F4, true
	case pixelgl.KeyF5:
		return w32.VK_F5, true
	case pixelgl.KeyF6:
		return w32.VK_F6, true
	case pixelgl.KeyF7:
		return w32.VK_F7, true
	case pixelgl.KeyF8:
		return w32.VK_F8, true
	case pixelgl.KeyF9:
		return w32.VK_F9, true
	case pixelgl.KeyF10:
		return w32.VK_F10, true
	case pixelgl.KeyF11:
		return w32.VK_F11, true
	case pixelgl.KeyF12:
		return w32.VK_F12, true
	case pixelgl.KeyF13:
		return w32.VK_F13, true
	case pixelgl.KeyF14:
		return w32.VK_F14, true
	case pixelgl.KeyF15:
		return w32.VK_F15, true
	case pixelgl.KeyF16:
		return w32.VK_F16, true
	case pixelgl.KeyF17:
		return w32.VK_F17, true
	case pixelgl.KeyF18:
		return w32.VK_F18, true
	case pixelgl.KeyF19:
		return w32.VK_F19, true
	case pixelgl.KeyF20:
		return w32.VK_F20, true
	case pixelgl.KeyF21:
		return w32.VK_F21, true
	case pixelgl.KeyF22:
		return w32.VK_F22, true
	case pixelgl.KeyF23:
		return w32.VK_F23, true
	case pixelgl.KeyF24:
		return w32.VK_F24, true
	case pixelgl.KeyF25:
		panic(fmt.Errorf("support for key %v not yet implemented", button))
	case pixelgl.KeyKP0:
		return w32.VK_NUMPAD0, true
	case pixelgl.KeyKP1:
		return w32.VK_NUMPAD1, true
	case pixelgl.KeyKP2:
		return w32.VK_NUMPAD2, true
	case pixelgl.KeyKP3:
		return w32.VK_NUMPAD3, true
	case pixelgl.KeyKP4:
		return w32.VK_NUMPAD4, true
	case pixelgl.KeyKP5:
		return w32.VK_NUMPAD5, true
	case pixelgl.KeyKP6:
		return w32.VK_NUMPAD6, true
	case pixelgl.KeyKP7:
		return w32.VK_NUMPAD7, true
	case pixelgl.KeyKP8:
		return w32.VK_NUMPAD8, true
	case pixelgl.KeyKP9:
		return w32.VK_NUMPAD9, true
	case pixelgl.KeyKPDecimal:
		return w32.VK_DECIMAL, true
	case pixelgl.KeyKPDivide:
		return w32.VK_DIVIDE, true
	case pixelgl.KeyKPMultiply:
		return w32.VK_MULTIPLY, true
	case pixelgl.KeyKPSubtract:
		return w32.VK_SUBTRACT, true
	case pixelgl.KeyKPAdd:
		return w32.VK_ADD, true
	case pixelgl.KeyKPEnter:
		return w32.VK_RETURN, true
	case pixelgl.KeyKPEqual:
		return '=', true
	case pixelgl.KeyLeftShift:
		return w32.VK_LSHIFT, true
	case pixelgl.KeyLeftControl:
		return w32.VK_LCONTROL, true
	case pixelgl.KeyLeftAlt:
		return w32.VK_LMENU, true
	case pixelgl.KeyLeftSuper:
		return w32.VK_LWIN, true
	case pixelgl.KeyRightShift:
		return w32.VK_RSHIFT, true
	case pixelgl.KeyRightControl:
		return w32.VK_RCONTROL, true
	case pixelgl.KeyRightAlt:
		return w32.VK_RMENU, true
	case pixelgl.KeyRightSuper:
		return w32.VK_RWIN, true
	case pixelgl.KeyMenu:
		return w32.VK_MENU, true
	default:
		panic(fmt.Errorf("support for key %v not yet implemented", button))
	}
}
