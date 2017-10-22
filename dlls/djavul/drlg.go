package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"github.com/kr/pretty"
	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/diablo"
	"github.com/sanctuary/djavul/gendung"
	dinit "github.com/sanctuary/djavul/init"
	"github.com/sanctuary/djavul/l1"
)

// dumpL1Maps stores the tile ID, dungeon piece ID and arch num maps for the
// first 256 seeds of the procedurally generated Cathedral dungeon.
func dumpL1Maps() {
	dinit.Archives()
	*gendung.DLvl = 1
	*gendung.DType = gendung.Cathedral
	diablo.LoadLevelGraphics()
	for seed := int32(0); seed <= 0xFF; seed++ {
		l1.CreateDungeon(seed, 0)
		if err := dumpL1Map(seed); err != nil {
			log.Fatalf("+%v", err)
		}
	}
}

// dumpL1Map stores the tile ID, dungeon piece ID and arch num maps for the
// given seed of the procedurally generated Cathedral dungeon.
func dumpL1Map(seed int32) error {
	path := fmt.Sprintf("l1_tiles_%08X.bin", seed)
	if err := dumpData(path, *gendung.TileIDMap); err != nil {
		return errors.WithStack(err)
	}
	path = fmt.Sprintf("l1_pieces_%08X.bin", seed)
	if err := dumpData(path, *gendung.PieceIDMap); err != nil {
		return errors.WithStack(err)
	}
	path = fmt.Sprintf("l1_arches_%08X.bin", seed)
	if err := dumpData(path, *gendung.ArchNumMap); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// dumpData stores the data to the given path in little endian format.
func dumpData(path string, data interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	if err := binary.Write(f, binary.LittleEndian, data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// drlgCheck validates the implemenation of the dynamic random level generation
// of Cathedral maps.
func drlgCheck() error {
	dinit.Archives()
	*gendung.DLvl = 1
	*gendung.DType = gendung.Cathedral
	diablo.LoadLevelGraphics()
	seed := int32(123)
	l1.CreateDungeon(seed, 0)
	if err := check(*gendung.TileIDMap, "tiles", "12a0410904ebf2507b6b7017f0ae191ae476686b"); err != nil {
		path := fmt.Sprintf("l1_tiles_%08X.bin", seed)
		f, err := os.Open(path)
		if err != nil {
			return errors.WithStack(err)
		}
		defer f.Close()
		var tiles [40][40]uint8
		if err := binary.Read(f, binary.LittleEndian, &tiles); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.ArchNumMap)
		pretty.Println("want:", tiles)

		return errors.WithStack(err)
	}
	if err := check(*gendung.PieceIDMap, "pieces", "e15a7afb7505cb01b0b3d1befce5b8d4833ae1c6"); err != nil {
		path := fmt.Sprintf("l1_pieces_%08X.bin", seed)
		f, err := os.Open(path)
		if err != nil {
			return errors.WithStack(err)
		}
		defer f.Close()
		var pieces [112][112]int32
		if err := binary.Read(f, binary.LittleEndian, &pieces); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.ArchNumMap)
		pretty.Println("want:", pieces)

		return errors.WithStack(err)
	}
	if err := check(*gendung.ArchNumMap, "arches", "5438e3d7761025a2ee6f7fec155c840fc289f5dd"); err != nil {
		path := fmt.Sprintf("l1_arches_%08X.bin", seed)
		f, err := os.Open(path)
		if err != nil {
			return errors.WithStack(err)
		}
		defer f.Close()
		var arches [112][112]int8
		if err := binary.Read(f, binary.LittleEndian, &arches); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.ArchNumMap)
		pretty.Println("want:", arches)

		return errors.WithStack(err)
	}
	fmt.Println("PASS")
	return nil
}

// check validates the data against the given SHA1 hashsum.
func check(data interface{}, name, want string) error {
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		return errors.WithStack(err)
	}
	sum := sha1.Sum(buf.Bytes())
	got := fmt.Sprintf("%040x", sum[:])
	if got != want {
		return errors.Errorf("SHA1 hash mismatch for %v, seed 0; expected %q, got %q", name, want, got)
	}
	return nil
}