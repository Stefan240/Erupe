package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetUdTacticsRewardList represents the MSG_MHF_GET_UD_TACTICS_REWARD_LIST
type MsgMhfGetUdTacticsRewardList struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetUdTacticsRewardList) Opcode() network.PacketID {
	return network.MSG_MHF_GET_UD_TACTICS_REWARD_LIST
}

// Parse parses the packet from binary
func (m *MsgMhfGetUdTacticsRewardList) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetUdTacticsRewardList) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
