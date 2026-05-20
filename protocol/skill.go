package protocol

type MsgC2SAskSkill struct {
	MsgHead
	Unknown0 byte
	Unknown1 byte
	Unknown2 byte
	Unknown3 byte
	Code     byte
	Type     byte
	Param    uint32
}

func (msg *MsgC2SAskSkill) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskSkill) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SSkillSlotInfo struct {
	MsgHead
	SlotInfo [0xD]byte
}

func (msg *MsgC2SSkillSlotInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SSkillSlotInfo) SetSize() {
	msg.Size = msg.GetSize()
}

type SkillRespTarget struct {
	ID      uint32
	Status  uint32
	Status2 int16
}

type MsgS2CSkillRespPos struct {
	MsgHead
	Unknown0 byte
	Unknown1 byte
}

func (msg *MsgS2CSkillRespPos) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSkillRespPos) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CSkillResp struct {
	MsgHead
	Code          byte
	Unknown0      byte
	Flag          byte
	CharacterCell uint32
	TargetCell    uint32
	HP            uint16
	MP            uint16
	TargetCount   byte
}

func (msg *MsgS2CSkillResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSkillResp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CSkillBlessResp struct {
	MsgHead
	Code      byte
	Unknown0  byte
	TargetUID uint32
	Unknown1  uint32
	HP        uint16
	MP        uint16
	Unknown2  uint16
}

func (msg *MsgS2CSkillBlessResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSkillBlessResp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CSkillStatusUpdate struct {
	MsgHead
	Code byte
	Flag byte
	HP   uint16
	MP   uint16
	Cell uint32
}

func (msg *MsgS2CSkillStatusUpdate) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSkillStatusUpdate) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CAnsSkillCaps struct {
	MsgHead
	Unknown0  byte
	Code      byte
	Unknown1  byte
	TargetUID uint32
	Flag      byte
	Cap       [0xB]uint16
	Unknown3  [0x6]byte
}

func (msg *MsgS2CAnsSkillCaps) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsSkillCaps) SetSize() {
	msg.Size = msg.GetSize()
}
