package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/common/bfutil"
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfCreateGuild represents the MSG_MHF_CREATE_GUILD
type MsgMhfCreateGuild struct {
	AckHandle uint32
	Unk0      uint8
	Unk1      uint8
	Name      string
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfCreateGuild) Opcode() network.PacketID {
	return network.MSG_MHF_CREATE_GUILD
}

// Parse parses the packet from binary
func (m *MsgMhfCreateGuild) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint8()
	m.Unk1 = bf.ReadUint8()
	nameLength := bf.ReadUint16()
	nameBytes := bfutil.UpToNull(bf.ReadBytes(uint(nameLength)))
	m.Name = pctx.StrConv.MustDecode(nameBytes)

	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfCreateGuild) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
