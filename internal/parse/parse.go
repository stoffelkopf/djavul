// Package parse provides access to the data of the diablo.exe exectuable.
package parse

import (
	"bytes"
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
}
