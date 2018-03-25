# djavul

Preservation project for an all time classic, Diablo 1.

**Note**, djavul requires an original copy of `diablo.exe` and `diabdat.mpq`. None of the Diablo 1 game assets are provided by this project.

**Project aim and objectives**

The aim of this project is to provide an open source reference re-implementation of the Diablo 1 game engine.

To achieve this aim, the following objectives have been identified.

1. Develop an understanding of the inner workings of the Diablo 1 game engine (subproject [notes](https://github.com/sanctuary/notes)).
2. Convert the original game assets to file formats with open specifications (subproject [formats](https://github.com/sanctuary/formats)).
3. Provide a framework of extensive test cases for comparing the re-implementation against the original Diablo 1 game engine.
4. Split the engine into self-contained modules that may be validated and verified independently.
5. Implement a set of tools around these modules which through interaction provide the functionality of the original Diablo 1 game engine.
6. Validate that - given a deterministic seed - the re-implementation achieve pixel perfection, by mirroring the mouse and keyboard input and comparing the graphic and audio output against the original Diablo 1 game.

## Installation

1. Install [Go](https://golang.org/doc/install).
2. Install [mingw-w64-gcc](https://aur.archlinux.org/packages/mingw-w64-gcc/) for cross-compilation to 32-bit Windows.

Install the [Djavul frontend](https://github.com/sanctuary/djavul/tree/master/cmd/djavul) executable `djavul-frontend.exe`.

```bash
go get -u github.com/sanctuary/djavul/...
```

Install the [Djavul backend](https://github.com/sanctuary/djavul/tree/master/dlls/djavul) library `djavul.dll`.

```bash
$ cd $GOPATH/src/github.com/sanctuary
$ make -C dlls/djavul
```

Install the Djavul backend executable `djavul.exe` by applying the patch file [djavul.patch](https://github.com/sanctuary/djavul-patch) to `diablo.exe` (v1.09b).

```bash
# Download the patch file djavul.patch.
wget https://github.com/sanctuary/djavul-patch/raw/master/djavul.patch
# Install binpatch.
go get github.com/mewkiz/cmd/binpatch
# Apply djavul.patch to diablo.exe to produce djavul.exe.
binpatch -o djavul.exe diablo.exe < djavul.patch
```

Place the `djavul-frontend.exe` binary, the `djavul.dll` library and the `djavul.exe` binary in the game directory (the one containing `diablo.exe`).

Copy `diabdat.mpq` to the game directory and extract its contents to a `diabdat` subdirectory of the game directory.

```bash
# Extract DIABDAT.MPQ archive.
go get github.com/sanctuary/mpq
mpq -dir diabdat -m diabdat.mpq
```

## Run

Start the Djavul frontend.

```bash
# Execute from within the game directory.
djavul-frontend
```

Start the Djavul backend.

```bash
# Execute from within the game directory.
# Specify the IP-address of the system running the Djavul frontend.
wine djavul.exe -ip 192.168.1.100
```

## Progress

### 2018-03-25

* Running the Djavul frontend (`djavul-frontend`) on a Linux host system (left) and the Djavul backend (`djavul.exe`) on a Windows XP VirtualBox guest (right).

![Screenshot from 2018-03-25](https://github.com/sanctuary/graphics/blob/master/djavul/screenshot_2018-03-25.png)

### 2017-12-23

* Decode legacy Diablo 1 save files.
    - [cmd/sv](cmd/sv)

### 2017-12-22

* Play sounds using Beep.

![Screenshot from 2017-12-22](https://github.com/sanctuary/graphics/blob/master/djavul/screenshot_2017-12-22.png)

### 2017-12-21

* Forward mouse and keyboard input from Djavul frontend to Diablo 1 game engine.

![Screenshot from 2017-12-21](https://github.com/sanctuary/graphics/blob/master/djavul/screenshot_2017-12-21.png)

### 2017-12-16

* Render control panel.

![Screenshot from 2017-12-16](https://github.com/sanctuary/graphics/blob/master/djavul/screenshot_2017-12-16.png)

### 2017-11-24

* Correct rendering of Tristram.

![Screenshot from 2017-11-24](https://github.com/sanctuary/graphics/blob/master/djavul/screenshot_2017-11-24.png)

### 2017-11-22

* Mirror output using Pixel.

![Screenshot from 2017-11-22](https://github.com/sanctuary/graphics/blob/master/djavul/screenshot_2017-11-22.png)

### 2017-11-13

* Dungeon generation of Cathedral.
    - [cmd/l1](cmd/l1)

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
