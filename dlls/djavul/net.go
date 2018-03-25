//+build djavul

package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
	"log"

	npipe "net" //"github.com/natefinch/npipe"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/internal/proto"
)

// initFrontConn initializes the connection to the front-end.
func initFrontConn(frontendIP string) error {
	// Initialize TCP connection.
	fmt.Printf("Connecting to %q.\n", frontendIP+proto.TCPReadPipe)
	tcpR, err := npipe.Dial("tcp", frontendIP+proto.TCPReadPipe)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Printf("Writing to %q.\n", frontendIP+proto.TCPReadPipe)
	proto.EncTCP = gob.NewEncoder(tcpR)

	//fmt.Printf("Connecting to %q.\n", frontendIP+proto.TCPWritePipe)
	//tcpW, err := npipe.Dial("tcp", frontendIP+proto.TCPWritePipe)
	//if err != nil {
	//	return errors.WithStack(err)
	//}
	//fmt.Printf("Reading from %q.\n", frontendIP+proto.TCPWritePipe)
	//proto.DecTCP = gob.NewDecoder(tcpW)

	// Initialize UDP connection.
	fmt.Printf("Connecting to %q.\n", frontendIP+proto.UDPReadPipe)
	udpR, err := npipe.Dial("tcp", frontendIP+proto.UDPReadPipe)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Printf("Writing to %q.\n", frontendIP+proto.UDPReadPipe)
	proto.EncUDP = gob.NewEncoder(udpR)

	fmt.Printf("Connecting to %q.\n", frontendIP+proto.UDPWritePipe)
	udpW, err := npipe.Dial("tcp", frontendIP+proto.UDPWritePipe)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Printf("Reading from %q.\n", frontendIP+proto.UDPWritePipe)
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
