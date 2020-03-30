package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfRegistGuildAdventure represents the MSG_MHF_REGIST_GUILD_ADVENTURE
type MsgMhfRegistGuildAdventure struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfRegistGuildAdventure) Opcode() network.PacketID {
	return network.MSG_MHF_REGIST_GUILD_ADVENTURE
}

// Parse parses the packet from binary
func (m *MsgMhfRegistGuildAdventure) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfRegistGuildAdventure) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
