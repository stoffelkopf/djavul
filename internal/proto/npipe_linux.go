package proto

import (
	"net"
)

// Note, there exist no Windows implementation of Unix sockets, so the
// djavul.exe backend cannot connect to djavul-frontend.exe. Thus, to avoid
// confusion, we panic rather than implement implementation for only one.

// Dial connects to the remote host over a named pipe.
func (i *NamedPipe) Dial() (net.Conn, error) {
	panic("proto.NamedPipe.Dial: not yet implemented on Linux")
	//return net.Dial("unix", i.Addr())
}

// Listen listens for incoming connections on a named pipe.
func (i *NamedPipe) Listen() (net.Listener, error) {
	panic("proto.NamedPipe.Dial: not yet implemented on Linux")
	//return net.Listen("unix", i.Addr())
}

// Addr returns the address of the named pipe.
func (i *NamedPipe) Addr() string {
	panic("proto.NamedPipe.Dial: not yet implemented on Linux")
	//if i.ip != "localhost" {
	//	panic(fmt.Errorf("proto.NamedPipe.Addr: support for named pipe to remote IP (%s) not yet supported over Unix socket", i.ip))
	//}
	//return fmt.Sprintf(`/tmp/djavul_%s.sock`, i.pipe)
}
