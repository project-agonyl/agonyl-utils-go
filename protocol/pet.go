package protocol

type MsgS2CPetBuy struct {
	MsgHead
	PetIndex uint8
	NewPet   ZoneServerPetInfo
	DwMoney  uint32
}

func (msg *MsgS2CPetBuy) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPetBuy) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPetBuy(pcId uint32) MsgS2CPetBuy {
	msg := MsgS2CPetBuy{
		MsgHead: MsgHead{
			Protocol: S2CPetBuy,
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

type MsgS2CPutInPet struct {
	MsgHead
	Pet ZoneServerPetInfo
}

func (msg *MsgS2CPutInPet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPutInPet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPutInPet(pcId uint32) MsgS2CPutInPet {
	msg := MsgS2CPutInPet{
		MsgHead: MsgHead{
			Protocol: S2CPutInPet,
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

type MsgS2CPutOutPet struct {
	MsgHead
	DwPetID uint32
}

func (msg *MsgS2CPutOutPet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPutOutPet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPutOutPet(pcId uint32) MsgS2CPutOutPet {
	msg := MsgS2CPutOutPet{
		MsgHead: MsgHead{
			Protocol: S2CPutOutPet,
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

type MsgS2CActivePet struct {
	MsgHead
	DwActivePetID uint32
	DwInactPetID  uint32
	PetIndex      uint8
	PC2ndStat     ZoneServerPc2Stat
}

func (msg *MsgS2CActivePet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CActivePet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CActivePet(pcId uint32) MsgS2CActivePet {
	msg := MsgS2CActivePet{
		MsgHead: MsgHead{
			Protocol: S2CActivePet,
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

type MsgS2CInactivePet struct {
	MsgHead
	DwInactPetID uint32
	PetIndex     uint8
	PC2ndStat    ZoneServerPc2Stat
}

func (msg *MsgS2CInactivePet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CInactivePet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CInactivePet(pcId uint32) MsgS2CInactivePet {
	msg := MsgS2CInactivePet{
		MsgHead: MsgHead{
			Protocol: S2CInactivePet,
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

type MsgS2CSeePet struct {
	MsgHead
	DwOPCID   uint32
	PetAppear ZoneServerPetAppear
	PetStat   uint8
}

func (msg *MsgS2CSeePet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSeePet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSeePet(pcId uint32) MsgS2CSeePet {
	msg := MsgS2CSeePet{
		MsgHead: MsgHead{
			Protocol: S2CSeePet,
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

type MsgS2CPetDie struct {
	MsgHead
	DwPetID   uint32
	PC2ndStat ZoneServerPc2Stat
}

func (msg *MsgS2CPetDie) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPetDie) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPetDie(pcId uint32) MsgS2CPetDie {
	msg := MsgS2CPetDie{
		MsgHead: MsgHead{
			Protocol: S2CPetDie,
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

type MsgS2CPetSell struct {
	MsgHead
	PetIndex uint8
	PetID    uint32
	DwMoney  uint32
}

func (msg *MsgS2CPetSell) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPetSell) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPetSell(pcId uint32) MsgS2CPetSell {
	msg := MsgS2CPetSell{
		MsgHead: MsgHead{
			Protocol: S2CPetSell,
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

type MsgS2CFeedPet struct {
	MsgHead
	ItemID    ZoneServerItemId
	Result    uint8
	PetID     uint32
	PetIndex  uint8
	WQuantity uint16
}

func (msg *MsgS2CFeedPet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CFeedPet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CFeedPet(pcId uint32) MsgS2CFeedPet {
	msg := MsgS2CFeedPet{
		MsgHead: MsgHead{
			Protocol: S2CFeedPet,
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

type MsgS2CRevivePet struct {
	MsgHead
	DwItemID  uint32
	PetUpdate ZoneServerPetStat
}

func (msg *MsgS2CRevivePet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CRevivePet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CRevivePet(pcId uint32) MsgS2CRevivePet {
	msg := MsgS2CRevivePet{
		MsgHead: MsgHead{
			Protocol: S2CRevivePet,
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

type MsgS2CShueCombination struct {
	MsgHead
	ByResult      uint8
	UPresentMoney uint32
	ByInvenIndex  uint8
	StPetInfo     ZoneServerPetInfo
}

func (msg *MsgS2CShueCombination) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CShueCombination) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CShueCombination(pcId uint32) MsgS2CShueCombination {
	msg := MsgS2CShueCombination{
		MsgHead: MsgHead{
			Protocol: S2CShueCombination,
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

type MsgC2SActivePet struct {
	MsgHead
	DwPetID uint32
}

func (msg *MsgC2SActivePet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SActivePet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SActivePet(pcId uint32) MsgC2SActivePet {
	msg := MsgC2SActivePet{
		MsgHead: MsgHead{
			Protocol: C2SActivePet,
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

type MsgC2SInactivePet struct {
	MsgHead
}

func (msg *MsgC2SInactivePet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SInactivePet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SInactivePet(pcId uint32) MsgC2SInactivePet {
	msg := MsgC2SInactivePet{
		MsgHead: MsgHead{
			Protocol: C2SInactivePet,
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

type MsgC2SPetBuy struct {
	MsgHead
	DwNPCID uint32
	Code    uint16
}

func (msg *MsgC2SPetBuy) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPetBuy) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SPetBuy(pcId uint32) MsgC2SPetBuy {
	msg := MsgC2SPetBuy{
		MsgHead: MsgHead{
			Protocol: C2SPetBuy,
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

type MsgC2SPetSell struct {
	MsgHead
	DwNPCID uint32
	PetID   uint32
}

func (msg *MsgC2SPetSell) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPetSell) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SPetSell(pcId uint32) MsgC2SPetSell {
	msg := MsgC2SPetSell{
		MsgHead: MsgHead{
			Protocol: C2SPetSell,
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

type MsgC2SFeedPet struct {
	MsgHead
	PetID uint32
}

func (msg *MsgC2SFeedPet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SFeedPet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SFeedPet(pcId uint32) MsgC2SFeedPet {
	msg := MsgC2SFeedPet{
		MsgHead: MsgHead{
			Protocol: C2SFeedPet,
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

type MsgC2SPutInPet struct {
	MsgHead
	DwPetID uint32
}

func (msg *MsgC2SPutInPet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPutInPet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SPutInPet(pcId uint32) MsgC2SPutInPet {
	msg := MsgC2SPutInPet{
		MsgHead: MsgHead{
			Protocol: C2SPutInPet,
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

type MsgC2SPutOutPet struct {
	MsgHead
	DwPetID uint32
}

func (msg *MsgC2SPutOutPet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPutOutPet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SPutOutPet(pcId uint32) MsgC2SPutOutPet {
	msg := MsgC2SPutOutPet{
		MsgHead: MsgHead{
			Protocol: C2SPutOutPet,
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

type MsgC2SRevivePet struct {
	MsgHead
	DwItemID uint32
	DwPetID  uint32
}

func (msg *MsgC2SRevivePet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SRevivePet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SRevivePet(pcId uint32) MsgC2SRevivePet {
	msg := MsgC2SRevivePet{
		MsgHead: MsgHead{
			Protocol: C2SRevivePet,
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

type MsgC2SShueCombination struct {
	MsgHead
	DwNPCID    uint32
	ArrPetInfo [0x2]ZoneServerPetInfo
}

func (msg *MsgC2SShueCombination) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SShueCombination) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SShueCombination(pcId uint32) MsgC2SShueCombination {
	msg := MsgC2SShueCombination{
		MsgHead: MsgHead{
			Protocol: C2SShueCombination,
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
