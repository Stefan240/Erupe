package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfPostTinyBin represents the MSG_MHF_POST_TINY_BIN
type MsgMhfPostTinyBin struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfPostTinyBin) Opcode() network.PacketID {
	return network.MSG_MHF_POST_TINY_BIN
}

// Parse parses the packet from binary
func (m *MsgMhfPostTinyBin) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfPostTinyBin) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
