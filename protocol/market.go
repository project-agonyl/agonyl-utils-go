package protocol

import "encoding/binary"

type MsgC2SOpenMarket struct {
	MsgHead
	Items [80]byte
	Msg   [0x40]byte
}

func (msg *MsgC2SOpenMarket) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgC2SOpenMarket) SetSize() {
	msg.Size = msg.GetSize()
}
