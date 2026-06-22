package protocol

type MsgC2SQuestInfo struct {
	MsgHead
	ItemID     ZoneServerItemId
	InvenIndex uint8
}

func (msg *MsgC2SQuestInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SQuestInfo) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CQuestInfo struct {
	MsgHead
	Title [0x40]byte
	Desc  [0x180]byte
}

func (msg *MsgS2CQuestInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CQuestInfo) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SQuestStart struct {
	MsgHead
	ItemID     ZoneServerItemId
	InvenIndex uint8
}

func (msg *MsgC2SQuestStart) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SQuestStart) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CQuestStart struct {
	MsgHead
	ItemID  ZoneServerItemId
	ErrCode uint8
}

func (msg *MsgS2CQuestStart) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CQuestStart) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CQuestContinue struct {
	MsgHead
	Title [0x40]byte
	Desc  [0x180]byte
}

func (msg *MsgS2CQuestContinue) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CQuestContinue) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SQuestCancel struct {
	MsgHead
}

func (msg *MsgC2SQuestCancel) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SQuestCancel) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CQuestCancel struct {
	MsgHead
	ErrCode uint8
}

func (msg *MsgS2CQuestCancel) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CQuestCancel) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CQuestState struct {
	MsgHead
	SzState [0xc0]byte
}

func (msg *MsgS2CQuestState) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CQuestState) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CQuestOk struct {
	MsgHead
	Lore int32
}

func (msg *MsgS2CQuestOk) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CQuestOk) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CPartyQuest struct {
	MsgHead
	ByPartyQuestIndex uint8
	ByPartyQuestStep  uint8
	ByResultCode      uint8
}

func (msg *MsgS2CPartyQuest) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPartyQuest) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPartyQuest(pcId uint32) MsgS2CPartyQuest {
	msg := MsgS2CPartyQuest{
		MsgHead: MsgHead{
			Protocol: S2CPartyQuest,
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

type MsgS2CPartyQuestTimer struct {
	MsgHead
	ByLimitTime uint8
	ByCode      uint8
}

func (msg *MsgS2CPartyQuestTimer) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPartyQuestTimer) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPartyQuestTimer(pcId uint32) MsgS2CPartyQuestTimer {
	msg := MsgS2CPartyQuestTimer{
		MsgHead: MsgHead{
			Protocol: S2CPartyQuestTimer,
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

type MsgC2SPartyQuest struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SPartyQuest) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPartyQuest) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SPartyQuest(pcId uint32) MsgC2SPartyQuest {
	msg := MsgC2SPartyQuest{
		MsgHead: MsgHead{
			Protocol: C2SPartyQuest,
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
