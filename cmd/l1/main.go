// The l1 tool provides dynamic random level generation of Cathedral maps.
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
		// Initial seed for dungeon generation.
		seed int32
		// Active quest ID.
		questID quests.QuestID
		// Store output in raw binary format.
		raw bool
	)
	var s, d int64
	var q string
	flag.Int64Var(&d, "dlvl", 1, "dungeon level (1-16)")
	flag.Int64Var(&s, "seed", 0, "initial seed")
	flag.BoolVar(&raw, "raw", false, "raw output format")
	flag.StringVar(&q, "quest", "", "active quest")
	flag.Usage = usage
	flag.Parse()
	dlvl = uint8(d)
	switch {
	case s >= -2147483648 || s <= 2147483647:
		seed = int32(s)
	default:
		panic(fmt.Errorf("invalid seed; expected >= -2147483648 and <= 2147483647, got %d", s))
	}
	if len(q) > 0 {
		if err := questID.Set(q); err != nil {
			log.Fatalf("unable to set quest ID; %v", err)
		}
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

	// Dump dungeon.
	if err := dump(*gendung.TileIDMap, raw); err != nil {
		log.Fatalf("%+v", err)
	}
}

// dump dumps the map to standard output.
func dump(m [40][40]uint8, raw bool) error {
	if raw {
		if err := binary.Write(os.Stdout, binary.LittleEndian, m); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
	buf, err := json.Marshal(m)
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
