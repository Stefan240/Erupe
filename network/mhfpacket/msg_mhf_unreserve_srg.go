package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfUnreserveSrg represents the MSG_MHF_UNRESERVE_SRG
type MsgMhfUnreserveSrg struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUnreserveSrg) Opcode() network.PacketID {
	return network.MSG_MHF_UNRESERVE_SRG
}

// Parse parses the packet from binary
func (m *MsgMhfUnreserveSrg) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUnreserveSrg) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
