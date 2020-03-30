package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetDistDescription represents the MSG_MHF_GET_DIST_DESCRIPTION
type MsgMhfGetDistDescription struct {
	AckHandle uint32
	Unk0      uint8
	EntryID   uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetDistDescription) Opcode() network.PacketID {
	return network.MSG_MHF_GET_DIST_DESCRIPTION
}

// Parse parses the packet from binary
func (m *MsgMhfGetDistDescription) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint8()
	m.EntryID = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetDistDescription) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
