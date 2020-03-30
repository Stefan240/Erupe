package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysReserve79 represents the MSG_SYS_reserve79
type MsgSysReserve79 struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve79) Opcode() network.PacketID {
	return network.MSG_SYS_reserve79
}

// Parse parses the packet from binary
func (m *MsgSysReserve79) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve79) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
