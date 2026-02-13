package protocol

import "encoding/binary"

type MsgC2SReqClanInfo struct {
	MsgHead
}

func (msg *MsgC2SReqClanInfo) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgC2SReqClanInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SReqClanInfo(pcId uint32) MsgC2SReqClanInfo {
	msg := MsgC2SReqClanInfo{
		MsgHead: MsgHead{
			Protocol: C2SReqClanInfo,
			MsgHeadNoProtocol: MsgHeadNoProtocol{
				Ctrl: 0x03,
				Cmd:  0xFF,
				PcId: pcId,
			},
		},
	}
	msg.SetSize()
	return msg
}

type ClanMate struct {
	CharacterName [0x15]byte
	Unknown1      [0xB]byte
	Class         byte
	Unknown2      [0x3]byte
}

type MsgS2CClanInfo struct {
	MsgHead
	ClanName  [0x20]byte
	Unknown1  uint16
	Unknown2  uint16
	Unknown3  byte
	Unknown4  uint16
	Unknown5  uint32
	Unknown6  uint32
	ClanMates [0xD]ClanMate
}

func (msg *MsgS2CClanInfo) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgS2CClanInfo) SetSize() {
	msg.Size = msg.GetSize()
}
