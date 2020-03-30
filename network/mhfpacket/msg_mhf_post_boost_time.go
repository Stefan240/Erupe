package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfPostBoostTime represents the MSG_MHF_POST_BOOST_TIME
type MsgMhfPostBoostTime struct {
	AckHandle uint32
	BoostTime uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfPostBoostTime) Opcode() network.PacketID {
	return network.MSG_MHF_POST_BOOST_TIME
}

// Parse parses the packet from binary
func (m *MsgMhfPostBoostTime) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	m.BoostTime = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfPostBoostTime) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
