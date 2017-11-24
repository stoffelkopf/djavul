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

## Progress

2017-11-22

* Correct rendering of Tristram.

![Screenshot from 2017-11-24](https://github.com/sanctuary/graphics/blob/master/djavul/screenshot_2017-11-24.png)

2017-11-22

* Mirror output using Pixel.

![Screenshot from 2017-11-22](https://github.com/sanctuary/graphics/blob/master/djavul/screenshot_2017-11-22.png)

2017-11-13

* Dungeon generation for Cathedral.
    - [cmd/l1](cmd/l1)

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
