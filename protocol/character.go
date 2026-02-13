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
