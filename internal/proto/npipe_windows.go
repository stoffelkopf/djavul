package proto

import (
	"fmt"
	"net"

	"github.com/natefinch/npipe"
)

// Dial connects to the remote host over a named pipe.
func (i *NamedPipe) Dial() (net.Conn, error) {
	return npipe.Dial(i.Addr())
}

// Listen listens for incoming connections on a named pipe.
func (i *NamedPipe) Listen() (net.Listener, error) {
	return npipe.Listen(i.Addr())
}

// Addr returns the address of the named pipe.
func (i *NamedPipe) Addr() string {
	ip := i.ip
	if ip == "localhost" {
		ip = "."
	}
	return fmt.Sprintf(`\\%s\pipe\%s`, ip, i.pipe)
}
