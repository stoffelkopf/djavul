all: hook_aa.dll bbbbbbbb.dll

hook_aa.dll:
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -buildmode=c-shared -o hook_aa.dll github.com/sanctuary/djavul/dlls/hook_aa

bbbbbbbb.dll:
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -buildmode=c-shared -o bbbbbbbb.dll github.com/sanctuary/djavul/dlls/bbbbbbbb
