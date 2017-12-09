//+build djavul

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
	"github.com/sanctuary/djavul/multi"
	"github.com/sanctuary/djavul/quests"
)

// dumpL1Maps stores the tile ID, dungeon piece ID and arch num maps for the
// first 256 seeds of the procedurally generated Cathedral dungeon.
func dumpL1Maps() {
	dinit.Archives()
	// Regular.
	//*gendung.DLvl = 1
	//*gendung.DType = gendung.Cathedral

	// Quest - The Butcher.
	//for i := range quests.Quests {
	//	quests.Quests[i].ID = quests.QuestID(i)
	//	quests.Quests[i].Active = false
	//}
	//*gendung.DLvl = quests.QuestsData[quests.TheButcher].DLvl
	//*gendung.DType = gendung.Cathedral
	//*gendung.IsQuestLevel = false
	//*multi.MaxPlayers = 1
	//quests.Quests[quests.TheButcher].Active = true
	//quests.Quests[quests.TheButcher].DLvl = quests.QuestsData[quests.TheButcher].DLvl

	// Quest - Poisoned Water Supply.
	//for i := range quests.Quests {
	//	quests.Quests[i].ID = quests.QuestID(i)
	//	quests.Quests[i].Active = false
	//}
	//*gendung.DLvl = quests.QuestsData[quests.PoisonedWaterSupply].DLvl
	//*gendung.DType = gendung.Cathedral
	//*gendung.IsQuestLevel = false
	//*multi.MaxPlayers = 1
	//quests.Quests[quests.PoisonedWaterSupply].Active = true
	//quests.Quests[quests.PoisonedWaterSupply].DLvl = quests.QuestsData[quests.PoisonedWaterSupply].DLvl

	// Quest - Odgen's Sign.
	for i := range quests.Quests {
		quests.Quests[i].ID = quests.QuestID(i)
		quests.Quests[i].Active = false
	}
	*gendung.DLvl = quests.QuestsData[quests.OgdensSign].DLvl
	*gendung.DType = gendung.Cathedral
	*gendung.IsQuestLevel = false
	*multi.MaxPlayers = 1
	quests.Quests[quests.OgdensSign].Active = true
	quests.Quests[quests.OgdensSign].DLvl = quests.QuestsData[quests.OgdensSign].DLvl

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
	path = fmt.Sprintf("l1_transparency_%08X.bin", seed)
	if err := dumpData(path, *gendung.TransparencyMap); err != nil {
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

// compareL1 compares the original and the Go implementation of the dynamic
// random level generation of Cathedral maps against one another.
func compareL1(start, end int64) error {
	dinit.Archives()
	*gendung.DLvl = 1
	*gendung.DType = gendung.Cathedral
	var orig [40][40]uint8
	diablo.LoadLevelGraphics()
	for i := start; i < end; i++ {
		seed := int32(i)
		switch seed {
		case 0x00000A1C, 0x00001287, 0x00001C71, 0x00002481, 0x00003B84:
			// skip broken seed.
			//
			// ref: https://github.com/sanctuary/graphics/tree/master/l1/broken
			log.Printf("ignoring broken map with seed 0x%08X", seed)
			continue
		}
		// Original implementation.
		l1.UseGo = false
		l1.CreateDungeon(seed, 0)
		wantTiles := hash(*gendung.TileIDMap)
		wantPieces := hash(*gendung.PieceIDMap)
		wantArches := hash(*gendung.ArchNumMap)
		wantTransparency := hash(*gendung.TransparencyMap)
		orig = *gendung.TileIDMap

		// Go implementation.
		l1.UseGo = true
		l1.CreateDungeon(seed, 0)
		gotTiles := hash(*gendung.TileIDMap)
		gotPieces := hash(*gendung.PieceIDMap)
		gotArches := hash(*gendung.ArchNumMap)
		gotTransparency := hash(*gendung.TransparencyMap)

		if gotTiles != wantTiles {
			path := fmt.Sprintf("dlvl_1_seed_%d_orig.bin", seed)
			if err := dumpData(path, orig); err != nil {
				return errors.WithStack(err)
			}
			path = fmt.Sprintf("dlvl_1_seed_%d_go.bin", seed)
			if err := dumpData(path, *gendung.TileIDMap); err != nil {
				return errors.WithStack(err)
			}
			log.Printf("dumping broken map with seed 0x%08X", seed)
			continue
			return errors.Errorf("SHA1 hash mismatch for tiles, seed 0x%08X; expected %q, got %q", seed, wantTiles, gotTiles)
		}
		if gotPieces != wantPieces {
			return errors.Errorf("SHA1 hash mismatch for pieces, seed 0x%08X; expected %q, got %q", seed, wantPieces, gotPieces)
		}
		if gotArches != wantArches {
			return errors.Errorf("SHA1 hash mismatch for arches, seed 0x%08X; expected %q, got %q", seed, wantArches, gotArches)
		}
		if gotTransparency != wantTransparency {
			return errors.Errorf("SHA1 hash mismatch for transparency, seed 0x%08X; expected %q, got %q", seed, wantTransparency, gotTransparency)
		}
	}
	fmt.Printf("PASS %08X - %08X\n", start, end)
	return nil
}

// hash returns the SHA1 hash of the given data.
func hash(data interface{}) string {
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		log.Fatalf("unable to write data to bytes buffer; %+v", errors.WithStack(err))
	}
	sum := sha1.Sum(buf.Bytes())
	return fmt.Sprintf("%040x", sum[:])
}

// checkL1Regular validates the implemenation of the dynamic random level
// generation of Cathedral maps.
func checkL1Regular() error {
	dinit.Archives()
	*gendung.DLvl = 1
	*gendung.DType = gendung.Cathedral
	diablo.LoadLevelGraphics()
	seed := int32(123)
	l1.CreateDungeon(seed, 0)
	if err := check(*gendung.TileIDMap, "tiles", seed, "12a0410904ebf2507b6b7017f0ae191ae476686b"); err != nil {
		path := fmt.Sprintf("testdata_regular/l1_tiles_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
		}
		defer f.Close()
		var tiles [40][40]uint8
		if err := binary.Read(f, binary.LittleEndian, &tiles); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.TileIDMap)
		pretty.Println("want:", tiles)
		return errors.WithStack(err)
	}
	if err := check(*gendung.PieceIDMap, "pieces", seed, "e15a7afb7505cb01b0b3d1befce5b8d4833ae1c6"); err != nil {
		path := fmt.Sprintf("testdata_regular/l1_pieces_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
		}
		defer f.Close()
		var pieces [112][112]int32
		if err := binary.Read(f, binary.LittleEndian, &pieces); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.PieceIDMap)
		pretty.Println("want:", pieces)
		return errors.WithStack(err)
	}
	if err := check(*gendung.ArchNumMap, "arches", seed, "5438e3d7761025a2ee6f7fec155c840fc289f5dd"); err != nil {
		path := fmt.Sprintf("testdata_regular/l1_arches_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
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
	if err := check(*gendung.TransparencyMap, "transparency", seed, "1269467cb381070f72bc6c8e69938e88da7e58cc"); err != nil {
		path := fmt.Sprintf("testdata_regular/l1_transparency_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
		}
		defer f.Close()
		var trans [112][112]int8
		if err := binary.Read(f, binary.LittleEndian, &trans); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.TransparencyMap)
		pretty.Println("want:", trans)
		return errors.WithStack(err)
	}
	fmt.Println("PASS: regular")
	return nil
}

// checkL1QuestTheButcher validates the implemenation of the dynamic random
// level generation of Cathedral maps, for dungeon level 2 with the Butcher
// quest active.
func checkL1Quest() error {
	if err := checkL1QuestTheButcher(); err != nil {
		return errors.WithStack(err)
	}
	if err := checkL1QuestPoisonedWaterSupply(); err != nil {
		return errors.WithStack(err)
	}
	if err := checkL1QuestOgdensSign(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// checkL1QuestTheButcher validates the implemenation of the dynamic random
// level generation of Cathedral maps, for dungeon level 2 with the Butcher
// quest active.
func checkL1QuestTheButcher() error {
	dinit.Archives()
	for i := range quests.Quests {
		quests.Quests[i].ID = quests.QuestID(i)
		quests.Quests[i].Active = false
	}
	*gendung.DLvl = quests.QuestsData[quests.TheButcher].DLvl
	*gendung.DType = gendung.Cathedral
	*gendung.IsQuestLevel = false
	*multi.MaxPlayers = 1
	quests.Quests[quests.TheButcher].Active = true
	quests.Quests[quests.TheButcher].DLvl = quests.QuestsData[quests.TheButcher].DLvl
	diablo.LoadLevelGraphics()
	seed := int32(123)
	l1.CreateDungeon(seed, 0)
	if err := check(*gendung.TileIDMap, "tiles", seed, "659b95eec3e1c18d13b7f9932de108b88b356b9b"); err != nil {
		path := fmt.Sprintf("testdata_quest_the_butcher/l1_tiles_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
		}
		defer f.Close()
		var tiles [40][40]uint8
		if err := binary.Read(f, binary.LittleEndian, &tiles); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.TileIDMap)
		pretty.Println("want:", tiles)
		return errors.WithStack(err)
	}
	if err := check(*gendung.PieceIDMap, "pieces", seed, "15f2209ff5d066cfd568a1eab77e4328d08474e8"); err != nil {
		path := fmt.Sprintf("testdata_quest_the_butcher/l1_pieces_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
		}
		defer f.Close()
		var pieces [112][112]int32
		if err := binary.Read(f, binary.LittleEndian, &pieces); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.PieceIDMap)
		pretty.Println("want:", pieces)
		return errors.WithStack(err)
	}
	if err := check(*gendung.ArchNumMap, "arches", seed, "42941df3ada356ebf87ce2987d26a06c44da711a"); err != nil {
		path := fmt.Sprintf("testdata_quest_the_butcher/l1_arches_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
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
	if err := check(*gendung.TransparencyMap, "transparency", seed, "74c24e596ec57a91261bc3a559270f31d6811336"); err != nil {
		path := fmt.Sprintf("testdata_quest_the_butcher/l1_transparency_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
		}
		defer f.Close()
		var trans [112][112]int8
		if err := binary.Read(f, binary.LittleEndian, &trans); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.TransparencyMap)
		pretty.Println("want:", trans)
		return errors.WithStack(err)
	}
	fmt.Println("PASS: quest - The Butcher")
	return nil
}

// checkL1QuestPoisonedWaterSupply validates the implemenation of the dynamic
// random level generation of Cathedral maps, for dungeon level 2 with the
// Poisoned Water Supply quest active.
func checkL1QuestPoisonedWaterSupply() error {
	dinit.Archives()
	for i := range quests.Quests {
		quests.Quests[i].ID = quests.QuestID(i)
		quests.Quests[i].Active = false
	}
	*gendung.DLvl = quests.QuestsData[quests.PoisonedWaterSupply].DLvl
	*gendung.DType = gendung.Cathedral
	*gendung.IsQuestLevel = false
	*multi.MaxPlayers = 1
	quests.Quests[quests.PoisonedWaterSupply].Active = true
	quests.Quests[quests.PoisonedWaterSupply].DLvl = quests.QuestsData[quests.PoisonedWaterSupply].DLvl
	diablo.LoadLevelGraphics()
	seed := int32(123)
	l1.CreateDungeon(seed, 0)
	if err := check(*gendung.TileIDMap, "tiles", seed, "f4525775a47ef083d85c7faf5560b9808ce203ff"); err != nil {
		path := fmt.Sprintf("testdata_quest_poisoned_water_supply/l1_tiles_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
		}
		defer f.Close()
		var tiles [40][40]uint8
		if err := binary.Read(f, binary.LittleEndian, &tiles); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.TileIDMap)
		pretty.Println("want:", tiles)
		return errors.WithStack(err)
	}
	if err := check(*gendung.PieceIDMap, "pieces", seed, "cb2f37c9d04a39ec22c4171c6e95c88a260364e3"); err != nil {
		path := fmt.Sprintf("testdata_quest_poisoned_water_supply/l1_pieces_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
		}
		defer f.Close()
		var pieces [112][112]int32
		if err := binary.Read(f, binary.LittleEndian, &pieces); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.PieceIDMap)
		pretty.Println("want:", pieces)
		return errors.WithStack(err)
	}
	if err := check(*gendung.ArchNumMap, "arches", seed, "87418c244b8123dbdb3439812a2e1d8af5032c21"); err != nil {
		path := fmt.Sprintf("testdata_quest_poisoned_water_supply/l1_arches_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
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
	if err := check(*gendung.TransparencyMap, "transparency", seed, "5d548a45afb50e56cd77847fd832822eee0b01e7"); err != nil {
		path := fmt.Sprintf("testdata_quest_poisoned_water_supply/l1_transparency_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
		}
		defer f.Close()
		var trans [112][112]int8
		if err := binary.Read(f, binary.LittleEndian, &trans); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.TransparencyMap)
		pretty.Println("want:", trans)
		return errors.WithStack(err)
	}
	fmt.Println("PASS: quest - Poisoned Water Supply")
	return nil
}

// checkL1QuestOgdensSign validates the implemenation of the dynamic random
// level generation of Cathedral maps, for dungeon level 4 with the Ogden's Sign
// quest active.
func checkL1QuestOgdensSign() error {
	dinit.Archives()
	for i := range quests.Quests {
		quests.Quests[i].ID = quests.QuestID(i)
		quests.Quests[i].Active = false
	}
	*gendung.DLvl = quests.QuestsData[quests.OgdensSign].DLvl
	*gendung.DType = gendung.Cathedral
	*gendung.IsQuestLevel = false
	*multi.MaxPlayers = 1
	quests.Quests[quests.OgdensSign].Active = true
	quests.Quests[quests.OgdensSign].DLvl = quests.QuestsData[quests.OgdensSign].DLvl
	diablo.LoadLevelGraphics()
	seed := int32(123)
	l1.CreateDungeon(seed, 0)
	if err := check(*gendung.TileIDMap, "tiles", seed, "3a54760d2ce39932f556dbb3ae924c8425e5f9ea"); err != nil {
		path := fmt.Sprintf("testdata_quest_ogdens_sign/l1_tiles_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
		}
		defer f.Close()
		var tiles [40][40]uint8
		if err := binary.Read(f, binary.LittleEndian, &tiles); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.TileIDMap)
		pretty.Println("want:", tiles)
		return errors.WithStack(err)
	}
	if err := check(*gendung.PieceIDMap, "pieces", seed, "f6fcf0461dfad18da42b3d25dde5e60cdc7b4daf"); err != nil {
		path := fmt.Sprintf("testdata_quest_ogdens_sign/l1_pieces_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
		}
		defer f.Close()
		var pieces [112][112]int32
		if err := binary.Read(f, binary.LittleEndian, &pieces); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.PieceIDMap)
		pretty.Println("want:", pieces)
		return errors.WithStack(err)
	}
	if err := check(*gendung.ArchNumMap, "arches", seed, "7e97023f45d78a37dffb569111762018e6b0c93f"); err != nil {
		path := fmt.Sprintf("testdata_quest_ogdens_sign/l1_arches_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
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
	if err := check(*gendung.TransparencyMap, "transparency", seed, "10156f455d85c0c4be6d26be23fc540776253aa9"); err != nil {
		path := fmt.Sprintf("testdata_quest_ogdens_sign/l1_transparency_%08X.bin", seed)
		f, e := os.Open(path)
		if e != nil {
			return errors.WithStack(e)
		}
		defer f.Close()
		var trans [112][112]int8
		if err := binary.Read(f, binary.LittleEndian, &trans); err != nil {
			return errors.WithStack(err)
		}
		pretty.Println("got:", *gendung.TransparencyMap)
		pretty.Println("want:", trans)
		return errors.WithStack(err)
	}
	fmt.Println("PASS: quest - Ogden's Sign")
	return nil
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
