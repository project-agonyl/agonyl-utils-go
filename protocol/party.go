package protocol

type MsgS2CAskParty struct {
	MsgHead
	DwAskPCID uint32
}

func (msg *MsgS2CAskParty) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAskParty) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAskParty(pcId uint32) MsgS2CAskParty {
	msg := MsgS2CAskParty{
		MsgHead: MsgHead{
			Protocol: S2CAskParty,
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

type MsgS2CAnsParty struct {
	MsgHead
	Result uint8
}

func (msg *MsgS2CAnsParty) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsParty) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsParty(pcId uint32) MsgS2CAnsParty {
	msg := MsgS2CAnsParty{
		MsgHead: MsgHead{
			Protocol: S2CAnsParty,
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

type MsgS2CPartyIn struct {
	MsgHead
	Member ZoneServerPartyMember
}

func (msg *MsgS2CPartyIn) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPartyIn) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPartyIn(pcId uint32) MsgS2CPartyIn {
	msg := MsgS2CPartyIn{
		MsgHead: MsgHead{
			Protocol: S2CPartyIn,
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

type MsgS2CPartyList struct {
	MsgHead
	Member [0x4]ZoneServerPartyMember
}

func (msg *MsgS2CPartyList) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPartyList) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPartyList(pcId uint32) MsgS2CPartyList {
	msg := MsgS2CPartyList{
		MsgHead: MsgHead{
			Protocol: S2CPartyList,
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

type MsgS2CPartyOut struct {
	MsgHead
	DwMemberID uint32
}

func (msg *MsgS2CPartyOut) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPartyOut) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPartyOut(pcId uint32) MsgS2CPartyOut {
	msg := MsgS2CPartyOut{
		MsgHead: MsgHead{
			Protocol: S2CPartyOut,
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

type MsgS2CPartyWhere struct {
	MsgHead
	DwMemberID  uint32
	WMapIndex   uint16
	DwCellIndex uint32
}

func (msg *MsgS2CPartyWhere) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPartyWhere) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPartyWhere(pcId uint32) MsgS2CPartyWhere {
	msg := MsgS2CPartyWhere{
		MsgHead: MsgHead{
			Protocol: S2CPartyWhere,
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

type MsgS2CPartyHp struct {
	MsgHead
	DwMemberID uint32
	WHP        uint16
}

func (msg *MsgS2CPartyHp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPartyHp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPartyHp(pcId uint32) MsgS2CPartyHp {
	msg := MsgS2CPartyHp{
		MsgHead: MsgHead{
			Protocol: S2CPartyHp,
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

type MsgC2SAskParty struct {
	MsgHead
	DwOPCID uint32
}

func (msg *MsgC2SAskParty) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskParty) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskParty(pcId uint32) MsgC2SAskParty {
	msg := MsgC2SAskParty{
		MsgHead: MsgHead{
			Protocol: C2SAskParty,
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

type MsgC2SAnsParty struct {
	MsgHead
	DwAskPCID uint32
	BAccept   uint8
}

func (msg *MsgC2SAnsParty) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAnsParty) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAnsParty(pcId uint32) MsgC2SAnsParty {
	msg := MsgC2SAnsParty{
		MsgHead: MsgHead{
			Protocol: C2SAnsParty,
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

type MsgC2SOutParty struct {
	MsgHead
}

func (msg *MsgC2SOutParty) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SOutParty) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SOutParty(pcId uint32) MsgC2SOutParty {
	msg := MsgC2SOutParty{
		MsgHead: MsgHead{
			Protocol: C2SOutParty,
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
