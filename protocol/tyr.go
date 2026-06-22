package protocol

type MsgS2CTyrError struct {
	MsgHead
	WTyrErrCode uint16
}

func (msg *MsgS2CTyrError) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CTyrError) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CTyrEntry struct {
	MsgHead
	ByMyUnit  uint8
	WEntryCnt uint16
	SzTyrInfo [0x80]byte
}

func (msg *MsgS2CTyrEntry) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CTyrEntry) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CTyrEntry(pcId uint32) MsgS2CTyrEntry {
	msg := MsgS2CTyrEntry{
		MsgHead: MsgHead{
			Protocol: S2CTyrEntry,
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

type MsgS2CTyrInfo struct {
	MsgHead
	ByTyrStat      uint8
	ByUnitIdx      uint8
	WEntryCnt      uint16
	WUnitMemberCnt uint16
	EntryList      [0x30]ZoneServerTyrEntry
}

func (msg *MsgS2CTyrInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CTyrInfo) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2STyrEntry struct {
	MsgHead
	DwNPCID        uint32
	ByUnitSelected uint8
}

func (msg *MsgC2STyrEntry) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2STyrEntry) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2STyrEntry(pcId uint32) MsgC2STyrEntry {
	msg := MsgC2STyrEntry{
		MsgHead: MsgHead{
			Protocol: C2STyrEntry,
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

type MsgC2STyrInfo struct {
	MsgHead
	DwNPCID   uint32
	ByUnitIdx uint8
}

func (msg *MsgC2STyrInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2STyrInfo) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2STyrJoin struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2STyrJoin) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2STyrJoin) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2STyrJoin(pcId uint32) MsgC2STyrJoin {
	msg := MsgC2STyrJoin{
		MsgHead: MsgHead{
			Protocol: C2STyrJoin,
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
