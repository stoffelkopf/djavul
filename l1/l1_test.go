//+build !djavul

package l1_test

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/diablo"
	"github.com/sanctuary/djavul/gendung"
	dinit "github.com/sanctuary/djavul/init"
	"github.com/sanctuary/djavul/l1"
	"github.com/sanctuary/djavul/multi"
	"github.com/sanctuary/djavul/quests"
)

func TestCreateDungeon(t *testing.T) {
	// Load level graphics.
	//
	//    - levels/l1data/l1.cel
	//    - levels/l1data/l1s.cel
	//    - levels/l1data/l1.til
	//    - levels/l1data/l1.min
	dinit.Archives()
	golden := []struct {
		// meta.
		dungeonName string
		// pre.
		dlvl    uint8
		dtype   gendung.DungeonType
		questID quests.QuestID
		seed    int32
		// post.
		tiles        string
		pieces       string
		arches       string
		transparency string
	}{
		{
			// meta.
			dungeonName: "Cathedral",
			// pre.
			dlvl:    1,
			dtype:   gendung.Cathedral,
			questID: quests.Invalid,
			seed:    123,
			// post.
			tiles:        "12a0410904ebf2507b6b7017f0ae191ae476686b",
			pieces:       "e15a7afb7505cb01b0b3d1befce5b8d4833ae1c6",
			arches:       "5438e3d7761025a2ee6f7fec155c840fc289f5dd",
			transparency: "1269467cb381070f72bc6c8e69938e88da7e58cc",
		},
		{
			// meta.
			dungeonName: "The Butcher",
			// pre.
			dlvl:    2,
			dtype:   gendung.Cathedral,
			questID: quests.TheButcher,
			seed:    123,
			// post.
			tiles:        "659b95eec3e1c18d13b7f9932de108b88b356b9b",
			pieces:       "15f2209ff5d066cfd568a1eab77e4328d08474e8",
			arches:       "42941df3ada356ebf87ce2987d26a06c44da711a",
			transparency: "5f4e2e570b8a631d94fb3852c38ace0fa0397c7a",
		},
		{
			// meta.
			dungeonName: "Poisoned Water Supply",
			// pre.
			dlvl:    2,
			dtype:   gendung.Cathedral,
			questID: quests.PoisonedWaterSupply,
			seed:    123,
			// post.
			tiles:        "f4525775a47ef083d85c7faf5560b9808ce203ff",
			pieces:       "cb2f37c9d04a39ec22c4171c6e95c88a260364e3",
			arches:       "87418c244b8123dbdb3439812a2e1d8af5032c21",
			transparency: "5d548a45afb50e56cd77847fd832822eee0b01e7",
		},
	}

	for _, g := range golden {
		// Load level graphics.
		*gendung.DType = g.dtype
		diablo.LoadLevelGraphics()

		// Establish pre-conditions.
		*gendung.DLvl = g.dlvl
		*multi.MaxPlayers = 1
		for i := range quests.Quests {
			quests.Quests[i].Active = false
		}
		*gendung.IsQuestLevel = false
		if g.questID != quests.Invalid {
			quests.Quests[g.questID].Active = true
			quests.Quests[g.questID].DLvl = g.dlvl
		}
		// Generate dungeon based on the given seed.
		l1.CreateDungeon(g.seed, 0)
		if err := check(*gendung.TileIDMap, "tiles", g.seed, g.tiles); err != nil {
			t.Errorf("%s (dlvl=%d): %v", g.dungeonName, g.dlvl, errors.WithStack(err))
		}
		if err := check(*gendung.PieceIDMap, "pieces", g.seed, g.pieces); err != nil {
			t.Errorf("%s (dlvl=%d): %v", g.dungeonName, g.dlvl, errors.WithStack(err))
		}
		if err := check(*gendung.ArchNumMap, "arches", g.seed, g.arches); err != nil {
			t.Errorf("%s (dlvl=%d): %v", g.dungeonName, g.dlvl, errors.WithStack(err))
		}
		if err := check(*gendung.TransparencyMap, "transparency", g.seed, g.transparency); err != nil {
			t.Errorf("%s (dlvl=%d): %v", g.dungeonName, g.dlvl, errors.WithStack(err))
		}
	}
}

// check validates the data against the given SHA1 hashsum.
func check(data interface{}, name string, seed int32, want string) error {
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		return errors.WithStack(err)
	}
	sum := sha1.Sum(buf.Bytes())
	got := fmt.Sprintf("%040x", sum[:])
	if got != want {
		return errors.Errorf("SHA1 hash mismatch for %v, seed 0x%08X; expected %q, got %q", name, seed, want, got)
	}
	return nil
}
