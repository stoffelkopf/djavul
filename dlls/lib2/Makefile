all: lib2.dll

lib2.dll:
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -buildmode=c-shared -o $@ .
