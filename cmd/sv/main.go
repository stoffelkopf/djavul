// The sv tool decodes Diablo 1 save files.
//
// Usage:
//
//    sv [OPTION]... FILE...
//
// Flags:
//
//    -p string
//          password (multi: "szqnlsk1" or computer name) (default "xrgyrkj1")
package main

import (
	"flag"
	"fmt"
	"hash"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/sha1"
)

func usage() {
	const use = `
Decode Diablo 1 save files.

Usage:

	sv [OPTION]... FILE...

Flags:
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line flags.
	var (
		// Output in JSON format.
		jsonOutput bool
		// Password for decoding save files.
		password string
	)
	flag.BoolVar(&jsonOutput, "json", false, "output in JSON format")
	flag.StringVar(&password, "p", "xrgyrkj1", `password (multi: "szqnlsk1" or computer name)`)
	flag.Usage = usage
	flag.Parse()
	for _, path := range flag.Args() {
		data, err := decodeFile(path, password)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		switch {
		case jsonOutput:
			// Output in JSON format.
			name := filepath.Base(path)
			switch name {
			case "hero":
				if err := outputHero(data); err != nil {
					log.Fatalf("%+v", err)
				}
			default:
				panic(fmt.Errorf("support for presenting %q in JSON format not yet implemented", name))
			}
		default:
			// Raw output.
			if _, err := os.Stdout.Write(data); err != nil {
				log.Fatalf("%+v", err)
			}
		}
	}
}

// decodeFile decodes the given save file.
func decodeFile(path, password string) ([]byte, error) {
	enc, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dec := decode(enc, password)
	return dec, nil
}

// decode decodes the given save file buffer.
func decode(enc []byte, password string) []byte {
	h := initKey(password)
	n := len(enc) - 8
	if n <= 0 {
		return nil
	}
	if n%sha1.BlockSize != 0 {
		return nil
	}
	block := make([]byte, sha1.BlockSize)
	dst := make([]byte, n)
	for i := 0; i < n; i += sha1.BlockSize {
		copy(block, enc[i:i+sha1.BlockSize])
		digest := h.Sum(nil)
		for j := range block {
			block[j] ^= digest[j%sha1.Size]
		}
		h.Write(block)
		copy(dst[i:], block)
	}
	return dst
}

// initKey returns a running hash of the codec key, as used for decoding save
// files.
//
// ref: 0x4035DB
func initKey(password string) hash.Hash {
	setSeed(28760)
	rnd := make([]byte, 136)
	for i := range rnd {
		rnd[i] = byte(rand())
	}
	block := make([]byte, sha1.BlockSize)
	pwd := []byte(password)
	for i := range block {
		block[i] = pwd[i%len(pwd)]
	}
	digest := sha1.Sum(block)
	for i := range rnd {
		rnd[i] ^= digest[i%len(digest)]
	}
	h := sha1.New()
	h.Write(rnd[72:])
	return h
}
