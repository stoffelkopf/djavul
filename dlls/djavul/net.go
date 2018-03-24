//+build djavul

package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/internal/proto"
)

// initFrontConn initializes the connection to the front-end.
func initFrontConn() error {
	// Initialize TCP connection.
	const tmpDir = `C:\temp\djavul`
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return errors.WithStack(err)
	}
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
	udpWPath := filepath.Join(tmpDir, "udp_w")
	udpW, err := os.Open(udpWPath)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Printf("Connected to %q.\n", udpRPath)
	proto.EncUDP = gob.NewEncoder(udpR)
	proto.DecUDP = gob.NewDecoder(udpW)
	go recvActions()
	return nil
}

// recvActions receives action packets from the front-end.
func recvActions() {
	proto.Actions = make(chan proto.EngineAction)
	for {
		var pkt proto.PacketUDP
		if err := proto.DecUDP.Decode(&pkt); err != nil {
			if errors.Cause(err) == io.EOF {
				fmt.Println("disconnected")
				break
			}
		}
		switch pkt.Cmd {
		case proto.CmdButtonPressedAction:
			var action proto.ButtonPressedAction
			if err := binary.Read(bytes.NewReader(pkt.Data), binary.LittleEndian, &action); err != nil {
				log.Fatalf("%+v", errors.WithStack(err))
			}
			proto.Actions <- action
		case proto.CmdButtonReleasedAction:
			var action proto.ButtonReleasedAction
			if err := binary.Read(bytes.NewReader(pkt.Data), binary.LittleEndian, &action); err != nil {
				log.Fatalf("%+v", errors.WithStack(err))
			}
			proto.Actions <- action
		default:
			panic(fmt.Errorf("support for packet cmd %v not yet implemented", pkt.Cmd))
		}
	}
}
