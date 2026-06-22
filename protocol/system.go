package protocol

type MsgC2SOpCmd struct {
	MsgHead
	WOpCmd  uint16
	SzParam [0x40]byte
}

func (msg *MsgC2SOpCmd) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SOpCmd) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SOpCmd(pcId uint32) MsgC2SOpCmd {
	msg := MsgC2SOpCmd{
		MsgHead: MsgHead{
			Protocol: C2SOpCmd,
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

type MsgC2SUnknownChar struct {
	MsgHead
	ByType       uint8
	DwID         uint32
	WRetProtocol uint16
}

func (msg *MsgC2SUnknownChar) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SUnknownChar) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SUnknownChar(pcId uint32) MsgC2SUnknownChar {
	msg := MsgC2SUnknownChar{
		MsgHead: MsgHead{
			Protocol: C2SUnknownChar,
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

type MsgS2CUnknownChar struct {
	MsgHead
	ByType       uint8
	ByCode       uint8
	DwID         uint32
	WRetProtocol uint16
}

func (msg *MsgS2CUnknownChar) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CUnknownChar) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CUnknownChar(pcId uint32) MsgS2CUnknownChar {
	msg := MsgS2CUnknownChar{
		MsgHead: MsgHead{
			Protocol: S2CUnknownChar,
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

type MsgS2CPayInfo struct {
	MsgHead
	ByPayMode   uint8
	LRemainTime int32
}

func (msg *MsgS2CPayInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPayInfo) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SPayInfo struct {
	MsgHead
}

func (msg *MsgC2SPayInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPayInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SPayInfo(pcId uint32) MsgC2SPayInfo {
	msg := MsgC2SPayInfo{
		MsgHead: MsgHead{
			Protocol: C2SPayInfo,
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
