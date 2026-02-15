package protocol

import "encoding/binary"

type MsgS2CError struct {
	MsgHead
	Code uint16
	Msg  [64]byte
}

func (msg *MsgS2CError) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgS2CError) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CError(pcId uint32, code uint16, msg string) *MsgS2CError {
	msgS2CError := MsgS2CError{
		MsgHead: MsgHead{Protocol: S2CError, MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x03, Cmd: 0xFF}},
		Code:    code,
	}

	copy(msgS2CError.Msg[:], msg)
	msgS2CError.PcId = pcId
	msgS2CError.SetSize()
	return &msgS2CError
}
