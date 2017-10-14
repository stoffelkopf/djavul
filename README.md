# djavul

Experimenting with C-shared libraries in Go.

## Installation

Convert `diablo.exe` version 1.09b to NASM using [bin2asm](https://godoc.org/github.com/decomp/exp/cmd/bin2asm) and [dump_imports](https://godoc.org/github.com/decomp/exp/cmd/dump_imports).

Edit the import data in `_rdata.asm` to include the new DLL files. To keep addresses fixed, replace `shell32.dll` with `aaaaaaa.dll`; e.g.

```diff
diff --git a/_rdata.asm b/_rdata.asm
index 96ed9bf..059bb03 100644
--- a/_rdata.asm
+++ b/_rdata.asm
@@ -319,15 +319,15 @@ iat_kernel32:
                         dd      imp_GetTimeZoneInformation - IMAGE_BASE
                         dd      0x00000000

-; --- [ shell32.dll ] ---------------------------------------------------------
-
-iat_shell32:
-  ia_ShellExecuteA:
-                        dd      imp_ShellExecuteA - IMAGE_BASE
-  ia_SHGetSpecialFolderLocation:
-                        dd      imp_SHGetSpecialFolderLocation - IMAGE_BASE
-  ia_SHGetPathFromIDListA:
-                        dd      imp_SHGetPathFromIDListA - IMAGE_BASE
+; --- [ aaaaaaa.dll ] ---------------------------------------------------------
+
+iat_aaaaaaa:
+  ia_Bbbbbbbbbbbbb:
+                        dd      imp_Bbbbbbbbbbbbb - IMAGE_BASE
+  ia_Cccccccccccccccccccccccccc:
+                        dd      imp_Cccccccccccccccccccccccccc - IMAGE_BASE
+  ia_Dddddddddddddddddddd:
+                        dd      imp_Dddddddddddddddddddd - IMAGE_BASE
                         dd      0x00000000

 ; --- [ storm.dll ] ---------------------------------------------------------
@@ -32542,11 +32542,11 @@ import_table:
                         dd      szAdvapi32_dll - IMAGE_BASE
                         dd      iat_advapi32 - IMAGE_BASE      ; pIAT_first_trunk

-                        dd      int_shell32 - IMAGE_BASE      ; pINT_first_trunk
+                        dd      int_aaaaaaa - IMAGE_BASE      ; pINT_first_trunk
                         dd      0x00000000                         ; TimeDateStamp
                         dd      0x00000000                         ; pForwardChain
-                        dd      szShell32_dll - IMAGE_BASE
-                        dd      iat_shell32 - IMAGE_BASE      ; pIAT_first_trunk
+                        dd      szAaaaaaa_dll - IMAGE_BASE
+                        dd      iat_aaaaaaa - IMAGE_BASE      ; pIAT_first_trunk

                         dd      int_version - IMAGE_BASE      ; pINT_first_trunk
                         dd      0x00000000                         ; TimeDateStamp
@@ -32732,12 +32732,12 @@ int_kernel32:
                         dd      imp_GetTimeZoneInformation - IMAGE_BASE
                         dd      0x00000000

-; --- [ shell32.dll ] ---------------------------------------------------------
+; --- [ aaaaaaa.dll ] ---------------------------------------------------------

-int_shell32:
-                        dd      imp_ShellExecuteA - IMAGE_BASE
-                        dd      imp_SHGetSpecialFolderLocation - IMAGE_BASE
-                        dd      imp_SHGetPathFromIDListA - IMAGE_BASE
+int_aaaaaaa:
+                        dd      imp_Bbbbbbbbbbbbb - IMAGE_BASE
+                        dd      imp_Cccccccccccccccccccccccccc - IMAGE_BASE
+                        dd      imp_Dddddddddddddddddddd - IMAGE_BASE
                         dd      0x00000000

 ; --- [ storm.dll ] ---------------------------------------------------------
@@ -33826,25 +33826,25 @@ szAdvapi32_dll:
                         db      'ADVAPI32.dll', 0x00 ; 0x0008228C
                         align 2, db 0x00

-; --- [ SHELL32.dll ] ---------------------------------------------------------
+; --- [ aaaaaaa.dll ] ---------------------------------------------------------

-imp_SHGetPathFromIDListA:
-                        dw      0x0063
-                        db      'SHGetPathFromIDListA', 0x00 ; 0x0008229A
+imp_Dddddddddddddddddddd:
+                        dw      0x0000
+                        db      'Dddddddddddddddddddd', 0x00 ; 0x0008229A
                         align 2, db 0x00

-imp_SHGetSpecialFolderLocation:
-                        dw      0x0066
-                        db      'SHGetSpecialFolderLocation', 0x00 ; 0x000822B2
+imp_Cccccccccccccccccccccccccc:
+                        dw      0x0001
+                        db      'Cccccccccccccccccccccccccc', 0x00 ; 0x000822B2
                         align 2, db 0x00

-imp_ShellExecuteA:
-                        dw      0x008C
-                        db      'ShellExecuteA', 0x00 ; 0x000822D0
+imp_Bbbbbbbbbbbbb:
+                        dw      0x0002
+                        db      'Bbbbbbbbbbbbb', 0x00 ; 0x000822D0
                         align 2, db 0x00

-szShell32_dll:
-                        db      'SHELL32.dll', 0x00 ; 0x000822E0
+szAaaaaaa_dll:
+                        db      'aaaaaaa.dll', 0x00 ; 0x000822E0
                         align 2, db 0x00

 ; --- [ VERSION.dll ] ---------------------------------------------------------
```

Edit the assembly in `_text.asm` to remove the call to `init_run_office_from_start_menu` (i.e. `0x41A84C`); e.g.

```diff
diff --git a/_text.asm b/_text.asm
index 3bbd269..9b4e3d6 100644
--- a/_text.asm
+++ b/_text.asm
@@ -44795,7 +44795,7 @@ sub_41A7C3:
   addr_41A7C6:          db      0xE8, 0x57, 0xF5, 0x02, 0x00                    ; CALL .+193879
   addr_41A7CB:          db      0x32, 0xC9                                      ; XOR CL, CL
   addr_41A7CD:          db      0xE8, 0x5A, 0x02, 0x00, 0x00                    ; CALL .+602
-  addr_41A7D2:          db      0xE8, 0x75, 0x00, 0x00, 0x00                    ; CALL .+117
+times (0x41A7D7 - _text_vstart) - ($ - $$) nop
   addr_41A7D7:          db      0xA1, 0x94, 0x4B, 0x63, 0x00                    ; MOV EAX, [+0x634b94]
   addr_41A7DC:          db      0x85, 0xC0                                      ; TEST EAX, EAX
   addr_41A7DE:          db      0x74, 0x0D
```

### Building shared libraries

Install cross compiler for Windows.

```bash
pacman -Sy gcc-multilib mingw-w64-gcc
```

Build shared library.

```bash
GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -buildmode=c-shared -o aaaaaaa.dll github.com/sanctuary/djavul/aaaaaaa
```

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
