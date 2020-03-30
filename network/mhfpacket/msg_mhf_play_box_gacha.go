package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfPlayBoxGacha represents the MSG_MHF_PLAY_BOX_GACHA
type MsgMhfPlayBoxGacha struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfPlayBoxGacha) Opcode() network.PacketID {
	return network.MSG_MHF_PLAY_BOX_GACHA
}

// Parse parses the packet from binary
func (m *MsgMhfPlayBoxGacha) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfPlayBoxGacha) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
