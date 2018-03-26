package sha1_test

import (
	"bytes"
	"hash"
	"testing"

	"github.com/sanctuary/djavul/d1/sha1"
)

func TestSum(t *testing.T) {
	golden := []struct {
		input string
		want  [sha1.Size]byte
	}{
		{
			input: "xrgyrkj1xrgyrkj1xrgyrkj1xrgyrkj1xrgyrkj1xrgyrkj1xrgyrkj1xrgyrkj1",
			want:  [sha1.Size]byte{0x7A, 0xDC, 0xF4, 0x4C, 0xDC, 0x67, 0xC7, 0xBB, 0x56, 0xBE, 0x05, 0x81, 0x8C, 0x9E, 0x18, 0x0C, 0x3C, 0x3E, 0x7B, 0xA4},
		},
	}
	for _, g := range golden {
		got := sha1.Sum([]byte(g.input))
		if !bytes.Equal(got, g.want[:]) {
			t.Errorf("SHA-1 hashsum mismatch for %q; expected %x, got %x", g.input, g.want, got)
		}
	}
}

// Ensure that Context implements the hash.Hash interface.
var _ hash.Hash = (*sha1.Context)(nil)
