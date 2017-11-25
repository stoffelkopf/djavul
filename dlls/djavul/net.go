package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/internal/proto"
)

// initFrontConn initializes the connection to the front-end.
func initFrontConn() error {
	// Initialize TCP connection.
	const tmpDir = "/tmp/djavul"
	tcpRPath := filepath.Join(tmpDir, "tcp_r")
	tcpR, err := os.OpenFile(tcpRPath, os.O_RDWR, 0644)
	if err != nil {
		return errors.WithStack(err)
	}
	//tcpWPath := filepath.Join(tmpDir, "tcp_w")
	//tcpW, err := os.Open(tcpWPath)
	//if err != nil {
	//	return errors.WithStack(err)
	//}
	fmt.Printf("Connected to %q.\n", tcpRPath)
	proto.EncTCP = gob.NewEncoder(tcpR)
	//proto.DecTCP = gob.NewDecoder(tcpW)

	// Initialize UDP connection.
	udpRPath := filepath.Join(tmpDir, "udp_r")
	udpR, err := os.OpenFile(udpRPath, os.O_RDWR, 0644)
	if err != nil {
		return errors.WithStack(err)
	}
	//udpWPath := filepath.Join(tmpDir, "udp_w")
	//udpW, err := os.Open(udpWPath)
	//if err != nil {
	//	return errors.WithStack(err)
	//}
	fmt.Printf("Connected to %q.\n", udpRPath)
	proto.EncUDP = gob.NewEncoder(udpR)
	//proto.DecUDP = gob.NewDecoder(udpW)
	return nil
}
