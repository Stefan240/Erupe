package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfSaveFavoriteQuest represents the MSG_MHF_SAVE_FAVORITE_QUEST
type MsgMhfSaveFavoriteQuest struct {
	AckHandle uint32
	DataSize  uint16
	Data      []byte
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfSaveFavoriteQuest) Opcode() network.PacketID {
	return network.MSG_MHF_SAVE_FAVORITE_QUEST
}

// Parse parses the packet from binary
func (m *MsgMhfSaveFavoriteQuest) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	m.DataSize = bf.ReadUint16()
	m.Data = bf.ReadBytes(uint(m.DataSize))
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfSaveFavoriteQuest) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
