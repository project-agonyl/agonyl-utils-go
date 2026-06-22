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

func NewMsgS2CClanInfo(pcId uint32) MsgS2CClanInfo {
	msg := MsgS2CClanInfo{
		MsgHead: MsgHead{
			Protocol: S2CClanInfo,
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

type MsgS2CClan struct {
	MsgHead
	RequestType  uint8
	Result       uint8
	ClanID       uint32
	MySocialInfo ZoneServerSocialinfo
}

func (msg *MsgS2CClan) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CClan) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CClan(pcId uint32) MsgS2CClan {
	msg := MsgS2CClan{
		MsgHead: MsgHead{
			Protocol: S2CClan,
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

type MsgS2CAskClan struct {
	MsgHead
	DwAnsPCID uint32
}

func (msg *MsgS2CAskClan) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAskClan) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAskClan(pcId uint32) MsgS2CAskClan {
	msg := MsgS2CAskClan{
		MsgHead: MsgHead{
			Protocol: S2CAskClan,
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

type MsgS2CJoinClan struct {
	MsgHead
	ResultCode   uint8
	MySocialInfo ZoneServerSocialinfo
}

func (msg *MsgS2CJoinClan) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CJoinClan) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CJoinClan(pcId uint32) MsgS2CJoinClan {
	msg := MsgS2CJoinClan{
		MsgHead: MsgHead{
			Protocol: S2CJoinClan,
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

type MsgS2CBoltClan struct {
	MsgHead
	ResultCode   uint8
	MySocialInfo ZoneServerSocialinfo
}

func (msg *MsgS2CBoltClan) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CBoltClan) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CBoltClan(pcId uint32) MsgS2CBoltClan {
	msg := MsgS2CBoltClan{
		MsgHead: MsgHead{
			Protocol: S2CBoltClan,
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

type MsgS2CUpdateClan struct {
	MsgHead
	SzClanName    [0x20]byte
	DwClanID      uint32
	DwMarkID      uint32
	DwStorageID   uint32
	ClanRank      uint8
	WAgitId       uint16
	ArrKnightInfo [0xd]ZoneServerKnightInfo
}

func (msg *MsgS2CUpdateClan) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CUpdateClan) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CUpdateClan(pcId uint32) MsgS2CUpdateClan {
	msg := MsgS2CUpdateClan{
		MsgHead: MsgHead{
			Protocol: S2CUpdateClan,
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

type MsgS2CClanFail struct {
	MsgHead
	ByErrorCode  uint8
	DwClanID     uint32
	MySocialInfo ZoneServerSocialinfo
}

func (msg *MsgS2CClanFail) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CClanFail) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CClanFail(pcId uint32) MsgS2CClanFail {
	msg := MsgS2CClanFail{
		MsgHead: MsgHead{
			Protocol: S2CClanFail,
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

type MsgS2CChangeNation struct {
	MsgHead
	ByErrorCode  uint8
	MySocialInfo ZoneServerSocialinfo
	DwLore       uint32
}

func (msg *MsgS2CChangeNation) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CChangeNation) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CChangeNation(pcId uint32) MsgS2CChangeNation {
	msg := MsgS2CChangeNation{
		MsgHead: MsgHead{
			Protocol: S2CChangeNation,
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

type MsgC2SClan struct {
	MsgHead
	RequestType  uint8
	SzClanName   [0x20]byte
	SzClanPasswd [0x10]byte
}

func (msg *MsgC2SClan) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SClan) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SClan(pcId uint32) MsgC2SClan {
	msg := MsgC2SClan{
		MsgHead: MsgHead{
			Protocol: C2SClan,
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

type MsgC2SJoinClan struct {
	MsgHead
	DwAskPCID uint32
}

func (msg *MsgC2SJoinClan) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SJoinClan) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SJoinClan(pcId uint32) MsgC2SJoinClan {
	msg := MsgC2SJoinClan{
		MsgHead: MsgHead{
			Protocol: C2SJoinClan,
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

type MsgC2SAnsClan struct {
	MsgHead
	BAccept   uint8
	DwAskPCID uint32
}

func (msg *MsgC2SAnsClan) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAnsClan) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAnsClan(pcId uint32) MsgC2SAnsClan {
	msg := MsgC2SAnsClan{
		MsgHead: MsgHead{
			Protocol: C2SAnsClan,
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

type MsgC2SBoltClan struct {
	MsgHead
	BDismissal   uint8
	SzKnightName [0xd]byte
}

func (msg *MsgC2SBoltClan) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SBoltClan) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SBoltClan(pcId uint32) MsgC2SBoltClan {
	msg := MsgC2SBoltClan{
		MsgHead: MsgHead{
			Protocol: C2SBoltClan,
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

type MsgC2SChangeNation struct {
	MsgHead
	NewNation uint8
}

func (msg *MsgC2SChangeNation) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SChangeNation) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SChangeNation(pcId uint32) MsgC2SChangeNation {
	msg := MsgC2SChangeNation{
		MsgHead: MsgHead{
			Protocol: C2SChangeNation,
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
