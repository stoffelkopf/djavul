//+build ddraw

// Package capture implements screenshot capturing.
package capture

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewkiz/pkg/osutil"
	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/dx"
	"github.com/sanctuary/djavul/internal/ddraw"
	"github.com/sanctuary/djavul/scrollrt"
)

// Screenshot captures a screenshot and stores it within the game directory as
// "screenXX.png". While the screenshot is being taken, the game is paused
// temporarily as indicated by the red screen.
//
// ref: 0x40311B
func Screenshot() {
	fmt.Println("capture.Screenshot")
	const (
		width  = 640
		height = 480
	)
	// Re-draw.
	scrollrt.DrawMainW()
	// Get active palette.
	entries := make([]ddraw.PALETTEENTRY, 256)
	if err := (*dx.DDP).GetEntries(0, entries); err != nil {
		log.Fatalf("%+v", err)
	}
	// Set active palette to red-scale.
	SetRedPalette(entries)
	// Store screenshot.
	pal := dx.PalFromEntries(entries)
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	dx.LockMutex()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := pal[(*dx.ScreenBuf).Row[y].Pixels[x]]
			dst.Set(x, y, c)
		}
	}
	dx.UnlockMutex()
	path, err := GenFileName()
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	log.Printf("creating `%v`.", path)
	if err := imgutil.WriteFile(path, dst); err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	// Pause temporarily with red screen.
	time.Sleep(300 * time.Millisecond)
	// Restore active palette.
	if err := (*dx.DDP).SetEntries(0, entries); err != nil {
		log.Fatalf("%+v", err)
	}
}

// GenFileName returns an unused file name within the game directory, (e.g.
// "screenXX.png").
//
// ref: 0x4033A8
func GenFileName() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", errors.WithStack(err)
	}
	gameDir := filepath.Dir(exePath)
	for i := 0; ; i++ {
		name := fmt.Sprintf("screen%02d.png", i)
		path := filepath.Join(gameDir, name)
		if !osutil.Exists(path) {
			return path, nil
		}
	}
}

// SetRedPalette sets the system palette to red-scale.
//
// ref: 0x403470
func SetRedPalette(entries []ddraw.PALETTEENTRY) {
	red := make([]ddraw.PALETTEENTRY, len(entries))
	for i, e := range entries {
		red[i].Red = e.Red
		red[i].Green = 0
		red[i].Blue = 0
		red[i].Flags = 0
	}
	if err := (*dx.DDP).SetEntries(0, red); err != nil {
		log.Fatalf("%+v", err)
	}
}
