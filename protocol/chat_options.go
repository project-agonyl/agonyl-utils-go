package protocol

type MsgC2SChatOpt struct {
	MsgHead
	Opt0     byte
	Unknown0 byte
	Opt1     byte
	Unknown1 byte
	Opt2     byte
	Unknown2 byte
}

func (msg *MsgC2SChatOpt) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SChatOpt) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SChatOpt(pcId uint32) MsgC2SChatOpt {
	msg := MsgC2SChatOpt{
		MsgHead: MsgHead{
			Protocol: C2SChatWindowOpt,
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

type MsgS2CChatOpt struct {
	MsgHead
	Opt0     byte
	Unknown0 byte
	Opt1     byte
	Unknown1 byte
	Opt2     byte
	Unknown2 byte
}

func (msg *MsgS2CChatOpt) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CChatOpt) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CChatOpt(pcId uint32) MsgS2CChatOpt {
	msg := MsgS2CChatOpt{
		MsgHead: MsgHead{
			Protocol: S2CChatWindowOpt,
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
