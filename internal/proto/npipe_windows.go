package proto

import (
	"net"

	"github.com/natefinch/npipe"
)

// Dial connects to the remote host over a named pipe.
func (i *NamedPipe) Dial() (net.Conn, error) {
	return npipe.Dial(i.addr)
}

// Listen listens for incoming connections on a named pipe.
func (i *NamedPipe) Listen() (net.Listener, error) {
	return npipe.Listen(i.addr)
}
