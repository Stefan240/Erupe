package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetGachaPlayHistory represents the MSG_MHF_GET_GACHA_PLAY_HISTORY
type MsgMhfGetGachaPlayHistory struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetGachaPlayHistory) Opcode() network.PacketID {
	return network.MSG_MHF_GET_GACHA_PLAY_HISTORY
}

// Parse parses the packet from binary
func (m *MsgMhfGetGachaPlayHistory) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetGachaPlayHistory) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
