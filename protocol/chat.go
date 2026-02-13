package protocol

import (
	"bytes"
	"encoding/binary"
)

type SayType byte

const (
	Notice     SayType = 0x0C
	System     SayType = 0x00
	General    SayType = 0x01
	Whisper    SayType = 0x03
	Party      SayType = 0x04
	Knighthood SayType = 0x05
	Country    SayType = 0x06
	Alliance   SayType = 0x08
	Shout      SayType = 0xF1
)

type MsgC2SSay struct {
	MsgHead
	SayType SayType
	SayPC   [0x15]byte
	Words   [0x40]byte
}

func (msg *MsgC2SSay) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgC2SSay) SetSize() {
	msg.Size = msg.GetSize()
}

func (msg *MsgC2SSay) GetBytes() []byte {
	var buffer bytes.Buffer
	_ = binary.Write(&buffer, binary.LittleEndian, msg)
	return buffer.Bytes()
}

func NewMsgC2SSay(pcId uint32, sayType SayType, sayPC string, words string) MsgC2SSay {
	msg := MsgC2SSay{
		MsgHead: MsgHead{
			MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x03, Cmd: 0xFF, PcId: pcId},
			Protocol:          C2SSay,
		},
		SayType: sayType,
	}
	copy(msg.SayPC[:], sayPC)
	copy(msg.Words[:], words)
	msg.SetSize()
	return msg
}

type MsgS2CSay struct {
	MsgHead
	SayType SayType
	SayPcId uint32
	SayPC   [0x15]byte
	Words   [0x40]byte
}

func (msg *MsgS2CSay) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgS2CSay) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSay(pcId uint32, sayType SayType, sayPC string, words string) MsgS2CSay {
	msgS2CSay := MsgS2CSay{
		MsgHead: MsgHead{
			Protocol:          C2SSay,
			MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x03, Cmd: 0xFF, PcId: pcId},
		},
		SayType: sayType,
		SayPcId: pcId,
	}

	copy(msgS2CSay.SayPC[:], sayPC)
	copy(msgS2CSay.Words[:], words)
	msgS2CSay.SetSize()
	return msgS2CSay
}
