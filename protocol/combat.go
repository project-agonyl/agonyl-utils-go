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
