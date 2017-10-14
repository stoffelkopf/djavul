# djavul

Experimenting with C-shared libraries in Go.

## Installation

Convert `diablo.exe` version 1.09b to NASM using [bin2asm](https://godoc.org/github.com/decomp/exp/cmd/bin2asm) and [dump_imports](https://godoc.org/github.com/decomp/exp/cmd/dump_imports).

Edit the import data in `_rdata.asm` to include the new DLL files; e.g.

```diff
diff --git a/_rdata.asm b/_rdata.asm
index 96ed9bf..22e69df 100644
--- a/_rdata.asm
+++ b/_rdata.asm
@@ -553,6 +553,13 @@ iat_version:
                         dd      imp_VerQueryValueA - IMAGE_BASE
                         dd      0x00000000
 
+; --- [ foo.dll ] --------------------------------------------------------------
+
+iat_foo:
+  ia_Foo:
+                        dd      imp_Foo - IMAGE_BASE
+                        dd      0x00000000
+
    iat_size             equ     $ - iat
 
 ; === [/ Import Address Tables (IATs) ] ========================================
@@ -32554,6 +32561,12 @@ import_table:
                         dd      szVersion_dll - IMAGE_BASE
                         dd      iat_version - IMAGE_BASE      ; pIAT_first_trunk
 
+                        dd      int_foo - IMAGE_BASE      ; pINT_first_trunk
+                        dd      0x00000000                         ; TimeDateStamp
+                        dd      0x00000000                         ; pForwardChain
+                        dd      szFoo_dll - IMAGE_BASE
+                        dd      iat_foo - IMAGE_BASE      ; pIAT_first_trunk
+
    times 5              dd      0x00000000
 
    import_table_size    equ     $ - import_table
@@ -32859,6 +32872,12 @@ int_version:
                         dd      imp_VerQueryValueA - IMAGE_BASE
                         dd      0x00000000
 
+; --- [ foo.dll ] ---------------------------------------------------------
+
+int_foo:
+                        dd      imp_Foo - IMAGE_BASE
+                        dd      0x00000000
+
 ; === [/ Import Name Tables (INTs) ] ===========================================
 
    int_size             equ     $ - int
@@ -33868,6 +33887,17 @@ szVersion_dll:
                         db      'VERSION.dll', 0x00 ; 0x0008232E
                         align 2, db 0x00
 
+; --- [ foo.dll ] --------------------------------------------------------------
+
+imp_Foo:
+                        dw      0x0000
+                        db      'Foo', 0x00
+                        align 2, db 0x00
+
+szFoo_dll:
+                        db      'foo.dll', 0x00
+                        align 2, db 0x00
+
 ; === [/ dll and function names ] ==============================================
 
    _rdata_vsize         equ     $ - $$
```

Edit the assembly in `_text.asm` to call the new functions; e.g.

```diff
diff --git a/_text.asm b/_text.asm
index da8ac48..c3e2a57 100644
--- a/_text.asm
+++ b/_text.asm
@@ -12183,8 +12183,10 @@ sub_408B4A:
   addr_408B53:          db      0x53                                            ; PUSH EBX
   addr_408B54:          db      0x56                                            ; PUSH ESI
   addr_408B55:          db      0x8B, 0x75, 0x08                                ; MOV ESI, [EBP+0x8]
-  addr_408B58:          db      0x8B, 0xCE                                      ; MOV ECX, ESI
-  addr_408B5A:          db      0xE8, 0x95, 0x02, 0x00, 0x00                    ; CALL .+661
+  call DWORD [ia_Foo]
+times (0x408B5F - _text_vstart) - ($ - $$) db 0x90
+  ;addr_408B58:          db      0x8B, 0xCE                                      ; MOV ECX, ESI
+  ;addr_408B5A:          db      0xE8, 0x95, 0x02, 0x00, 0x00                    ; CALL .+661
   addr_408B5F:          db      0x89, 0x35, 0xEC, 0x56, 0x52, 0x00              ; MOV [+0x5256ec], ESI
   addr_408B65:          db      0xE8, 0x1B, 0x9D, 0x04, 0x00                    ; CALL .+302363
   addr_408B6A:          db      0x85, 0xC0                                      ; TEST EAX, EAX
```

### Building shared libraries

Install cross compiler for Windows.

```bash
pacman -Sy gcc-multilib mingw-w64-gcc
```

Build shared library.

```bash
GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -buildmode=c-shared -o foo.dll import/path/to/foo
```

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
