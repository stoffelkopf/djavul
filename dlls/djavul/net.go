package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/internal/proto"
)

// initFrontConn initializes the connection to the front-end.
func initFrontConn() error {
	// Initialize TCP connection.
	tcpConn, err := net.Dial("tcp", "localhost:6667")
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Println("Connected to:", tcpConn.RemoteAddr())
	proto.EncTCP = gob.NewEncoder(tcpConn)
	proto.DecTCP = gob.NewDecoder(tcpConn)
	// Initialize UDP connection.
	laddr, err := net.ResolveUDPAddr("udp", "localhost:12345")
	if err != nil {
		return errors.WithStack(err)
	}
	raddr, err := net.ResolveUDPAddr("udp", "localhost:6666")
	if err != nil {
		return errors.WithStack(err)
	}
	udpConn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Println("Connected to:", udpConn.RemoteAddr())
	proto.EncUDP = gob.NewEncoder(udpConn)
	proto.DecUDP = gob.NewDecoder(udpConn)
	return nil
}
