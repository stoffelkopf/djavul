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
	}

	for _, g := range golden {
		// Load level graphics.
		*gendung.DType = g.dtype
		diablo.LoadLevelGraphics()

		// Establish pre-conditions.
		*gendung.DLvl = g.dlvl
		if g.questID != quests.Invalid {
			*multi.MaxPlayers = 1
			for i := range quests.Quests {
				quests.Quests[i].Active = false
			}
			*gendung.IsQuestLevel = false
			quests.Quests[g.questID].Active = true
			quests.Quests[g.questID].DLvl = g.dlvl
		}
		// Generate dungeon based on the given seed.
		l1.CreateDungeon(g.seed, 0)
		if err := check(*gendung.TileIDMap, "tiles", g.seed, "12a0410904ebf2507b6b7017f0ae191ae476686b"); err != nil {
			t.Errorf("%s (dlvl=%d): %v", g.dungeonName, g.dlvl, errors.WithStack(err))
			continue
		}
		if err := check(*gendung.PieceIDMap, "pieces", g.seed, "e15a7afb7505cb01b0b3d1befce5b8d4833ae1c6"); err != nil {
			t.Errorf("%s (dlvl=%d): %v", g.dungeonName, g.dlvl, errors.WithStack(err))
			continue
		}
		if err := check(*gendung.ArchNumMap, "arches", g.seed, "5438e3d7761025a2ee6f7fec155c840fc289f5dd"); err != nil {
			t.Errorf("%s (dlvl=%d): %v", g.dungeonName, g.dlvl, errors.WithStack(err))
			continue
		}
		if err := check(*gendung.TransparencyMap, "transparency", g.seed, "1269467cb381070f72bc6c8e69938e88da7e58cc"); err != nil {
			t.Errorf("%s (dlvl=%d): %v", g.dungeonName, g.dlvl, errors.WithStack(err))
			continue
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
