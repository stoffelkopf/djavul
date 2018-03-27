//+build djavul

package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
	"log"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/internal/proto"
)

// initFrontConn initializes the connection to the front-end.
func initFrontConn(stable, unstable proto.IPC) error {
	// Initialize stable connection.
	fmt.Printf("Connecting to %q.\n", stable.Addr())
	stableConn, err := stable.Dial()
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Printf("Connected to %q.\n", stableConn.RemoteAddr())
	proto.EncStable = gob.NewEncoder(stableConn)
	//proto.DecStable = gob.NewDecoder(stableConn)

	// Initialize unstable connection.
	fmt.Printf("Connecting to %q.\n", unstable.Addr())
	unstableConn, err := unstable.Dial()
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Printf("Connected to %q.\n", unstableConn.RemoteAddr())
	proto.EncUnstable = gob.NewEncoder(unstableConn)
	proto.DecUnstable = gob.NewDecoder(unstableConn)

	go recvActions()
	return nil
}

// recvActions receives action packets from the front-end.
func recvActions() {
	proto.Actions = make(chan proto.EngineAction)
	for {
		var pkt proto.PacketUnstable
		if err := proto.DecUnstable.Decode(&pkt); err != nil {
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
