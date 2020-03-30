package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfOperateJoint represents the MSG_MHF_OPERATE_JOINT
type MsgMhfOperateJoint struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfOperateJoint) Opcode() network.PacketID {
	return network.MSG_MHF_OPERATE_JOINT
}

// Parse parses the packet from binary
func (m *MsgMhfOperateJoint) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfOperateJoint) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}
