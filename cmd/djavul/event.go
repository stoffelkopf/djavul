package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/internal/proto"
)

// loop initiates the event loop of the front-end.
func loop(win *pixelgl.Window) {
	// Relay window events.
	winEvents := make(chan WindowEvent)
	go relayWindowEvents(win, winEvents)
	// Relay events from and actions to the Diablo 1 game engine.
	gameEvents := make(chan proto.EngineEvent)
	gameActions := make(chan proto.EngineAction)
	go relayEngineEvents(win, gameEvents, gameActions)
	for {
		select {
		// Handle window events.
		case e := <-winEvents:
			switch e := e.(type) {
			case WindowClosedEvent:
				dbg.Printf("window closed")
				return
			case ButtonPressedEvent:
				dbg.Printf("button pressed: %v at %v", e.Button, e.Pos)
				gameActions <- buttonPressed(win, e)
			case ButtonReleasedEvent:
				dbg.Printf("button released: %v at %v", e.Button, e.Pos)
				gameActions <- buttonReleased(win, e)
			}
		// Handle engine event.
		//case e := <-gameEvents:
		default:
			// Poll window events.
			time.Sleep(time.Second / 100)
			win.UpdateInput()
		}
	}
}

// === [ Window events ] =======================================================

// relayWindowEvents relays window events from the window.
func relayWindowEvents(win *pixelgl.Window, winEvents chan WindowEvent) {
	// pixelButtons is a list of all mouse buttons and keyboard keys in PixelGL.
	var pixelButtons = []pixelgl.Button{pixelgl.MouseButton1, pixelgl.MouseButton2, pixelgl.MouseButton3, pixelgl.MouseButton4, pixelgl.MouseButton5, pixelgl.MouseButton6, pixelgl.MouseButton7, pixelgl.MouseButton8, pixelgl.MouseButtonLast, pixelgl.MouseButtonLeft, pixelgl.MouseButtonRight, pixelgl.MouseButtonMiddle, pixelgl.KeySpace, pixelgl.KeyApostrophe, pixelgl.KeyComma, pixelgl.KeyMinus, pixelgl.KeyPeriod, pixelgl.KeySlash, pixelgl.Key0, pixelgl.Key1, pixelgl.Key2, pixelgl.Key3, pixelgl.Key4, pixelgl.Key5, pixelgl.Key6, pixelgl.Key7, pixelgl.Key8, pixelgl.Key9, pixelgl.KeySemicolon, pixelgl.KeyEqual, pixelgl.KeyA, pixelgl.KeyB, pixelgl.KeyC, pixelgl.KeyD, pixelgl.KeyE, pixelgl.KeyF, pixelgl.KeyG, pixelgl.KeyH, pixelgl.KeyI, pixelgl.KeyJ, pixelgl.KeyK, pixelgl.KeyL, pixelgl.KeyM, pixelgl.KeyN, pixelgl.KeyO, pixelgl.KeyP, pixelgl.KeyQ, pixelgl.KeyR, pixelgl.KeyS, pixelgl.KeyT, pixelgl.KeyU, pixelgl.KeyV, pixelgl.KeyW, pixelgl.KeyX, pixelgl.KeyY, pixelgl.KeyZ, pixelgl.KeyLeftBracket, pixelgl.KeyBackslash, pixelgl.KeyRightBracket, pixelgl.KeyGraveAccent, pixelgl.KeyWorld1, pixelgl.KeyWorld2, pixelgl.KeyEscape, pixelgl.KeyEnter, pixelgl.KeyTab, pixelgl.KeyBackspace, pixelgl.KeyInsert, pixelgl.KeyDelete, pixelgl.KeyRight, pixelgl.KeyLeft, pixelgl.KeyDown, pixelgl.KeyUp, pixelgl.KeyPageUp, pixelgl.KeyPageDown, pixelgl.KeyHome, pixelgl.KeyEnd, pixelgl.KeyCapsLock, pixelgl.KeyScrollLock, pixelgl.KeyNumLock, pixelgl.KeyPrintScreen, pixelgl.KeyPause, pixelgl.KeyF1, pixelgl.KeyF2, pixelgl.KeyF3, pixelgl.KeyF4, pixelgl.KeyF5, pixelgl.KeyF6, pixelgl.KeyF7, pixelgl.KeyF8, pixelgl.KeyF9, pixelgl.KeyF10, pixelgl.KeyF11, pixelgl.KeyF12, pixelgl.KeyF13, pixelgl.KeyF14, pixelgl.KeyF15, pixelgl.KeyF16, pixelgl.KeyF17, pixelgl.KeyF18, pixelgl.KeyF19, pixelgl.KeyF20, pixelgl.KeyF21, pixelgl.KeyF22, pixelgl.KeyF23, pixelgl.KeyF24, pixelgl.KeyF25, pixelgl.KeyKP0, pixelgl.KeyKP1, pixelgl.KeyKP2, pixelgl.KeyKP3, pixelgl.KeyKP4, pixelgl.KeyKP5, pixelgl.KeyKP6, pixelgl.KeyKP7, pixelgl.KeyKP8, pixelgl.KeyKP9, pixelgl.KeyKPDecimal, pixelgl.KeyKPDivide, pixelgl.KeyKPMultiply, pixelgl.KeyKPSubtract, pixelgl.KeyKPAdd, pixelgl.KeyKPEnter, pixelgl.KeyKPEqual, pixelgl.KeyLeftShift, pixelgl.KeyLeftControl, pixelgl.KeyLeftAlt, pixelgl.KeyLeftSuper, pixelgl.KeyRightShift, pixelgl.KeyRightControl, pixelgl.KeyRightAlt, pixelgl.KeyRightSuper, pixelgl.KeyMenu, pixelgl.KeyLast}
	for {
		if win.Closed() {
			winEvents <- WindowClosedEvent{}
		}
		for _, button := range pixelButtons {
			if win.JustPressed(button) {
				vec := win.MousePosition()
				bounds := win.Bounds()
				vec.Y = bounds.H() - vec.Y
				pos := point(vec)
				winEvents <- ButtonPressedEvent{
					Button: button,
					Pos:    pos,
				}
			}
			if win.JustReleased(button) {
				vec := win.MousePosition()
				bounds := win.Bounds()
				vec.Y = bounds.H() - vec.Y
				pos := point(vec)
				winEvents <- ButtonReleasedEvent{
					Button: button,
					Pos:    pos,
				}
			}
		}
	}
}

// WindowEvent is a window event.
type WindowEvent interface {
	// isWindowEvent ensures that only window events can be assigned to the
	// WindowEvent interface.
	isWindowEvent()
}

// WindowClosedEvent signals that the window has been closed.
type WindowClosedEvent struct{}

// ButtonPressedEvent signals that a mouse button or keyboard key has been
// pressed.
type ButtonPressedEvent struct {
	// Mouse button or keyboard key pressed.
	Button pixelgl.Button
	// Mouse position on screen.
	Pos image.Point
}

// ButtonReleasedEvent signals that a mouse button or keyboard key has been
// released.
type ButtonReleasedEvent struct {
	// Mouse button or keyboard key released.
	Button pixelgl.Button
	// Mouse position on screen.
	Pos image.Point
}

// isWindowEvent ensures that only window events can be assigned to the
// WindowEvent interface.
func (WindowClosedEvent) isWindowEvent()   {}
func (ButtonPressedEvent) isWindowEvent()  {}
func (ButtonReleasedEvent) isWindowEvent() {}

// === [ Engine events ] =======================================================

// relayEngineEvents relays events to and from the Diablo 1 game engine.
func relayEngineEvents(win *pixelgl.Window, gameEvents chan proto.EngineEvent, gameActions chan proto.EngineAction) {
	tmpDir := initTmpDir()
	// Relay events on unstable connection to the Diablo 1 game engine.
	go relayEngineUnstableEvents(tmpDir, win, gameEvents, gameActions)
	go relayEngineUnstableActions(tmpDir, win, gameActions)
	// Relay events on stable connection to the Diablo 1 game engine.
	//go relayEngineStableActions(tmpDir, win, gameActions)
	relayEngineStableEvents(tmpDir, win, gameEvents, gameActions)
}

// relayEngineStableActions relays actions on stable connection to the Diablo 1
// game engine.
func relayEngineStableActions(tmpDir string, win *pixelgl.Window, gameActions chan proto.EngineAction) {
	// Open pipe for writing.
	wpath := filepath.Join(tmpDir, "tcp_w")
	if err := syscall.Mkfifo(wpath, 0666); err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	w, err := os.OpenFile(wpath, os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	defer w.Close()
	fmt.Printf("Writing to %q.\n", wpath)
	enc := gob.NewEncoder(w)
	_ = enc
	for {
		action := <-gameActions
		switch action := action.(type) {
		default:
			panic(fmt.Errorf("support for action %T not yet implemented", action))
		}
	}
}

// relayEngineUnstableActions relays actions on unstable connection to the
// Diablo 1 game engine.
func relayEngineUnstableActions(tmpDir string, win *pixelgl.Window, gameActions chan proto.EngineAction) {
	// Open pipe for writing.
	wpath := filepath.Join(tmpDir, "udp_w")
	if err := syscall.Mkfifo(wpath, 0666); err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	w, err := os.OpenFile(wpath, os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	defer w.Close()
	fmt.Printf("Writing to %q.\n", wpath)
	enc := gob.NewEncoder(w)
	for {
		action := <-gameActions
		pkt := proto.NewAction(action)
		if err := enc.Encode(pkt); err != nil {
			die(err)
		}
	}
}

// relayEngineStableEvents relays events on stable connection to the Diablo 1
// game engine.
func relayEngineStableEvents(tmpDir string, win *pixelgl.Window, gameEvents chan proto.EngineEvent, gameActions chan proto.EngineAction) {
	// Open pipe for reading.
	rpath := filepath.Join(tmpDir, "tcp_r")
	if err := syscall.Mkfifo(rpath, 0666); err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	r, err := os.OpenFile(rpath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	defer r.Close()
	fmt.Printf("Listening on %q.\n", rpath)
	dec := gob.NewDecoder(r)
	for {
		var cmd proto.CommandTCP
		if err := dec.Decode(&cmd); err != nil {
			if errors.Cause(err) == io.EOF {
				//fmt.Println("disconnected")
				dec = gob.NewDecoder(r)
				continue
			}
			log.Fatalf("%+v", errors.WithStack(err))
		}
		switch cmd {
		case proto.CmdLoadFile:
			var data proto.LoadFile
			if err := dec.Decode(&data); err != nil {
				log.Fatalf("%+v", errors.WithStack(err))
			}
			//fmt.Println("recv pkg:", data)
			ExecLoadFile(&data)
		}
	}
}

// relayEngineUnstableEvents relays events on unstable connection to the Diablo
// 1 game engine.
func relayEngineUnstableEvents(tmpDir string, win *pixelgl.Window, gameEvents chan proto.EngineEvent, gameActions chan proto.EngineAction) {
	// Open pipe for reading.
	rpath := filepath.Join(tmpDir, "udp_r")
	if err := syscall.Mkfifo(rpath, 0666); err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	r, err := os.OpenFile(rpath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	defer r.Close()
	fmt.Printf("Listening on %q.\n", rpath)
	dec := gob.NewDecoder(r)
	frames := 0
	var start time.Time
	for {
		var pkg proto.PacketUDP
		if err := dec.Decode(&pkg); err != nil {
			if errors.Cause(err) == io.EOF {
				//fmt.Println("disconnected")
				dec = gob.NewDecoder(r)
				continue
			}
			log.Printf("unable to decode UDP packet; %+v", errors.WithStack(err))
			continue
		}
		//fmt.Println("recv cmd:", pkg.Cmd)
		switch pkg.Cmd {
		case proto.CmdUpdateScreen:
			//fmt.Println("recv cmd: UpdateScreen")
			//win.Clear(colornames.Black)
			ExecDrawImages(win, pkg.Data)
			if start == (time.Time{}) {
				start = time.Now()
			} else {
				frames++
				fps := float64(frames) / (float64(time.Since(start)) / float64(time.Second))
				win.SetTitle(fmt.Sprintf("FPS: %.02f", fps))
			}
			win.Update()
		}
	}
}

// ### [ Old, remove code below ] ##############################################

func ExecLoadFile(cmd *proto.LoadFile) {
	switch filepath.Ext(cmd.Path) {
	case ".cel":
		if err := loadPics(cmd.Path); err != nil {
			log.Fatalf("%+v", err)
		}
	}
}

func ExecDrawImages(win *pixelgl.Window, data []byte) {
	var cmds []proto.DrawImage
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&cmds); err != nil {
		log.Printf("unable to parse body of DrawImages; %v", errors.WithStack(err))
		return
	}
	for _, cmd := range cmds {
		ExecDrawImage(win, cmd)
	}
}

func ExecDrawImage(win *pixelgl.Window, cmd proto.DrawImage) {
	//fmt.Println("recv pkg:", cmd)
	sprite := getSprite(cmd.Path, cmd.FrameNum)
	const screenHeight = 480
	dp := cmd.Dp
	dp.Y = screenHeight - 1 - dp.Y
	frame := sprite.Frame()
	if cmd.Sr != image.ZR {
		pic := sprite.Picture()
		picBounds := pic.Bounds()
		frame = pixelRect(picBounds, cmd.Sr)
		sprite.Set(pic, frame)
		dp.Y -= int(frame.H())
	}
	bounds := pixel.R(0, 0, frame.W(), frame.H())
	sprite.Draw(win, pixel.IM.Moved(bounds.Center().Add(pixelVec(dp))))
}

// ### [ Helper functions ] ####################################################

// initTmpDir initializes the temporary directory used by Djavul to store Unix
// sockets for communication with the Diablo 1 game engine.
func initTmpDir() string {
	const tmpDir = "/tmp/djavul"
	if err := os.RemoveAll(tmpDir); err != nil {
		die(err)
	}
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		die(err)
	}
	return tmpDir
}
