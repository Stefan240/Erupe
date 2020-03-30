package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireMonthlyReward represents the MSG_MHF_ACQUIRE_MONTHLY_REWARD
type MsgMhfAcquireMonthlyReward struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireMonthlyReward) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_MONTHLY_REWARD
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireMonthlyReward) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireMonthlyReward) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
