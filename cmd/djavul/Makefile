all: djavul-frontend.exe

djavul-frontend.exe:
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -o $@ .

.PHONY: all clean

clean:
	rm -f djavul-frontend.exe
