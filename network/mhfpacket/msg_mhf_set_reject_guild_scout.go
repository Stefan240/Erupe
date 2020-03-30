package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfSetRejectGuildScout represents the MSG_MHF_SET_REJECT_GUILD_SCOUT
type MsgMhfSetRejectGuildScout struct {
	AckHandle uint32
	Reject    bool
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfSetRejectGuildScout) Opcode() network.PacketID {
	return network.MSG_MHF_SET_REJECT_GUILD_SCOUT
}

// Parse parses the packet from binary
func (m *MsgMhfSetRejectGuildScout) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	m.Reject = bf.ReadBool()

	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfSetRejectGuildScout) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
