package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"path/filepath"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mewkiz/pkg/imgutil"
	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/internal/proto"
	"github.com/sanctuary/formats/image/cel"
	"github.com/sanctuary/formats/image/cel/config"
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

func listenTCP() {
	l, err := net.Listen("tcp", "127.0.0.1:6667")
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	fmt.Println("Listening on TCP")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("unable to accept connection; %v", err)
			continue
		}
		dec := gob.NewDecoder(conn)
		for {
			var cmd proto.CommandTCP
			if err := dec.Decode(&cmd); err != nil {
				if errors.Cause(err) == io.EOF {
					fmt.Println("disconnected")
					break
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
}

func listenUDP(win *pixelgl.Window) {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6666")
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}
	fmt.Println("Listening on UDP")
	dec := gob.NewDecoder(conn)
	for {
		var pkg proto.PacketUDP
		if err := dec.Decode(&pkg); err != nil {
			if errors.Cause(err) == io.EOF {
				fmt.Println("disconnected")
				dec = gob.NewDecoder(conn)
				continue
			}
			log.Printf("unable to decode UDP packet; %+v", errors.WithStack(err))
			continue
		}
		fmt.Println("recv cmd:", pkg.Cmd)
		switch pkg.Cmd {
		case proto.CmdDrawImage:
			ExecDrawImage(win, pkg.Data)
		case proto.CmdUpdateScreen:
			fmt.Println("recv cmd: UpdateScreen")
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
	go listenUDP(win)
	listenTCP()
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

func ExecDrawImage(win *pixelgl.Window, data []byte) {
	var cmd proto.DrawImage
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&cmd); err != nil {
		log.Printf("unable to parse body of LoadFile; %v", errors.WithStack(err))
		return
	}
	fmt.Println("recv pkg:", cmd)
	pics := getPictures(cmd.Path)
	pic := pics[cmd.FrameNum]
	sprite := pixel.NewSprite(pic, pic.Bounds())
	const screenHeight = 480
	sprite.Draw(win, pixel.IM.Moved(pic.Bounds().Center().Add(pixel.V(cmd.X, cmd.Y))))
}

// ### [ Helper functions ] ####################################################

// pictures maps from relative file path to decoded image frames.
var pictures = make(map[string][]pixel.Picture)

// getPictures returns the pictures associated with the given file path.
func getPictures(relPath string) []pixel.Picture {
	pics, ok := pictures[relPath]
	if !ok {
		panic(fmt.Errorf("unable to locate decoded image frames of %q", relPath))
	}
	return pics
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

// loadPics loads the frames of the given CEL image.
func loadPics(relPath string) error {
	name := filepath.Base(relPath)
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
	var pics []pixel.Picture
	for _, img := range imgs {
		pic := pixel.PictureDataFromImage(img)
		pics = append(pics, pic)
	}
	pictures[relPath] = pics
	return nil
}
