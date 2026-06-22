package protocol

type MsgC2SAskWarpZ2B struct {
	MsgHead
	BtWarpData ZoneServerBattleWarpData
}

func (msg *MsgC2SAskWarpZ2B) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskWarpZ2B) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskWarpZ2B(pcId uint32) MsgC2SAskWarpZ2B {
	msg := MsgC2SAskWarpZ2B{
		MsgHead: MsgHead{
			Protocol: C2SAskWarpZ2B,
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

type MsgS2CAnsWarpZ2B struct {
	MsgHead
	ByResult uint8
}

func (msg *MsgS2CAnsWarpZ2B) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsWarpZ2B) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsWarpZ2B(pcId uint32) MsgS2CAnsWarpZ2B {
	msg := MsgS2CAnsWarpZ2B{
		MsgHead: MsgHead{
			Protocol: S2CAnsWarpZ2B,
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

type MsgC2SAskWarpB2Z struct {
	MsgHead
}

func (msg *MsgC2SAskWarpB2Z) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskWarpB2Z) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskWarpB2Z(pcId uint32) MsgC2SAskWarpB2Z {
	msg := MsgC2SAskWarpB2Z{
		MsgHead: MsgHead{
			Protocol: C2SAskWarpB2Z,
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

type MsgS2CAnsWarpB2Z struct {
	MsgHead
	DwNewCell uint32
	WearList  [0xa]ZoneServerItemInWear
	HaveList  [0x1e]ZoneServerItemInInven
	PetActive ZoneServerPetInfo
	PetInven  [0x5]ZoneServerPetInfo
}

func (msg *MsgS2CAnsWarpB2Z) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsWarpB2Z) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsWarpB2Z(pcId uint32) MsgS2CAnsWarpB2Z {
	msg := MsgS2CAnsWarpB2Z{
		MsgHead: MsgHead{
			Protocol: S2CAnsWarpB2Z,
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

type MsgC2SAskPvp struct {
	MsgHead
	DwOtherPCID uint32
}

func (msg *MsgC2SAskPvp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskPvp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CAskPvp struct {
	MsgHead
	DwOtherPCID uint32
}

func (msg *MsgS2CAskPvp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAskPvp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAskPvp(pcId uint32) MsgS2CAskPvp {
	msg := MsgS2CAskPvp{
		MsgHead: MsgHead{
			Protocol: S2CAskPvp,
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

type MsgC2SAnsPvp struct {
	MsgHead
	DwOtherPCID uint32
	ByAnswer    uint8
}

func (msg *MsgC2SAnsPvp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAnsPvp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CAnsPvp struct {
	MsgHead
	DwOtherPCID uint32
	ByAnswer    uint8
}

func (msg *MsgS2CAnsPvp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsPvp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsPvp(pcId uint32) MsgS2CAnsPvp {
	msg := MsgS2CAnsPvp{
		MsgHead: MsgHead{
			Protocol: S2CAnsPvp,
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

type MsgC2SAskShopInfo struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SAskShopInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskShopInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskShopInfo(pcId uint32) MsgC2SAskShopInfo {
	msg := MsgC2SAskShopInfo{
		MsgHead: MsgHead{
			Protocol: C2SAskShopInfo,
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

type MsgS2CAnsShopInfo struct {
	MsgHead
	ByNation  uint8
	ByShop    uint8
	SzNames   [0x3][0x20]byte
	SzPCNames [0x2][0xd]byte
}

func (msg *MsgS2CAnsShopInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsShopInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsShopInfo(pcId uint32) MsgS2CAnsShopInfo {
	msg := MsgS2CAnsShopInfo{
		MsgHead: MsgHead{
			Protocol: S2CAnsShopInfo,
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

type MsgC2SAskGiveMyTax struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SAskGiveMyTax) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskGiveMyTax) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskGiveMyTax(pcId uint32) MsgC2SAskGiveMyTax {
	msg := MsgC2SAskGiveMyTax{
		MsgHead: MsgHead{
			Protocol: C2SAskGiveMyTax,
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

type MsgS2CAnsGiveMyTax struct {
	MsgHead
	ByAns   uint8
	DwMoney uint32
	DwTax   uint32
}

func (msg *MsgS2CAnsGiveMyTax) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsGiveMyTax) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsGiveMyTax(pcId uint32) MsgS2CAnsGiveMyTax {
	msg := MsgS2CAnsGiveMyTax{
		MsgHead: MsgHead{
			Protocol: S2CAnsGiveMyTax,
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
