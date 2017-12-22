package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// outputHero outputs the contents of the Hero struct in JSON format.
func outputHero(data []byte) error {
	var hero Hero
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &hero); err != nil {
		return errors.WithStack(err)
	}
	buf, err := json.MarshalIndent(hero, "", "\t")
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Println(string(buf))
	return nil
}

// Hero contains the most essential information about a player, as used in save
// files.
type Hero struct { // size = 0x4F2
	// padding
	_ [8]byte
	// offset 0008 (1 bytes)
	DAction int8
	// offset 0009 (1 bytes)
	Param1 int8
	// offset 000A (1 bytes)
	Param2 int8
	// offset 000B (1 bytes)
	DLvl int8
	// offset 000C (1 bytes)
	X int8
	// offset 000D (1 bytes)
	Y int8
	// offset 000E (1 bytes)
	TargetX int8
	// offset 000F (1 bytes)
	TargetY int8
	// offset 0010 (32 bytes)
	Name String32
	// offset 0030 (1 bytes)
	PlayerClass PlayerClass
	// offset 0031 (1 bytes)
	StrBase int8
	// offset 0032 (1 bytes)
	MagBase int8
	// offset 0033 (1 bytes)
	DexBase int8
	// offset 0034 (1 bytes)
	VitBase int8
	// offset 0035 (1 bytes)
	CLvl int8
	// offset 0036 (1 bytes)
	Points int8
	// offset 0037 (4 bytes)
	Exp int32
	// offset 003B (4 bytes)
	GoldTotal int32
	// offset 003F (4 bytes)
	HPBaseCur int32
	// offset 0043 (4 bytes)
	HPBaseMax int32
	// offset 0047 (4 bytes)
	MPBaseCur int32
	// offset 004B (4 bytes)
	MPBaseMax int32
	// offset 004F (37 bytes)
	SpellLvlFromSpellID [37]int8
	// offset 0074 (8 bytes)
	KnownSpells [2]uint32 // bitfield of known spells.
	// offset 007C (133 bytes)
	BodyItems [7]HeroItem
	// offset 0101 (760 bytes)
	InvItems [40]HeroItem
	// offset 03F9 (40 bytes)
	InvNumFromInvGrid [40]int8
	// offset 0421 (1 bytes)
	NInvItems int8
	// offset 0422 (152 bytes)
	BeltItems [8]HeroItem
	// padding
	_ [3]byte
	// offset 04BD (1 bytes)
	OnBattlenet int8
	// offset 04BE (1 bytes)
	HasManashild int8
	// padding
	_ [19]byte
	// offset 04D2 (4 bytes)
	Difficulty int32
	// padding
	_ [18]byte
}

// HeroItem contains the most essential information about the item of a player,
// as used in save files.
type HeroItem struct { // size = 0x13
	// offset 0000 (bytes 4)
	Seed int32
	// offset 0004 (bytes 2)
	CF uint16
	// offset 0006 (bytes 2)
	ItemID ItemID
	// offset 0008 (bytes 1)
	IdentifiedAndItemQuality uint8 // The first bit corresponds to identified and the remaining bits corresponds to item_quality.
	// offset 0009 (bytes 1)
	DurabilityCur int8
	// offset 000A (bytes 1)
	DurabilityMax int8
	// offset 000B (bytes 1)
	ChargesMin int8
	// offset 000C (bytes 1)
	ChargesMax int8
	// offset 000D (bytes 2)
	GoldPrice int16
	// offset 000F (bytes 4)
	OnlyUsedByEar uint32 // Stores the last 4 bytes of the ear name.
}

//go:generate stringer -linecomment -type PlayerClass

// PlayerClass specifies the class of a player.
type PlayerClass uint8

// MarshalText encodes the receiver into UTF-8-encoded text and returns the
// result.
func (c PlayerClass) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

// Player classes.
const (
	PlayerClassWarrior  PlayerClass = 0 // Warrior
	PlayerClassRogue    PlayerClass = 1 // Rogue
	PlayerClassSorceror PlayerClass = 2 // Sorceror
)

//go:generate stringer -linecomment -type ItemID

// ItemID represents the set of items.
type ItemID int16

// MarshalText encodes the receiver into UTF-8-encoded text and returns the
// result.
func (i ItemID) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// Item IDs.
const (
	ItemIDGold                         ItemID = 0   // Gold
	ItemIDShortSword                   ItemID = 1   // Short Sword
	ItemIDBuckler                      ItemID = 2   // Buckler
	ItemIDClub                         ItemID = 3   // Club
	ItemIDShortBow                     ItemID = 4   // Short Bow
	ItemIDShortStaffOfChargedBolt      ItemID = 5   // Short Staff of Charged Bolt
	ItemIDCleaver                      ItemID = 6   // Cleaver
	ItemIDTheUndeadCrown               ItemID = 7   // The Undead Crown
	ItemIDEmpyreanBand                 ItemID = 8   // Empyrean Band
	ItemIDMagicRock                    ItemID = 9   // Magic Rock
	ItemIDOpticAmulet                  ItemID = 10  // Optic Amulet
	ItemIDRingOfTruth                  ItemID = 11  // Ring of Truth
	ItemIDTavernSign                   ItemID = 12  // Tavern Sign
	ItemIDHarlequinCrest               ItemID = 13  // Harlequin Crest
	ItemIDVeilOfSteel                  ItemID = 14  // Veil of Steel
	ItemIDGoldenElixir                 ItemID = 15  // Golden Elixir
	ItemIDAnvilOfFury                  ItemID = 16  // Anvil of Fury
	ItemIDBlackMushroom                ItemID = 17  // Black Mushroom
	ItemIDBrain                        ItemID = 18  // Brain
	ItemIDFungalTome                   ItemID = 19  // Fungal Tome
	ItemIDSpectralElixir               ItemID = 20  // Spectral Elixir
	ItemIDBloodStone                   ItemID = 21  // Blood Stone
	ItemIDMapOfTheStars                ItemID = 22  // Map of The Stars
	ItemIDHeart                        ItemID = 23  // Heart
	ItemIDPotionOfHealing              ItemID = 24  // Potion of Healing
	ItemIDPotionOfMana                 ItemID = 25  // Potion of Mana
	ItemIDScrollOfIdentify             ItemID = 26  // Scroll of Identify
	ItemIDScrollOfTownPortal           ItemID = 27  // Scroll of Town Portal
	ItemIDArkainesValor                ItemID = 28  // Arkaine's Valor
	ItemIDPotionOfFullHealing          ItemID = 29  // Potion of Full Healing
	ItemIDPotionOfFullMana             ItemID = 30  // Potion of Full Mana
	ItemIDGriswoldsEdge                ItemID = 31  // Griswold's Edge
	ItemIDLightforge                   ItemID = 32  // Lightforge
	ItemIDStaffOfLazarus               ItemID = 33  // Staff of Lazarus
	ItemIDScrollOfResurrect            ItemID = 34  // Scroll of Resurrect
	ItemIDNull1                        ItemID = 35  // Null 1
	ItemIDNull2                        ItemID = 36  // Null 2
	ItemIDNull3                        ItemID = 37  // Null 3
	ItemIDNull4                        ItemID = 38  // Null 4
	ItemIDNull5                        ItemID = 39  // Null 5
	ItemIDNull6                        ItemID = 40  // Null 6
	ItemIDNull7                        ItemID = 41  // Null 7
	ItemIDNull8                        ItemID = 42  // Null 8
	ItemIDNull9                        ItemID = 43  // Null 9
	ItemIDNull10                       ItemID = 44  // Null 10
	ItemIDNull11                       ItemID = 45  // Null 11
	ItemIDNull12                       ItemID = 46  // Null 12
	ItemIDNull13                       ItemID = 47  // Null 13
	ItemIDBaseCap                      ItemID = 48  // Cap (base item)
	ItemIDBaseSkullCap                 ItemID = 49  // Skull Cap (base item)
	ItemIDBaseHelm                     ItemID = 50  // Helm (base item)
	ItemIDBaseFullHelm                 ItemID = 51  // Full Helm (base item)
	ItemIDBaseCrown                    ItemID = 52  // Crown (base item)
	ItemIDBaseGreatHelm                ItemID = 53  // Great Helm (base item)
	ItemIDBaseCape                     ItemID = 54  // Cape (base item)
	ItemIDBaseRags                     ItemID = 55  // Rags (base item)
	ItemIDBaseCloak                    ItemID = 56  // Cloak (base item)
	ItemIDBaseRobe                     ItemID = 57  // Robe (base item)
	ItemIDBaseQuiltedArmor             ItemID = 58  // Quilted Armor (base item)
	ItemIDBaseLeatherArmor             ItemID = 59  // Leather Armor (base item)
	ItemIDBaseHardLeatherArmor         ItemID = 60  // Hard Leather Armor (base item)
	ItemIDBaseStuddedLeatherArmor      ItemID = 61  // Studded Leather Armor (base item)
	ItemIDBaseRingMail                 ItemID = 62  // Ring Mail (base item)
	ItemIDBaseChainMail                ItemID = 63  // Chain Mail (base item)
	ItemIDBaseScaleMail                ItemID = 64  // Scale Mail (base item)
	ItemIDBaseBreastPlate              ItemID = 65  // Breast Plate (base item)
	ItemIDBaseSplintMail               ItemID = 66  // Splint Mail (base item)
	ItemIDBasePlateMail                ItemID = 67  // Plate Mail (base item)
	ItemIDBaseFieldPlate               ItemID = 68  // Field Plate (base item)
	ItemIDBaseGothicPlate              ItemID = 69  // Gothic Plate (base item)
	ItemIDBaseFullPlateMail            ItemID = 70  // Full Plate Mail (base item)
	ItemIDBaseBuckler                  ItemID = 71  // Buckler (base item)
	ItemIDBaseSmallShield              ItemID = 72  // Small Shield (base item)
	ItemIDBaseLargeShield              ItemID = 73  // Large Shield (base item)
	ItemIDBaseKiteShield               ItemID = 74  // Kite Shield (base item)
	ItemIDBaseTowerShield              ItemID = 75  // Tower Shield (base item)
	ItemIDBaseGothicShield             ItemID = 76  // Gothic Shield (base item)
	ItemIDBasePotionOfHealing          ItemID = 77  // Potion of Healing (base item)
	ItemIDBasePotionOfFullHealing      ItemID = 78  // Potion of Full Healing (base item)
	ItemIDBasePotionOfMana             ItemID = 79  // Potion of Mana (base item)
	ItemIDBasePotionOfFullMana         ItemID = 80  // Potion of Full Mana (base item)
	ItemIDBasePotionOfRejuvenation     ItemID = 81  // Potion of Rejuvenation (base item)
	ItemIDBasePotionOfFullRejuvenation ItemID = 82  // Potion of Full Rejuvenation (base item)
	ItemIDBaseElixirOfStrength         ItemID = 83  // Elixir of Strength (base item)
	ItemIDBaseElixirOfMagic            ItemID = 84  // Elixir of Magic (base item)
	ItemIDBaseElixirOfDexterity        ItemID = 85  // Elixir of Dexterity (base item)
	ItemIDBaseElixirOfVitality         ItemID = 86  // Elixir of Vitality (base item)
	ItemIDBaseScrollOfHealing          ItemID = 87  // Scroll of Healing (base item)
	ItemIDBaseScrollOfLightning        ItemID = 88  // Scroll of Lightning (base item)
	ItemIDBaseScrollOfIdentify         ItemID = 89  // Scroll of Identify (base item)
	ItemIDBaseScrollOfResurrect        ItemID = 90  // Scroll of Resurrect (base item)
	ItemIDBaseScrollOfFireWall         ItemID = 91  // Scroll of Fire Wall (base item)
	ItemIDBaseScrollOfInferno          ItemID = 92  // Scroll of Inferno (base item)
	ItemIDBaseScrollOfTownPortal       ItemID = 93  // Scroll of Town Portal (base item)
	ItemIDBaseScrollOfFlash            ItemID = 94  // Scroll of Flash (base item)
	ItemIDBaseScrollOfInfravision      ItemID = 95  // Scroll of Infravision (base item)
	ItemIDBaseScrollOfPhasing          ItemID = 96  // Scroll of Phasing (base item)
	ItemIDBaseScrollOfManaShield       ItemID = 97  // Scroll of Mana Shield (base item)
	ItemIDBaseScrollOfFlameWave        ItemID = 98  // Scroll of Flame Wave (base item)
	ItemIDBaseScrollOfFireball         ItemID = 99  // Scroll of Fireball (base item)
	ItemIDBaseScrollOfStoneCurse       ItemID = 100 // Scroll of Stone Curse (base item)
	ItemIDBaseScrollOfChainLightning   ItemID = 101 // Scroll of Chain Lightning (base item)
	ItemIDBaseScrollOfGuardian         ItemID = 102 // Scroll of Guardian (base item)
	ItemIDBaseNonItem                  ItemID = 103 // Non Item (base item)
	ItemIDBaseScrollOfNova             ItemID = 104 // Scroll of Nova (base item)
	ItemIDBaseScrollOfGolem            ItemID = 105 // Scroll of Golem (base item)
	ItemIDBaseScrollOfNone             ItemID = 106 // Scroll of None (base item)
	ItemIDBaseScrollOfTeleport         ItemID = 107 // Scroll of Teleport (base item)
	ItemIDBaseScrollOfApocalypse       ItemID = 108 // Scroll of Apocalypse (base item)
	ItemIDBaseBookQlvl2                ItemID = 109 // Book Qlvl 2 (base item)
	ItemIDBaseBookQlvl8                ItemID = 110 // Book Qlvl 8 (base item)
	ItemIDBaseBookQlvl14               ItemID = 111 // Book Qlvl 14 (base item)
	ItemIDBaseBookQlvl20               ItemID = 112 // Book Qlvl 20 (base item)
	ItemIDBaseDagger                   ItemID = 113 // Dagger (base item)
	ItemIDBaseShortSword               ItemID = 114 // Short Sword (base item)
	ItemIDBaseFalchion                 ItemID = 115 // Falchion (base item)
	ItemIDBaseScimitar                 ItemID = 116 // Scimitar (base item)
	ItemIDBaseClaymore                 ItemID = 117 // Claymore (base item)
	ItemIDBaseBlade                    ItemID = 118 // Blade (base item)
	ItemIDBaseSabre                    ItemID = 119 // Sabre (base item)
	ItemIDBaseLongSword                ItemID = 120 // Long Sword (base item)
	ItemIDBaseBroadSword               ItemID = 121 // Broad Sword (base item)
	ItemIDBaseBastardSword             ItemID = 122 // Bastard Sword (base item)
	ItemIDBaseTwoHandedSword           ItemID = 123 // Two Handed Sword (base item)
	ItemIDBaseGreatSword               ItemID = 124 // Great Sword (base item)
	ItemIDBaseSmallAxe                 ItemID = 125 // Small Axe (base item)
	ItemIDBaseAxe                      ItemID = 126 // Axe (base item)
	ItemIDBaseLargeAxe                 ItemID = 127 // Large Axe (base item)
	ItemIDBaseBroadAxe                 ItemID = 128 // Broad Axe (base item)
	ItemIDBaseBattleAxe                ItemID = 129 // Battle Axe (base item)
	ItemIDBaseGreatAxe                 ItemID = 130 // Great Axe (base item)
	ItemIDBaseMace                     ItemID = 131 // Mace (base item)
	ItemIDBaseMorningStar              ItemID = 132 // Morning Star (base item)
	ItemIDBaseWarHammer                ItemID = 133 // War Hammer (base item)
	ItemIDBaseSpikedClub               ItemID = 134 // Spiked Club (base item)
	ItemIDBaseClub                     ItemID = 135 // Club (base item)
	ItemIDBaseFlail                    ItemID = 136 // Flail (base item)
	ItemIDBaseMaul                     ItemID = 137 // Maul (base item)
	ItemIDBaseShortBow                 ItemID = 138 // Short Bow (base item)
	ItemIDBaseHuntersBow               ItemID = 139 // Hunter's Bow (base item)
	ItemIDBaseLongBow                  ItemID = 140 // Long Bow (base item)
	ItemIDBaseCompositeBow             ItemID = 141 // Composite Bow (base item)
	ItemIDBaseShortBattleBow           ItemID = 142 // Short Battle Bow (base item)
	ItemIDBaseLongBattleBow            ItemID = 143 // Long Battle Bow (base item)
	ItemIDBaseShortWarBow              ItemID = 144 // Short War Bow (base item)
	ItemIDBaseLongWarBow               ItemID = 145 // Long War Bow (base item)
	ItemIDBaseShortStaff               ItemID = 146 // Short Staff (base item)
	ItemIDBaseLongStaff                ItemID = 147 // Long Staff (base item)
	ItemIDBaseCompositeStaff           ItemID = 148 // Composite Staff (base item)
	ItemIDBaseQuarterStaff             ItemID = 149 // Quarter Staff (base item)
	ItemIDBaseWarStaff                 ItemID = 150 // War Staff (base item)
	ItemIDBaseRingQlvl5                ItemID = 151 // Ring Qlvl 5 (base item)
	ItemIDBaseRingQlvl10               ItemID = 152 // Ring Qlvl 10 (base item)
	ItemIDBaseRingQlvl15               ItemID = 153 // Ring Qlvl 15 (base item)
	ItemIDBaseAmuletQlvl8              ItemID = 154 // Amulet Qlvl 8 (base item)
	ItemIDBaseAmuletQlvl16             ItemID = 155 // Amulet Qlvl 16 (base item)
	ItemIDNull14                       ItemID = 156 // Null 14
	ItemIDNone                         ItemID = -1  // None
)

// ### [ Helper functions ] ####################################################

// String32 represents a 32-byte NULL-padded string.
type String32 [32]byte

// MarshalText encodes the receiver into UTF-8-encoded text and returns the
// result.
func (b String32) MarshalText() ([]byte, error) {
	pos := bytes.IndexByte(b[:], 0)
	if pos == -1 {
		pos = 0
	}
	return b[:pos], nil
}
