package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfCaravanRanking represents the MSG_MHF_CARAVAN_RANKING
type MsgMhfCaravanRanking struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfCaravanRanking) Opcode() network.PacketID {
	return network.MSG_MHF_CARAVAN_RANKING
}

// Parse parses the packet from binary
func (m *MsgMhfCaravanRanking) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfCaravanRanking) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
