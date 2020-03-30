package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/mhfpacket/pctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysStageDestruct represents the MSG_SYS_STAGE_DESTRUCT
type MsgSysStageDestruct struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysStageDestruct) Opcode() network.PacketID {
	return network.MSG_SYS_STAGE_DESTRUCT
}

// Parse parses the packet from binary
func (m *MsgSysStageDestruct) Parse(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysStageDestruct) Build(bf *byteframe.ByteFrame, pctx *pctx.PacketContext) error {
	// No data other than opcode
	return nil
}
