all: hook.dll lib2.dll

hook.dll:
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -buildmode=c-shared -o hook.dll github.com/sanctuary/djavul/dlls/hook

lib2.dll:
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -buildmode=c-shared -o lib2.dll github.com/sanctuary/djavul/dlls/lib2
