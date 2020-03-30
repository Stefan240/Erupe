package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetUdSelectedColorInfo represents the MSG_MHF_GET_UD_SELECTED_COLOR_INFO
type MsgMhfGetUdSelectedColorInfo struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetUdSelectedColorInfo) Opcode() network.PacketID {
	return network.MSG_MHF_GET_UD_SELECTED_COLOR_INFO
}

// Parse parses the packet from binary
func (m *MsgMhfGetUdSelectedColorInfo) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetUdSelectedColorInfo) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
