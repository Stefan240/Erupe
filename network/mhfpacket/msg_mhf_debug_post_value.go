package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfDebugPostValue represents the MSG_MHF_DEBUG_POST_VALUE
type MsgMhfDebugPostValue struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfDebugPostValue) Opcode() network.PacketID {
	return network.MSG_MHF_DEBUG_POST_VALUE
}

// Parse parses the packet from binary
func (m *MsgMhfDebugPostValue) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfDebugPostValue) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
