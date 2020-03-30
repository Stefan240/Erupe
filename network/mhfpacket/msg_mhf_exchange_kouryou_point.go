package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfExchangeKouryouPoint represents the MSG_MHF_EXCHANGE_KOURYOU_POINT
type MsgMhfExchangeKouryouPoint struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfExchangeKouryouPoint) Opcode() network.PacketID {
	return network.MSG_MHF_EXCHANGE_KOURYOU_POINT
}

// Parse parses the packet from binary
func (m *MsgMhfExchangeKouryouPoint) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfExchangeKouryouPoint) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
