package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysGetState represents the MSG_SYS_GET_STATE
type MsgSysGetState struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysGetState) Opcode() network.PacketID {
	return network.MSG_SYS_GET_STATE
}

// Parse parses the packet from binary
func (m *MsgSysGetState) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysGetState) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
