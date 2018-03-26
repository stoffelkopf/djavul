//+build !djavul

// Global variable wrappers for quests.cpp

package quests

import (
	"log"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/internal/parse"
)

// Global variables.
var (
	// QuestsData contains the data related to each quest ID.
	//
	// References:
	//    * https://github.com/sanctuary/notes/blob/master/enums.h#quest_id
	//
	// ref: 0x4A1AE0
	QuestsData = new([16]QuestData)

	// Quests contains the quests of the current game.
	//
	// PSX ref: 0x800DDA40
	// PSX def: QuestStruct quests[16];
	//
	// ref: 0x69BD10
	Quests = new([16]Quest)
)

// init initializes read-only data of structures from DIABLO.EXE.
func init() {
	if err := initDiabloStructs(); err != nil {
		log.Fatalf("%+v", err)
	}
}

// initDiabloStructs initializes read-only data of structures from DIABLO.EXE.
func initDiabloStructs() error {
	if err := parse.Data(parse.Offset(0x4A1AE0), QuestsData); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
