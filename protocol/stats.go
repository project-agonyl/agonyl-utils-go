package protocol

type MsgC2SPointInc struct {
	MsgHead
	Strength uint8
	Magic    uint8
	Dex      uint8
	Vit      uint8
	Mana     uint8
}

func (msg *MsgC2SPointInc) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPointInc) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SAskHeal struct {
	MsgHead
	ID uint32
}

func (msg *MsgC2SAskHeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskHeal) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SPointDec struct {
	MsgHead
	Strength uint8
	Magic    uint8
	Dex      uint8
	Vit      uint8
	Mana     uint8
}

func (msg *MsgC2SPointDec) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPointDec) SetSize() {
	msg.Size = msg.GetSize()
}

type CharacterPointStat struct {
	Point            uint16
	Strength         uint16
	Magic            uint16
	Dex              uint16
	Vit              uint16
	Mana             uint16
	MaxHPStore       uint32
	MaxMPStore       uint32
	HP               uint16
	MP               uint16
	MinAttack        uint16
	MinMagicAttack   uint16
	Defense          uint16
	FireAttack       uint16
	FireDefense      uint16
	IceAttack        uint16
	IceDefense       uint16
	LightningAttack  uint16
	LightningDefense uint16
	MaxHPBar         uint16
	MaxMPBar         uint16
	MaxAttack        uint16
	Unknown          uint16
	MaxMagicAttack   uint16
}

type MsgS2CPointIncResp struct {
	MsgHead
	CharacterPointStat
}

func (msg *MsgS2CPointIncResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPointIncResp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CUpdateExp struct {
	MsgHead
	Exp     uint32
	Unknown uint16
}

func (msg *MsgS2CUpdateExp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CUpdateExp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CAskHeal struct {
	MsgHead
	HP uint16
	MP uint16
}

func (msg *MsgS2CAskHeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAskHeal) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CPointDecResp struct {
	MsgHead
	Unknown0 byte
	CharacterPointStat
	LorePoint uint32
}

func (msg *MsgS2CPointDecResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPointDecResp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CPayMoney struct {
	MsgHead
	PayMoney uint32
	Unknown  byte
}

func (msg *MsgS2CPayMoney) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPayMoney) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SRecovery struct {
	MsgHead
	Fix  byte
	Type byte
}

func (msg *MsgC2SRecovery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SRecovery) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CRecovery struct {
	MsgHead
	Type  byte
	Fix   byte
	Bar   uint32
	Store uint32
}

func (msg *MsgS2CRecovery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CRecovery) SetSize() {
	msg.Size = msg.GetSize()
}
