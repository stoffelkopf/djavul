# djavul

[![Join the chat at https://gitter.im/sanctuary/notes](https://badges.gitter.im/sanctuary/notes.svg)](https://gitter.im/sanctuary/notes)
[![GoDoc](https://godoc.org/github.com/sanctuary/djavul?status.svg)](https://godoc.org/github.com/sanctuary/djavul)

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

## Design

For an overview of the key idea behind this project, refer to the [design documentation](DESIGN.md).

## Installation

For installation and run instructions, refer to the respective documentation:

* [Linux installation](INSTALL_linux.md)
* [Windows installation](INSTALL_windows.md)

## Progress

### 2018-05-12

* Interact with the Diablo 1 game engine from Python script.

![Screenshot from 2018-05-12](https://github.com/sanctuary/graphics/blob/master/djavul/screenshot_2018-05-12.png)

### 2018-03-30

* Running the Djavul frontend (`djavul-frontend.exe`) on a Windows 7 host system (left) and the Djavul backend (`djavul.exe`) on the same system using a modified version of [DiabloPatch](http://diablopat.ch/) (right).

![Screenshot from 2018-03-30](https://github.com/sanctuary/graphics/blob/master/djavul/screenshot_2018-03-30.png)

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

* Forward mouse and keyboard input from Djavul frontend to Diablo 1 game engine (also referred to as the Djavul backend).

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
