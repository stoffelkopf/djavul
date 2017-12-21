package main

import (
	"fmt"
	"image"
	"path/filepath"

	"github.com/faiface/pixel"
	"github.com/mewkiz/pkg/imgutil"
	"github.com/pkg/errors"
	"github.com/sanctuary/formats/image/cel"
	"github.com/sanctuary/formats/image/cel/config"
)

// loadIcon loads and returns the window icon.
func loadIcon() []pixel.Picture {
	path := "icon.png"
	icon, err := loadImage(path)
	if err != nil {
		warn.Printf("unable to load icon %q; %v", path, err)
		return nil
	}
	return []pixel.Picture{icon}
}

// sprites maps from relative file path to sprites, one per frame.
var sprites = make(map[string][]*pixel.Sprite)

// getSprite returns a sprite of the picture associated with the given file path
// and frame number.
func getSprite(relPath string, frameNum int) *pixel.Sprite {
	name := filepath.Base(relPath)
	switch name {
	case "l1.cel":
		const (
			frameWidth  = 32
			frameHeight = 32
		)
		pic := l1Sprite.Picture()
		picBounds := pic.Bounds()
		frame := l1Sprite.Frame()
		k := int(picBounds.W() / frameWidth)
		x := float64(frameWidth * (frameNum % k))
		y := picBounds.H() - frameHeight - float64(frameHeight*(frameNum/k))
		frame = pixel.R(x, y, x+frameWidth, y+frameHeight)
		l1Sprite.Set(pic, frame)
		return l1Sprite
	}
	ss, ok := sprites[relPath]
	if !ok {
		panic(fmt.Errorf("unable to locate decoded image frames of %q", relPath))
	}
	sprite := ss[frameNum]
	return sprite
}

// dirPictures maps from relative file path to decoded image frames based on
// direction.
var dirPictures = make(map[string][][]pixel.Picture)

// getPicturesForDir returns the pictures associated with the given file path
// and direction.
func getPicturesForDir(relPath string, direction int) []pixel.Picture {
	dirPics, ok := dirPictures[relPath]
	if !ok {
		panic(fmt.Errorf("unable to locate decoded image frames of %q", relPath))
	}
	if direction == 8 {
		direction = 0
	}
	if len(dirPics) <= direction {
		panic(fmt.Errorf("invalid direction for %q; expected < %d, got %d", relPath, len(dirPics), direction))
	}
	return dirPics[direction]
}

// absPath returns the absolute path to the given file, relative to the MPQ
// directory.
func absPath(relPath string) string {
	// mpqDir specifies a directory containing an extracted copy of the files
	// contained within DIABDAT.MPQ. Note that the extracted files should have
	// lowercase names.
	const mpqDir = "diabdat"
	return filepath.Join(mpqDir, relPath)
}

// l1.cel sprite
var l1Sprite *pixel.Sprite

// loadPics loads the frames of the given CEL image.
func loadPics(relPath string) error {
	name := filepath.Base(relPath)
	switch name {
	case "l1.cel":
		img, err := imgutil.ReadFile("l1.png")
		if err != nil {
			return errors.WithStack(err)
		}
		pic := pixel.PictureDataFromImage(img)
		const (
			frameWidth  = 32
			frameHeight = 32
		)
		frame := pixel.R(0, 0, frameWidth, frameHeight)
		l1Sprite = pixel.NewSprite(pic, frame)
		return nil
	}
	conf, err := config.Get(name)
	if err != nil {
		return errors.Errorf("unable to locate image config for %q; %v", name, err)
	}
	fmt.Println("decoding CEL image:", relPath)
	palPath := "levels/towndata/town.pal"
	if len(conf.Pals) > 0 {
		// TODO: Figure out how to handle multiple palettes.
		palPath = conf.Pals[0]
	}
	pal, err := cel.ParsePal(absPath(palPath))
	if err != nil {
		return errors.Errorf("unable to parse palette %q; %v", palPath, err)
	}
	path := absPath(relPath)

	// CEL archive.
	if conf.Nimgs != 0 {
		archiveImgs, err := cel.DecodeArchive(path, pal)
		if err != nil {
			return errors.Errorf("unable to parse CEL archive %q; %v", relPath, err)
		}
		var dirPics [][]pixel.Picture
		for _, archiveImg := range archiveImgs {
			var pics []pixel.Picture
			for _, img := range archiveImg {
				pic := pixel.PictureDataFromImage(img)
				pics = append(pics, pic)
			}
			dirPics = append(dirPics, pics)
		}
		dirPictures[relPath] = dirPics
		return nil
	}

	// CEL image.
	imgs, err := cel.DecodeAll(path, pal)
	if err != nil {
		return errors.Errorf("unable to decode CEL image %q; %v", relPath, err)
	}
	var ss []*pixel.Sprite
	for _, img := range imgs {
		pic := pixel.PictureDataFromImage(img)
		sprite := pixel.NewSprite(pic, pic.Bounds())
		ss = append(ss, sprite)
	}
	sprites[relPath] = ss
	return nil
}

// loadImage loads the given image file.
func loadImage(path string) (pixel.Picture, error) {
	img, err := imgutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return pixel.PictureDataFromImage(img), nil
}

// pixelRect returns the Pixel rectangle corresponding to the given Go image
// rectangle.
func pixelRect(picBounds pixel.Rect, r image.Rectangle) pixel.Rect {
	frameWidth := float64(r.Dx())
	frameHeight := float64(r.Dy())
	x := float64(r.Min.X)
	y := picBounds.H() - frameHeight - float64(r.Min.Y)
	return pixel.R(x, y, x+frameWidth, y+frameHeight)
}

// pixelVec returns the Pixel vector corresponding to the given Go image point.
func pixelVec(pt image.Point) pixel.Vec {
	return pixel.V(float64(pt.X), float64(pt.Y))
}

// point returns the Go image point corresponding to the given Pixel vector.
func point(v pixel.Vec) image.Point {
	return image.Pt(int(v.X), int(v.Y))
}
