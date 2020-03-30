package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireFesta represents the MSG_MHF_ACQUIRE_FESTA
type MsgMhfAcquireFesta struct {
	AckHandle uint32
	Unk0      uint32
	Unk1      uint32
	Unk2      uint32
	Unk3      uint8
	Unk4      uint8 // Hardcoded 0 in binary.
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireFesta) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_FESTA
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireFesta) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireFesta) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
