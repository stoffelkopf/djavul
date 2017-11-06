package quests

import (
	"fmt"

	"github.com/pkg/errors"
)

// A QuestID represents a specific quest.
type QuestID uint8

// Quest IDs.
const (
	TheMagicRock         QuestID = 0
	BlackMushroom        QuestID = 1
	GharbadTheWeak       QuestID = 2
	ZharTheMad           QuestID = 3
	Lachdanan            QuestID = 4
	Diablo               QuestID = 5
	TheButcher           QuestID = 6
	OgdensSign           QuestID = 7
	HallsOfTheBlind      QuestID = 8
	Valor                QuestID = 9
	AnvilOfFury          QuestID = 10
	WarlordOfBlood       QuestID = 11
	TheCurseOfKingLeoric QuestID = 12
	PoisonedWaterSupply  QuestID = 13
	TheChamberOfBone     QuestID = 14
	ArchbishopLazarus    QuestID = 15
	Invalid              QuestID = 0xFF
)

// String returns a string representation of the quest ID.
func (q QuestID) String() string {
	switch q {
	case TheMagicRock:
		return "The Magic Rock"
	case BlackMushroom:
		return "Black Mushroom"
	case GharbadTheWeak:
		return "Gharbad the Weak"
	case ZharTheMad:
		return "Zhar the Mad"
	case Lachdanan:
		return "Lachdanan"
	case Diablo:
		return "Diablo"
	case TheButcher:
		return "The Butcher"
	case OgdensSign:
		return "Ogdens Sign"
	case HallsOfTheBlind:
		return "Halls of the Blind"
	case Valor:
		return "Valor"
	case AnvilOfFury:
		return "Anvil of Fury"
	case WarlordOfBlood:
		return "Warlord of Blood"
	case TheCurseOfKingLeoric:
		return "The Curse of King Leoric"
	case PoisonedWaterSupply:
		return "Poisoned Water Supply"
	case TheChamberOfBone:
		return "The Chamber of Bone"
	case ArchbishopLazarus:
		return "Archbishop Lazarus"
	case Invalid:
		return "<invalid quest ID>"
	default:
		return fmt.Sprintf("<unknown quest ID: %d>", uint8(q))
	}
}

// Sets sets the quest ID based on the given string representation.
func (q *QuestID) Set(s string) error {
	switch s {
	case "The Magic Rock":
		*q = TheMagicRock
	case "Black Mushroom":
		*q = BlackMushroom
	case "Gharbad the Weak":
		*q = GharbadTheWeak
	case "Zhar the Mad":
		*q = ZharTheMad
	case "Lachdanan":
		*q = Lachdanan
	case "Diablo":
		*q = Diablo
	case "The Butcher":
		*q = TheButcher
	case "Ogdens Sign":
		*q = OgdensSign
	case "Halls of the Blind":
		*q = HallsOfTheBlind
	case "Valor":
		*q = Valor
	case "Anvil of Fury":
		*q = AnvilOfFury
	case "Warlord of Blood":
		*q = WarlordOfBlood
	case "The Curse of King Leoric":
		*q = TheCurseOfKingLeoric
	case "Poisoned Water Supply":
		*q = PoisonedWaterSupply
	case "The Chamber of Bone":
		*q = TheChamberOfBone
	case "Archbishop Lazarus":
		*q = ArchbishopLazarus
	default:
		return errors.Errorf(`invalid quest ID string representation; expected "The Magic Rock", "Black Mushroom", "Gharbad the Weak", "Zhar the Mad", "Lachdanan", "Diablo", "The Butcher", "Ogdens Sign", "Halls of the Blind", "Valor", "Anvil of Fury", "Warlord of Blood", "The Curse of King Leoric", "Poisoned Water Supply", "The Chamber of Bone", or "Archbishop Lazarus", got %q`, s)
	}
	return nil
}
