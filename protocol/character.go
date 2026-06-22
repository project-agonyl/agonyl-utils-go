package protocol

import "encoding/binary"

type MsgC2SAskDeletePlayer struct {
	MsgHead
	CharacterName [0x15]byte
}

func (msg *MsgC2SAskDeletePlayer) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgC2SAskDeletePlayer) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskDeletePlayer(pcId uint32, characterName string) MsgC2SAskDeletePlayer {
	msg := MsgC2SAskDeletePlayer{
		MsgHead: MsgHead{
			Protocol: C2SAskDeletePlayer,
			MsgHeadNoProtocol: MsgHeadNoProtocol{
				Ctrl: 0x03,
				Cmd:  0xFF,
				PcId: pcId,
			},
		},
	}

	copy(msg.CharacterName[:], characterName)
	msg.PcId = pcId
	msg.SetSize()
	return msg
}

type AclCharacterWear struct {
	ItemPtr    uint32
	ItemCode   uint32
	ItemOption uint32
	WearIndex  uint32
}

type CharacterInfo struct {
	Name     [0x15]byte
	SlotUsed byte
	Class    byte
	Nation   byte
	Level    uint32
	Wear     [0xA]AclCharacterWear
}

type MsgS2CCharacterList struct {
	MsgHead
	CharacterList [0x5]CharacterInfo
}

func (msg *MsgS2CCharacterList) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgS2CCharacterList) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CCharacterList(pcId uint32, characterList []CharacterInfo) MsgS2CCharacterList {
	msgS2CCharacterList := MsgS2CCharacterList{
		MsgHead: MsgHead{
			Protocol: S2CCharacterList,
			MsgHeadNoProtocol: MsgHeadNoProtocol{
				Ctrl: 0x03,
				Cmd:  0xFF,
				PcId: pcId,
			},
		},
		CharacterList: [5]CharacterInfo{},
	}

	for i := range 5 {
		if i < len(characterList) {
			msgS2CCharacterList.CharacterList[i] = characterList[i]
		} else {
			msgS2CCharacterList.CharacterList[i].Class = 255
		}
	}

	msgS2CCharacterList.SetSize()
	return msgS2CCharacterList
}

func NewMsgS2CCharacterListEmpty(pcId uint32) MsgS2CCharacterList {
	msgS2CCharacterList := MsgS2CCharacterList{
		MsgHead: MsgHead{
			Protocol: S2CCharacterList,
			MsgHeadNoProtocol: MsgHeadNoProtocol{
				Ctrl: 0x03,
				Cmd:  0xFF,
				PcId: pcId,
			},
		},
		CharacterList: [5]CharacterInfo{},
	}

	for i := range msgS2CCharacterList.CharacterList {
		msgS2CCharacterList.CharacterList[i].Class = 255
	}

	msgS2CCharacterList.SetSize()
	return msgS2CCharacterList
}

type MsgS2CWarpReady struct {
	MsgHead
	WNewMapIndex uint16
}

func (msg *MsgS2CWarpReady) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CWarpReady) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CWarpReady(pcId uint32) MsgS2CWarpReady {
	msg := MsgS2CWarpReady{
		MsgHead: MsgHead{
			Protocol: S2CWarpReady,
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

type MsgS2CWarp struct {
	MsgHead
	DwNewCell uint32
}

func (msg *MsgS2CWarp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CWarp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CWarp(pcId uint32) MsgS2CWarp {
	msg := MsgS2CWarp{
		MsgHead: MsgHead{
			Protocol: S2CWarp,
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

type MsgS2CWarpLogin struct {
	MsgHead
	DwNewCell uint32
	WearList  [0xa]ZoneServerItemInWear
	HaveList  [0x1e]ZoneServerItemInInven
	PetActive ZoneServerPetInfo
	PetInven  [0x5]ZoneServerPetInfo
}

func (msg *MsgS2CWarpLogin) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CWarpLogin) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CWarpLogin(pcId uint32) MsgS2CWarpLogin {
	msg := MsgS2CWarpLogin{
		MsgHead: MsgHead{
			Protocol: S2CWarpLogin,
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

type MsgS2CCharLogout struct {
	MsgHead
}

func (msg *MsgS2CCharLogout) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CCharLogout) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CCharLogout(pcId uint32) MsgS2CCharLogout {
	msg := MsgS2CCharLogout{
		MsgHead: MsgHead{
			Protocol: S2CCharLogout,
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
