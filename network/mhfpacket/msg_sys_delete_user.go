package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysDeleteUser represents the MSG_SYS_DELETE_USER
type MsgSysDeleteUser struct {
	CharID uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysDeleteUser) Opcode() network.PacketID {
	return network.MSG_SYS_DELETE_USER
}

// Parse parses the packet from binary
func (m *MsgSysDeleteUser) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysDeleteUser) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	bf.WriteUint32(m.CharID)

	return nil
}
