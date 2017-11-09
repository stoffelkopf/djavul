# l1

The l1 tool provides dynamic random level generation of Cathedral maps.

## Installation

```bash
go get github.com/sanctuary/djavul/cmd/l1
```

### Game assets

**Note**, `l1` requires an original copy of `diablo.exe` and `diabdat.mpq`. None of the Diablo 1 game assets are provided by this project.

```bash
# Extract diabdat.mpq to the "diabdat" directory.
go get github.com/sanctuary/mpq
mpq -dir diabdat -m /path/to/diabdat.mpq

# Copy diablo.exe version 1.09b to the current directory.
cp /path/to/diablo.exe .
```

## Usage

```bash
# Generate the Cathedral map of dungeon level 2 with seed 36 and The Butcher quest active.
l1 -dlvl 2 -seed 36 -quest "The Butcher"
```

![The Butcher, seed 36](https://github.com/sanctuary/graphics/blob/master/l1/maps/the_butcher_seed_36.png?raw=true)

```bash
# Generate the Cathedral map of dungeon level 2 with seed 1 and Poisoned Water Supply quest active.
l1 -dlvl 2 -seed 1 -quest "Poisoned Water Supply"
```

![Poisoned Water Supply, seed 1](https://github.com/sanctuary/graphics/blob/master/l1/maps/poisoned_water_supply_seed_1.png?raw=true)

```bash
# Generate the Cathedral map of dungeon level 4 with seed 215 and Ogden's Sign quest active.
l1 -dlvl 4 -seed 215 -quest "Ogden's Sign"
```

![Ogden's Sign, seed 215](https://github.com/sanctuary/graphics/blob/master/l1/maps/ogdens_sign_seed_215.png?raw=true)

```bash
# Generate the Cathedral map of dungeon level 4 with seed 1000.
l1 -dlvl 4 -seed 1000
```

![Dungeon level 4, seed 1000](https://github.com/sanctuary/graphics/blob/master/l1/maps/dlvl_4_seed_1000.png)
