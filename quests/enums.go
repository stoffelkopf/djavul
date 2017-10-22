package quests

// A QuestID represents a specific quest.
type QuestID uint8

// Quest IDs.
const (
	TheMagicRock QuestID = iota
	BlackMushroom
	GharbadTheWeak
	ZharTheMad
	Lachdanan
	Diablo
	TheButcher
	OgdensSign
	HallsOfTheBlind
	Valor
	AnvilOfFury
	WarlordOfBlood
	TheCurseOfKingLeoric
	PoisonedWaterSupply
	TheChamberOfBone
	ArchbishopLazarus
	Invalid QuestID = 0xFF
)
