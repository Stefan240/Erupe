package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfSaveRengokuData represents the MSG_MHF_SAVE_RENGOKU_DATA
type MsgMhfSaveRengokuData struct {
	AckHandle      uint32
	DataSize       uint32
	RawDataPayload []byte
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfSaveRengokuData) Opcode() network.PacketID {
	return network.MSG_MHF_SAVE_RENGOKU_DATA
}

// Parse parses the packet from binary
func (m *MsgMhfSaveRengokuData) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	m.DataSize = bf.ReadUint32()
	m.RawDataPayload = bf.ReadBytes(uint(m.DataSize))
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfSaveRengokuData) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
