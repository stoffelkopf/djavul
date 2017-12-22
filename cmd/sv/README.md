# sv

The sv tool decodes Diablo 1 save files.

## Installation

```bash
go get github.com/sanctuary/djavul/cmd/sv
```

## Usage

```bash
# Decode the contents of a single player `game` save file.
$ sv game | hexdump -C

00000000  52 45 54 4c 00 00 00 00  00 00 00 00 00 00 00 00  |RETL............|
00000010  00 00 00 00 4b 00 00 00  44 00 00 00 00 00 00 00  |....K...D.......|
...
00000a20  00 00 00 00 00 00 00 00  00 00 00 00 53 68 6f 72  |............Shor|
00000a30  74 20 53 77 6f 72 64 00  00 00 00 00 00 00 00 00  |t Sword.........|
...
```

```bash
# Decode the contents of a multi player `hero` save file.
$ sv -p "szqnlsk1" hero | hexdump -C

00000000  00 00 00 00 00 00 00 00  ff 00 00 00 49 41 49 41  |............IAIA|
00000010  66 6f 6f 62 61 72 00 00  00 00 00 00 00 00 00 00  |foobar..........|
...
```

```bash
# Output the contents of the `hero` save file in JSON format.
$ sv -json hero

{
    "DAction": -1,
    "Param1": 0,
    "Param2": 0,
    "DLvl": 0,
    "X": 75,
    "Y": 68,
    "TargetX": 75,
    "TargetY": 68,
    "Name": "a",
    "PlayerClass": "Warrior",
    "StrBase": 30,
    "MagBase": 10,
    "DexBase": 20,
    "VitBase": 25,
    "CLvl": 1,
    "Points": 0,
    "Exp": 0,
    "GoldTotal": 100,
    "HPBaseCur": 4480,
    "HPBaseMax": 4480,
    "MPBaseCur": 640,
    "MPBaseMax": 640,
    "SpellLvlFromSpellID": [
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
    ],
    "KnownSpells": [
        0, 0
    ],
    "BodyItems": [
        {
            "Seed": 1217677573,
            "CF": 0,
            "ItemID": "Short Sword",
            "IdentifiedAndItemQuality": 0,
            "DurabilityCur": 20,
            "DurabilityMax": 20,
            "ChargesMin": 0,
            "ChargesMax": 0,
            "GoldPrice": 0,
            "OnlyUsedByEar": 0
        },
        {
            "Seed": 441146358,
            "CF": 0,
            "ItemID": "Buckler",
            "IdentifiedAndItemQuality": 0,
            "DurabilityCur": 10,
            "DurabilityMax": 10,
            "ChargesMin": 0,
            "ChargesMax": 0,
            "GoldPrice": 0,
            "OnlyUsedByEar": 0
        }
    ],
    "InvItems": [
        {
            "Seed": 1306853907,
            "CF": 0,
            "ItemID": "Club",
            "IdentifiedAndItemQuality": 0,
            "DurabilityCur": 20,
            "DurabilityMax": 20,
            "ChargesMin": 0,
            "ChargesMax": 0,
            "GoldPrice": 0,
            "OnlyUsedByEar": 0
        },
        {
            "Seed": 808692058,
            "CF": 0,
            "ItemID": "Gold",
            "IdentifiedAndItemQuality": 0,
            "DurabilityCur": 0,
            "DurabilityMax": 0,
            "ChargesMin": 0,
            "ChargesMax": 0,
            "GoldPrice": 100,
            "OnlyUsedByEar": 0
        }
    ],
    "InvNumFromInvGrid": [
        -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0
    ],
    "NInvItems": 2,
    "BeltItems": [
        {
            "Seed": 1262929936,
            "CF": 0,
            "ItemID": "Potion of Healing",
            "IdentifiedAndItemQuality": 0,
            "DurabilityCur": 0,
            "DurabilityMax": 0,
            "ChargesMin": 0,
            "ChargesMax": 0,
            "GoldPrice": 0,
            "OnlyUsedByEar": 0
        },
        {
            "Seed": 1942205617,
            "CF": 0,
            "ItemID": "Potion of Healing",
            "IdentifiedAndItemQuality": 0,
            "DurabilityCur": 0,
            "DurabilityMax": 0,
            "ChargesMin": 0,
            "ChargesMax": 0,
            "GoldPrice": 0,
            "OnlyUsedByEar": 0
        }
    ],
    "OnBattlenet": 0,
    "HasManashild": 0,
    "Difficulty": 0
}
```
