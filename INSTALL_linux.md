## Installation

1. Install [Go](https://golang.org/doc/install).
2. Install [mingw-w64-gcc](https://aur.archlinux.org/packages/mingw-w64-gcc/) for cross-compilation to 32-bit Windows.

Install the [Djavul frontend](https://github.com/sanctuary/djavul/tree/master/cmd/djavul) executable `djavul-frontend.exe`.

```bash
go get -u github.com/sanctuary/djavul/...
```

Install the [Djavul backend](https://github.com/sanctuary/djavul/tree/master/dll/djavul) library `djavul.dll`.

```bash
$ cd $GOPATH/src/github.com/sanctuary
$ make -C dll/djavul
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
# Extract the diabdat.mpq archive.
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
# Specify the IP-address of the system running the Djavul frontend
# (default is localhost).
wine djavul.exe -ip 192.168.1.100
```
