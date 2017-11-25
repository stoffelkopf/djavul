package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mewkiz/pkg/imgutil"
	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/internal/proto"
	"github.com/sanctuary/formats/image/cel"
	"github.com/sanctuary/formats/image/cel/config"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	if err := front(); err != nil {
		log.Fatalf("%+v", err)
	}
}

// loadImage loads the given image file.
func loadImage(path string) (pixel.Picture, error) {
	img, err := imgutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return pixel.PictureDataFromImage(img), nil
}

func listenTCP(tmpDir string) {
	// Open pipe for writing.
	//wpath := filepath.Join(tmpDir, "tcp_w")
	//if err := syscall.Mkfifo(wpath, 0666); err != nil {
	//	log.Fatalf("%+v", errors.WithStack(err))
	//}
	//w, err := os.OpenFile(wpath, os.O_WRONLY, 0666)
	//if err != nil {
	//	log.Fatalf("%+v", errors.WithStack(err))
	//}
	//defer w.Close()
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
				fmt.Println("disconnected")
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
			fmt.Println("recv pkg:", data)
			ExecLoadFile(&data)
		}
	}
}

func listenUDP(win *pixelgl.Window, tmpDir string) {
	// Open pipe for writing.
	//wpath := filepath.Join(tmpDir, "udp_w")
	//if err := syscall.Mkfifo(wpath, 0666); err != nil {
	//	log.Fatalf("%+v", errors.WithStack(err))
	//}
	//w, err := os.OpenFile(wpath, os.O_WRONLY, 0666)
	//if err != nil {
	//	log.Fatalf("%+v", errors.WithStack(err))
	//}
	//defer w.Close()
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
				fmt.Println("disconnected")
				dec = gob.NewDecoder(r)
				continue
			}
			log.Printf("unable to decode UDP packet; %+v", errors.WithStack(err))
			continue
		}
		fmt.Println("recv cmd:", pkg.Cmd)
		switch pkg.Cmd {
		case proto.CmdUpdateScreen:
			fmt.Println("recv cmd: UpdateScreen")
			win.Clear(colornames.Black)
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

func front() error {
	var icons []pixel.Picture
	icon, err := loadImage("icon.png")
	if err != nil {
		log.Printf("unable to load icon %q; %v", "icon.png", err)
	} else {
		icons = append(icons, icon)
	}
	cfg := pixelgl.WindowConfig{
		Title:  "d1-frontend",
		Icon:   icons,
		Bounds: pixel.R(0, 0, 640, 480),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return errors.WithStack(err)
	}
	// Init temporary directory.
	const tmpDir = "/tmp/djavul"
	if err := os.RemoveAll(tmpDir); err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	// Listen on TCP and UDP.
	go listenUDP(win, tmpDir)
	listenTCP(tmpDir)
	return nil
}

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
	fmt.Println("recv pkg:", cmd)
	sprite := getSprite(cmd.Path, cmd.FrameNum)
	const screenHeight = 480
	frame := sprite.Frame()
	bounds := pixel.R(0, 0, frame.W(), frame.H())
	sprite.Draw(win, pixel.IM.Moved(bounds.Center().Add(pixel.V(cmd.X, cmd.Y))))
}

// ### [ Helper functions ] ####################################################

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
		fmt.Println("k:", k)
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
