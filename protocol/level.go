package protocol

import (
	"encoding/binary"
)

type MsgS2CLevelUp struct {
	MsgHead
	Level uint16
}

func (msg *MsgS2CLevelUp) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgS2CLevelUp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CLevelUp(level uint16) *MsgS2CLevelUp {
	msg := MsgS2CLevelUp{
		MsgHead: MsgHead{Protocol: S2CLevelUp, MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x03, Cmd: 0xFF}},
		Level:   level,
	}
	msg.SetSize()
	return &msg
}
