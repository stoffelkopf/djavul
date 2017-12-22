// NOTE: The SHA-1 implementation of Diablo 1 is not standard compliant.
//
//    ref: https://blog.skullsecurity.org/2012/battle-net-authentication-misconceptions

// Package sha1 implements a non-standard compliant SHA-1 hash algorithm, as
// used by Diablo 1.
package sha1

import (
	"encoding/binary"
)

const (
	// The blocksize of SHA-1 in bytes.
	BlockSize = 64
	// The size of a SHA-1 checksum in bytes.
	Size = 20
)

// Context holds the context of a SHA-1 running hashsum. Context implements the
// hash.Hash interface.
type Context struct {
	State  [5]uint32
	Count  [2]int32
	Buffer [BlockSize]uint8
}

// New returns a new initialized SHA-1 context.
func New() *Context {
	ctx := &Context{}
	ctx.Init()
	return ctx
}

// Final copies the message digest of the given SHA-1 context to dst.
//
// ref: 0x456A2B
func (ctx *Context) Final(digest []byte) {
	binary.LittleEndian.PutUint32(digest[0:], ctx.State[0])
	binary.LittleEndian.PutUint32(digest[4:], ctx.State[1])
	binary.LittleEndian.PutUint32(digest[8:], ctx.State[2])
	binary.LittleEndian.PutUint32(digest[12:], ctx.State[3])
	binary.LittleEndian.PutUint32(digest[16:], ctx.State[4])
}

// Update adds the data to the running hash of the given SHA-1 context.
//
// ref: 0x456A73
func (ctx *Context) Update(data []byte) {
	count0 := ctx.Count[0]
	n := int32(len(data))
	bits := count0 + 8*n
	if bits < count0 {
		ctx.Count[1]++
	}
	ctx.Count[0] = bits
	ctx.Count[1] += n >> 29
	for i := int32(0); i <= n-BlockSize; i += BlockSize {
		copy(ctx.Buffer[:], data[i:i+BlockSize])
		ctx.Transform()
	}
}

// Transform performs a SHA-1 transformation on the 64-byte block of the given
// SHA-1 context.
//
// ref: 0x456AC4
func (ctx *Context) Transform() {
	var buf [80]uint32
	for i := 0; i < len(ctx.Buffer); i += 4 {
		buf[i/4] = binary.LittleEndian.Uint32(ctx.Buffer[i:])
	}
	for i := 0; i < BlockSize; i++ {
		buf[i+16] = buf[i+0] ^ buf[i+2] ^ buf[i+8] ^ buf[i+13]
	}
	a := ctx.State[0]
	b := ctx.State[1]
	c := ctx.State[2]
	d := ctx.State[3]
	e := ctx.State[4]
	g := uint32(0)
	for _, v := range buf[0:20] {
		g = v + rol(a, 5) + e + (b&c | ^b&d) + 0x5A827999
		e = d
		d = c
		c = rol(b, 30)
		b = a
		a = g
	}
	for _, v := range buf[20:40] {
		g = v + rol(g, 5) + e + (b ^ c ^ d) + 0x6ED9EBA1
		e = d
		d = c
		c = rol(b, 30)
		b = a
		a = g
	}
	for _, v := range buf[40:60] {
		g = v + rol(g, 5) + e + (b&c | b&d | c&d) - 0x70E44324
		e = d
		d = c
		c = rol(b, 30)
		b = a
		a = g
	}
	for _, v := range buf[60:80] {
		g = v + rol(g, 5) + e + (b ^ c ^ d) - 0x359D3E2A
		e = d
		d = c
		c = rol(b, 30)
		b = a
		a = g
	}
	ctx.State[0] += g
	ctx.State[1] += b
	ctx.State[2] += c
	ctx.State[3] += d
	ctx.State[4] += e
}

// Init initializes the given SHA-1 context.
//
// ref: 0x456C82
func (ctx *Context) Init() {
	ctx.Count[0] = 0
	ctx.Count[1] = 0
	ctx.State[0] = 0x67452301
	ctx.State[1] = 0xEFCDAB89
	ctx.State[2] = 0x98BADCFE
	ctx.State[3] = 0x10325476
	ctx.State[4] = 0xC3D2E1F0
}

// Size returns the size of a SHA-1 checksum in bytes.
func (ctx *Context) Size() int {
	return Size
}

// BlockSize returns the blocksize of SHA-1 in bytes.
func (ctx *Context) BlockSize() int {
	return BlockSize
}

// Reset resets the Hash to its initial state.
func (ctx *Context) Reset() {
	ctx.Init()
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (ctx *Context) Sum(b []byte) []byte {
	ctx.Write(b)
	digest := make([]byte, Size)
	ctx.Final(digest)
	return digest
}

// Write writes data to the running SHA-1 hashsum. It never returns an error.
func (ctx *Context) Write(data []byte) (n int, err error) {
	ctx.Update(data)
	return len(data), nil
}

// Sum returns the SHA-1 checksum of the data.
func Sum(data []byte) []byte {
	ctx := New()
	return ctx.Sum(data)
}

// ### [ Helper functions ] ####################################################

// rol rotates x left shift bits.
func rol(x, shift uint32) uint32 {
	v := int32(x)
	return uint32((v >> (32 - shift)) | (v << shift))
}
