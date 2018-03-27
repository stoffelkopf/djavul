package proto

import (
	"fmt"
	"net"
)

// IPC is an inter-process communication handler.
type IPC interface {
	// Dial connects to the remote host.
	Dial() (net.Conn, error)
	// Listen listens for incoming connections.
	Listen() (net.Listener, error)
	// Addr returns the address of the IPC connection.
	Addr() string
}

// Stable and unstable pipes.
const (
	stablePipe   = "stable"
	unstablePipe = "unstable"
)

// NamedPipe is a named pipe IPC handler.
type NamedPipe struct {
	ip   string
	pipe string
}

// NewNamedPipe returns a new named pipe IPC handler based on the IP of the
// remote host and pipe name.
func NewNamedPipe(ip, pipe string) *NamedPipe {
	return &NamedPipe{
		ip:   ip,
		pipe: pipe,
	}
}

// NewStableNamedPipe returns a new stable named pipe IPC handler based on the
// IP of the remote host.
func NewStableNamedPipe(ip string) *NamedPipe {
	return NewNamedPipe(ip, stablePipe)
}

// NewUnstableNamedPipe returns a new unstable named pipe IPC handler based on
// the IP of the remote host.
func NewUnstableNamedPipe(ip string) *NamedPipe {
	return NewNamedPipe(ip, unstablePipe)
}

// Stable and unstable ports.
const (
	stablePort   = 1666
	unstablePort = 1667
)

// TCP is a TCP connection IPC handler.
type TCP struct {
	addr string
}

// NewTCP returns a new TCP connection handler based on the IP of the remote
// host and port number.
func NewTCP(ip string, port int) *TCP {
	addr := fmt.Sprintf("%s:%d", ip, port)
	return &TCP{addr: addr}
}

// NewStableTCP returns a new stable TCP connection IPC handler based on the IP
// of the remote host.
func NewStableTCP(ip string) *TCP {
	return NewTCP(ip, stablePort)
}

// NewUnstableTCP returns a new unstable TCP connection IPC handler based on the
// IP of the remote host.
func NewUnstableTCP(ip string) *TCP {
	return NewTCP(ip, unstablePort)
}

// Dial connects to the remote host over a TCP connection.
func (i *TCP) Dial() (net.Conn, error) {
	return net.Dial("tcp", i.addr)
}

// Listen listens for incoming connections on a TCP connection.
func (i *TCP) Listen() (net.Listener, error) {
	return net.Listen("tcp", i.addr)
}

// Addr returns the address of the TCP connection.
func (i *TCP) Addr() string {
	return i.addr
}

// TODO: Add UDP?
