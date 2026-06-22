package protocol

type MsgC2SAttackMonster struct {
	MsgHead
	Unknown0 uint32
	Fix      byte
	ID       uint32
	Unknown1 byte
}

func (msg *MsgC2SAttackMonster) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAttackMonster) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAttackMonster(pcId uint32) MsgC2SAttackMonster {
	msg := MsgC2SAttackMonster{
		MsgHead: MsgHead{
			Protocol: C2SAskAttack,
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

type MsgS2CAttackMonsterResp struct {
	MsgHead
	Unknown0    byte
	CellChar    uint32
	Status      byte
	SerialNo    uint16
	ID          uint16
	CellMon     uint32
	Damage      uint16
	Unknown1    uint16
	CharacterHP uint16
	CharacterMP uint16
}

func (msg *MsgS2CAttackMonsterResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAttackMonsterResp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAttackMonsterResp(pcId uint32) MsgS2CAttackMonsterResp {
	msg := MsgS2CAttackMonsterResp{
		MsgHead: MsgHead{
			Protocol: S2CAttackMonsterResp,
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

type MsgS2CSeeAttack struct {
	MsgHead
	DwPCID         uint32
	DwPCCell       uint32
	ByTargetType   uint8
	DwTargetID     uint32
	DwTargetCell   uint32
	ByCurrentStep  uint8
	ByTargetStatus uint8
	BFinish        uint8
}

func (msg *MsgS2CSeeAttack) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSeeAttack) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSeeAttack(pcId uint32) MsgS2CSeeAttack {
	msg := MsgS2CSeeAttack{
		MsgHead: MsgHead{
			Protocol: S2CSeeAttack,
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

type MsgS2CUpdatePkCount struct {
	MsgHead
	WRTime    uint16
	ByPKCount uint8
}

func (msg *MsgS2CUpdatePkCount) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CUpdatePkCount) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CUpdatePkCount(pcId uint32) MsgS2CUpdatePkCount {
	msg := MsgS2CUpdatePkCount{
		MsgHead: MsgHead{
			Protocol: S2CUpdatePkCount,
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

type MsgS2CPcAttack struct {
	MsgHead
	DwAttackPCID   uint32
	DwAttackPCCell uint32
	WPCHP          uint16
	WPCMP          uint16
	ByCurStep      uint8
	BDie           uint8
	ByPetHP        uint8
}

func (msg *MsgS2CPcAttack) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPcAttack) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPcAttack(pcId uint32) MsgS2CPcAttack {
	msg := MsgS2CPcAttack{
		MsgHead: MsgHead{
			Protocol: S2CPcAttack,
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
