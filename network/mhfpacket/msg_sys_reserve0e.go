package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysReserve0E represents the MSG_SYS_reserve0E
type MsgSysReserve0E struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve0E) Opcode() network.PacketID {
	return network.MSG_SYS_reserve0E
}

// Parse parses the packet from binary
func (m *MsgSysReserve0E) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve0E) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
