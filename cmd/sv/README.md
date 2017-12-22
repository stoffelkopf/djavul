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
$ sv -p szqnlsk1 hero | hexdump -C

00000000  00 00 00 00 00 00 00 00  ff 00 00 00 49 41 49 41  |............IAIA|
00000010  66 6f 6f 62 61 72 00 00  00 00 00 00 00 00 00 00  |foobar..........|
...
```
