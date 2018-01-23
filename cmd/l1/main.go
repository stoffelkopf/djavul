// The l1 tool provides dynamic random level generation of Cathedral maps.
//
// Usage:
//
//    l1 [OPTION]...
//
// Flags:
//
//    -dpieces
//          output dungeon pieces in CSV format
//    -dlvl int
//          dungeon level (1-16) (default 1)
//    -quest string
//          active quest
//    -raw
//          raw output format
//    -seed int
//          initial seed
package main

import (
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/diablo"
	"github.com/sanctuary/djavul/gendung"
	"github.com/sanctuary/djavul/l1"
	"github.com/sanctuary/djavul/multi"
	"github.com/sanctuary/djavul/quests"
	"github.com/sanctuary/formats/level/til"
)

func usage() {
	const use = `
Dynamic random level generation of Cathedral maps.

Usage:

	l1 [OPTION]...

Flags:
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line flags.
	var (
		// Dungeon level.
		dlvl uint8
		// Output dungeon pieces instead of tiles.
		outputDPieces bool
		// Active quest ID.
		questID quests.QuestID
		// Store output in raw binary format.
		raw bool
		// Initial seed for dungeon generation.
		seed int32
	)
	var d, s int64
	var q string
	flag.Int64Var(&d, "dlvl", 1, "dungeon level (1-16)")
	flag.BoolVar(&outputDPieces, "dpieces", false, "output dungeon pieces in CSV format")
	flag.StringVar(&q, "quest", "", "active quest")
	flag.BoolVar(&raw, "raw", false, "raw output format")
	flag.Int64Var(&s, "seed", 0, "initial seed")
	flag.Usage = usage
	flag.Parse()
	dlvl = uint8(d)
	if len(q) > 0 {
		if err := questID.Set(q); err != nil {
			log.Fatalf("unable to set quest ID; %v", err)
		}
	}
	switch {
	case s >= -2147483648 && s <= 2147483647:
		seed = int32(s)
	default:
		panic(fmt.Errorf("invalid seed; expected >= -2147483648 and <= 2147483647, got %d", s))
	}

	// Load level graphics.
	*gendung.DType = dtypeFromDLvl(dlvl)
	diablo.LoadLevelGraphics()

	// Generate dungeon.
	*gendung.DLvl = dlvl
	*multi.MaxPlayers = 1
	for i := range quests.Quests {
		quests.Quests[i].Active = false
	}
	*gendung.IsQuestLevel = false
	if questID != quests.Invalid {
		quests.Quests[questID].Active = true
		quests.Quests[questID].DLvl = dlvl
	}
	l1.CreateDungeon(seed, 0)

	// Dump dungeon pieces.
	tiles := *gendung.TileIDMap
	if outputDPieces {
		if err := dumpDungeonPieces(tiles); err != nil {
			log.Fatalf("%+v", err)
		}
		return
	}
	// Dump tiles.
	if err := dumpTiles(tiles, raw); err != nil {
		log.Fatalf("%+v", err)
	}
}

// dumpDungeonPieces dumps the dungeon pieces of the map to standard output.
func dumpDungeonPieces(tiles [40][40]uint8) error {
	dpieces, err := getDPieces(tiles)
	if err != nil {
		return errors.WithStack(err)
	}
	var vals []uint16
	for x := 0; x < 80; x++ {
		for y := 0; y < 80; y++ {
			vals = append(vals, dpieces[x][y])
		}
	}
	for i, v := range vals {
		if i != 0 {
			fmt.Print(",")
		}
		fmt.Print(v)
	}
	fmt.Println()
	return nil
}

// getDPieces converts the given tile map to a corresponding map of dungeon
// pieces.
func getDPieces(tiles [40][40]uint8) ([80][80]uint16, error) {
	tileDefs, err := til.Parse("diabdat/levels/l1data/l1.til")
	if err != nil {
		return [80][80]uint16{}, errors.WithStack(err)
	}
	var dpieces [80][80]uint16
	for x := 0; x < 40; x++ {
		for y := 0; y < 40; y++ {
			tile := tiles[y][x] // TODO: figure out why x and y are swapped.
			if tile != 0 {
				tileDef := tileDefs[tile-1]
				dpieces[2*x][2*y] = tileDef.Top + 1
				dpieces[2*x][2*y+1] = tileDef.Right + 1
				dpieces[2*x+1][2*y] = tileDef.Left + 1
				dpieces[2*x+1][2*y+1] = tileDef.Bottom + 1
			}
		}
	}
	return dpieces, nil
}

// dumpTiles dumps the tiles of the map to standard output.
func dumpTiles(tiles [40][40]uint8, raw bool) error {
	if raw {
		if err := binary.Write(os.Stdout, binary.LittleEndian, tiles); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
	buf, err := json.Marshal(tiles)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Println(string(buf))
	return nil
}

// dtypeFromDLvl returns the dungeon type of the given dungeon level.
func dtypeFromDLvl(dlvl uint8) gendung.DungeonType {
	switch dlvl {
	case 0:
		return gendung.Tristram
	case 1, 2, 3, 4:
		return gendung.Cathedral
	case 5, 6, 7, 8:
		return gendung.Catacombs
	case 9, 10, 11, 12:
		return gendung.Caves
	case 13, 14, 15, 16:
		return gendung.Hell
	}
	panic(fmt.Errorf("invalid dlvl; expected <= 16; got %d", dlvl))
}
