package proto

import (
	"net"
)

// Dial connects to the remote host over a named pipe.
func (i *NamedPipe) Dial() (net.Conn, error) {
	panic("proto.NamedPipe.Dial: not yet implemented on Linux")
}

// Listen listens for incoming connections on a named pipe.
func (i *NamedPipe) Listen() (net.Listener, error) {
	panic("proto.NamedPipe.Dial: not yet implemented on Linux")
}
