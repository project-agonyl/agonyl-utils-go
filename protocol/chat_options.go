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
