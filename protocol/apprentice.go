package protocol

type MsgS2CAskApprenticeIn struct {
	MsgHead
	DwRequestPCID uint32
}

func (msg *MsgS2CAskApprenticeIn) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAskApprenticeIn) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAskApprenticeIn(pcId uint32) MsgS2CAskApprenticeIn {
	msg := MsgS2CAskApprenticeIn{
		MsgHead: MsgHead{
			Protocol: S2CAskApprenticeIn,
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

type MsgS2CAnsApprenticeIn struct {
	MsgHead
	BResult          uint8
	StApprenticeInfo ZoneServerPartyMember
}

func (msg *MsgS2CAnsApprenticeIn) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsApprenticeIn) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsApprenticeIn(pcId uint32) MsgS2CAnsApprenticeIn {
	msg := MsgS2CAnsApprenticeIn{
		MsgHead: MsgHead{
			Protocol: S2CAnsApprenticeIn,
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

type MsgS2CApprenticeUpdate struct {
	MsgHead
	WMapIndex   uint16
	DwCellIndex uint32
}

func (msg *MsgS2CApprenticeUpdate) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CApprenticeUpdate) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CApprenticeUpdate(pcId uint32) MsgS2CApprenticeUpdate {
	msg := MsgS2CApprenticeUpdate{
		MsgHead: MsgHead{
			Protocol: S2CApprenticeUpdate,
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

type MsgS2CApprenticeHp struct {
	MsgHead
	WHP uint16
}

func (msg *MsgS2CApprenticeHp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CApprenticeHp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CApprenticeHp(pcId uint32) MsgS2CApprenticeHp {
	msg := MsgS2CApprenticeHp{
		MsgHead: MsgHead{
			Protocol: S2CApprenticeHp,
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

type MsgS2CAskApprenticeOut struct {
	MsgHead
}

func (msg *MsgS2CAskApprenticeOut) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAskApprenticeOut) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAskApprenticeOut(pcId uint32) MsgS2CAskApprenticeOut {
	msg := MsgS2CAskApprenticeOut{
		MsgHead: MsgHead{
			Protocol: S2CAskApprenticeOut,
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

type MsgC2SAskApprenticeIn struct {
	MsgHead
	DwRequestedPCID uint32
}

func (msg *MsgC2SAskApprenticeIn) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskApprenticeIn) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskApprenticeIn(pcId uint32) MsgC2SAskApprenticeIn {
	msg := MsgC2SAskApprenticeIn{
		MsgHead: MsgHead{
			Protocol: C2SAskApprenticeIn,
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

type MsgC2SAnsApprenticeIn struct {
	MsgHead
	DwRequestPCID uint32
	BAccepted     uint8
}

func (msg *MsgC2SAnsApprenticeIn) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAnsApprenticeIn) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAnsApprenticeIn(pcId uint32) MsgC2SAnsApprenticeIn {
	msg := MsgC2SAnsApprenticeIn{
		MsgHead: MsgHead{
			Protocol: C2SAnsApprenticeIn,
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

type MsgC2SAskApprenticeOut struct {
	MsgHead
}

func (msg *MsgC2SAskApprenticeOut) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskApprenticeOut) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskApprenticeOut(pcId uint32) MsgC2SAskApprenticeOut {
	msg := MsgC2SAskApprenticeOut{
		MsgHead: MsgHead{
			Protocol: C2SAskApprenticeOut,
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
