package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysInfokyserver represents the MSG_SYS_INFOKYSERVER
type MsgSysInfokyserver struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysInfokyserver) Opcode() network.PacketID {
	return network.MSG_SYS_INFOKYSERVER
}

// Parse parses the packet from binary
func (m *MsgSysInfokyserver) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysInfokyserver) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
