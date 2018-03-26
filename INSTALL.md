# Installation on Windows

**Note:** djavul requires an original copy of `diablo.exe` and `diabdat.mpq`. None of the Diablo 1 game assets are provided by this project.

These installation instructions assumes that Diablo 1 (version 1.09b) has been installed into a game directory (from now on referred to as `$GAMEDIR`) and that `diabdat.mpq` has been copied from the installation CD to the game directory.

The game directory should contain the following files, prior to installation of Djavul.

```
diabdat.mpq    (5cfd971abb25602731fef0c9b43eb7d7447f296e)
diablo.exe     (ebaee2acb462a0ae9c895a0e33079c94796cb0b6)
diabloui.dll   (281c720eab871df25760bf92b67d5bb40025f7e8)
standard.snp   (dc183027b11114fbc8bcb35e1b7befc75df077e6)
storm.dll      (a8d7f56fd81976e98e7dadf0efae4625f6749b84)
```

Note, `smackw32.dll` may optionally be added for video playback of SMK files. For reference the SHA1 hash sum of `smackw32.dll` is `d85bc8a2aec0eafa35f59baa1cc9293661c7b618`.

## Install development tools

* Install [Go](https://golang.org/doc/install)
* Install [MSYS2](https://www.msys2.org/)

## Configure environment variables

Open *Edit environment variables for your account* from the start menu.

* Configure `GOPATH` (e.g. set to `$HOME/go`).
* Add `$GOPATH/bin` to the semicolon separated list `PATH`.

Add the following lines to `$HOME/.bash_profile` of MSYS.

```bash
export GOPATH=$HOME/go
export PATH=$PATH:/c/Go/bin
export PATH=$PATH:$GOPATH/bin
```

## Install dependencies

```bash
# Start the *MSYS2 MSYS* terminal.
pacman -S git make nasm mingw-w64-i686-gcc mingw-w64-x86_64-gcc
```

Install [Microsoft .NET Framework Version 1.1 Redistributable Package](https://www.microsoft.com/en-us/download/details.aspx?id=26) and copy `msvcr71.dll` to the game directory.

```bash
cp $WINDIR/Microsoft.NET/Framework/v1.1.4322/msvcr71.dll $GAMEDIR/msvcr71.dll
```

## Install the Djavul frontend

Build `djavul-frontend.exe`:

```bash
# Run from the *MSYS2 MinGW 64-bit* terminal.
go get -v -u github.com/sanctuary/djavul/...
```

## Install the Djavul backend

Build `djavul.dll`:

```bash
# Run from the *MSYS2 MinGW 32-bit* terminal.
go get -v -u -tags djavul github.com/sanctuary/djavul/...
make -C $GOPATH/src/github.com/sanctuary/djavul/dll
ln -s $GOPATH/src/github.com/sanctuary/djavul/dll/crt0.dll $GAMEDIR/crt0.dll
ln -s $GOPATH/src/github.com/sanctuary/djavul/dll/djavul.dll $GAMEDIR/djavul.dll
```

Build `djavul.exe` by following the installation instructions on https://github.com/sanctuary/djavul-patch and place it in the game directory.

## Extract game assets

```bash
# Extract the diabdat.mpq archive.
go get github.com/sanctuary/mpq
cd $GAMEDIR
mpq -dir diabdat -m diabdat.mpq
```

The game directory should contain the following files and directories, after to installation of Djavul.

```diff
+diabdat/
+crt0.dll
 diabdat.mpq
 diablo.exe
 diabloui.dll
+djavul.dll
+djavul.exe
+msvcr71.dll
 standard.snp
 storm.dll
```

# Running on Windows

To start Djavul, first start the frontend as it will listen for incoming connections from the backend.

## Running the Djavul frontend

```bash
# Run from the *MSYS2 MinGW 64-bit* terminal.
cd $GAMEDIR
djavul-frontend.exe
```

## Running the Djavul backend

```bash
# Run from the cmd.exe terminal.
cd $GAMEDIR
djavul.exe -ip localhost
```

# Optional: Multiplayer over UDP

To enable Local Area Network over UDP, refer to the thread [The death of IPX protocol [IPX + UDP Fix]](http://www.lurkerlounge.com/forums/thread-353.html) at Lurker Lounge. In essence, download the `UDP_Diablo&Hellfire.zip` archive, extract it, and replace `battle.snp` with `sndi.nx`. Make sure the rename `sndi.nx` to `sndi.snp` for `storm.dll` to be able to locate it. For reference, the SHA1 hash sum of `sndi.nx` is `f08c935b60fb6b844dcc4690b1abbc8a7a4e1eda`.
