all: crt0.dll

crt0.dll:
	i686-w64-mingw32-gcc -shared -o $@ *.c -Wl,--out-implib,libcrt0_dll.a

.PHONY: all clean

clean:
	rm -f crt0.dll libcrt0_dll.a
