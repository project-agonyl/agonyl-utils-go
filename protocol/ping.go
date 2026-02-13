package protocol

import (
	"encoding/binary"
)

type MsgZACLChkTimeTick struct {
	MsgHeadNoProtocol
	TickCount uint32
	TickSvr   uint32
	TickClt   uint32
}

func (msg *MsgZACLChkTimeTick) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgZACLChkTimeTick) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgZACLChkTimeTick(pcId uint32, tickCount uint32, tickSvr uint32) *MsgZACLChkTimeTick {
	msg := MsgZACLChkTimeTick{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x01, Cmd: 0xF0, PcId: pcId},
		TickCount:         tickCount,
		TickSvr:           tickSvr,
	}
	msg.SetSize()
	return &msg
}
