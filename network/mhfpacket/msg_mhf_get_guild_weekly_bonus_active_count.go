package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetGuildWeeklyBonusActiveCount represents the MSG_MHF_GET_GUILD_WEEKLY_BONUS_ACTIVE_COUNT
type MsgMhfGetGuildWeeklyBonusActiveCount struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetGuildWeeklyBonusActiveCount) Opcode() network.PacketID {
	return network.MSG_MHF_GET_GUILD_WEEKLY_BONUS_ACTIVE_COUNT
}

// Parse parses the packet from binary
func (m *MsgMhfGetGuildWeeklyBonusActiveCount) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetGuildWeeklyBonusActiveCount) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
