package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfLoadPlateBox represents the MSG_MHF_LOAD_PLATE_BOX
type MsgMhfLoadPlateBox struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfLoadPlateBox) Opcode() network.PacketID {
	return network.MSG_MHF_LOAD_PLATE_BOX
}

// Parse parses the packet from binary
func (m *MsgMhfLoadPlateBox) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfLoadPlateBox) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
