package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"log"
	"os"

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
	for seed := uint32(0); seed <= 0xFF; seed++ {
		l1.CreateDungeon(seed, 0)
		if err := dumpL1Map(seed); err != nil {
			log.Fatalf("+%v", err)
		}
	}
}

// dumpL1Map stores the tile ID, dungeon piece ID and arch num maps for the
// given seed of the procedurally generated Cathedral dungeon.
func dumpL1Map(seed uint32) error {
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
	seed := uint32(0)
	l1.CreateDungeon(seed, 0)
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.LittleEndian, *gendung.TileIDMap); err != nil {
		return errors.WithStack(err)
	}
	sum := sha1.Sum(buf.Bytes())
	got := fmt.Sprintf("%040x", sum[:])
	want := "ff9be30983ee653ee2b3f33030091b8db3bd5a70"
	if got != want {
		return errors.Errorf("SHA1 hash mismatch for seed 0; expected %q, got %q", want, got)
	}
	fmt.Println("PASS")
	return nil
}
