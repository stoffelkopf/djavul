// Package parse provides access to the data of the diablo.exe exectuable.
package parse

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

// Data reads the data at the specified offset.
func Data(offset int64, data interface{}) error {
	r := bytes.NewReader(buf)
	if _, err := r.Seek(offset, io.SeekStart); err != nil {
		return errors.WithStack(err)
	}
	if err := binary.Read(r, binary.LittleEndian, data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DataFromAddr reads the data at the specified address.
func DataFromAddr(addr uint32, data interface{}) error {
	return Data(Offset(addr), data)
}

// MustDataFromAddr reads the data at the specified address. It panics on error.
func MustDataFromAddr(addr uint32, data interface{}) {
	if err := DataFromAddr(addr, data); err != nil {
		panic(err)
	}
}

// Offset returns the file offset of the given address in diablo.exe.
func Offset(addr uint32) int64 {
	switch {
	case addr >= 0x401000 && addr < 0x479000:
		// .text
		//   start: 0x401000
		//   end:   0x479000
		//
		// file offset: 0x00000400
		return 0x00000400 + int64(addr) - 0x401000
	case addr >= 0x479000 && addr < 0x483000:
		// .rdata
		//   start: 0x479000
		//   end:   0x483000
		//
		// file offset: 0x00077A00
		return 0x00077A00 + int64(addr) - 0x479000
	case addr >= 0x483000 && addr < 0x6AE000:
		// .data
		//   start: 0x483000
		//   end:   0x6AE000
		//
		// file offset: 0x00080E00
		return 0x00080E00 + int64(addr) - 0x483000
	case addr >= 0x6AE000 && addr < 0x6B2000:
		// .rsrc
		//   start: 0x6AE000
		//   end:   0x6B2000
		//
		// file offset: 0x000B5800
		return 0x000B5800 + int64(addr) - 0x6AE000
	default:
		panic(fmt.Errorf("unknown segment of address 0x%08X", addr))
	}
}

// buf contains the contents of diablo.exe.
var buf []byte

func init() {
	// Read file contents of diablo.exe.
	var err error
	buf, err = ioutil.ReadFile("diablo.exe")
	if err != nil {
		log.Fatalf("unable to read file; %v", errors.WithStack(err))
	}
	h := sha1.Sum(buf)
	// The size in bytes of diablo.exe v1.09 and v1.09b are identical.
	// Furthermore the offset to each function and global variable is identical,
	// thus we can effortlessly support both versions.
	switch hash := fmt.Sprintf("%040x", h[:]); hash {
	case "accedfe32775d4a1984451309608c2a2d39ad406":
		// v1.00dbg, not supported.
		log.Fatalf("support for diablo.exe v1.00dbg not yet implemented (expected version 1.09 or 1.09b)")
	case "9633d9bbf9a2dc3e88571cead6ac09d689c2abcf":
		// v1.00, not supported.
		log.Fatalf("support for diablo.exe v1.00 not yet implemented (expected version 1.09 or 1.09b)")
	case "958e3503839a0798d544868d5fc3dff05c02e9fa":
		// v1.04, not supported.
		log.Fatalf("support for diablo.exe v1.04 not yet implemented (expected version 1.09 or 1.09b)")
	case "d8af136407ea1c019c881f31a180a0bceda39226":
		// v1.07, not supported.
		log.Fatalf("support for diablo.exe v1.07 not yet implemented (expected version 1.09 or 1.09b)")
	case "c2c111a1825c0410eec93f2f0dc872dd49f8c0db":
		// v1.08, not supported.
		log.Fatalf("support for diablo.exe v1.08 not yet implemented (expected version 1.09 or 1.09b)")
	case "2119e1c8b818c27a06948979560cdeb4bec9ae65":
		// v1.09, supported.
	case "ebaee2acb462a0ae9c895a0e33079c94796cb0b6":
		// v1.09b, supported.
	case "e59538ac8de87063e5d3e921a0c5d629e8d54c4e":
		// v1.09b (no CD), supported.
	default:
		log.Fatalf("support for unknown version of diablo.exe with SHA1 hashsum %q not yet implemented (expected version 1.09 or 1.09b)", hash)
	}
}
