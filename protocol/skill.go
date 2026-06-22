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

func NewMsgC2SAskSkill(pcId uint32) MsgC2SAskSkill {
	msg := MsgC2SAskSkill{
		MsgHead: MsgHead{
			Protocol: C2SAskSkill,
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

func NewMsgC2SSkillSlotInfo(pcId uint32) MsgC2SSkillSlotInfo {
	msg := MsgC2SSkillSlotInfo{
		MsgHead: MsgHead{
			Protocol: C2SSkillSlotInfo,
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

func NewMsgS2CSkillRespPos(pcId uint32) MsgS2CSkillRespPos {
	msg := MsgS2CSkillRespPos{
		MsgHead: MsgHead{
			Protocol: S2CSkillRespPos,
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

func NewMsgS2CSkillResp(pcId uint32) MsgS2CSkillResp {
	msg := MsgS2CSkillResp{
		MsgHead: MsgHead{
			Protocol: S2CSkillResp,
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

func NewMsgS2CSkillBlessResp(pcId uint32) MsgS2CSkillBlessResp {
	msg := MsgS2CSkillBlessResp{
		MsgHead: MsgHead{
			Protocol: S2CSkillBlessResp,
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

func NewMsgS2CSkillStatusUpdate(pcId uint32) MsgS2CSkillStatusUpdate {
	msg := MsgS2CSkillStatusUpdate{
		MsgHead: MsgHead{
			Protocol: S2CSkillStatusUpdate,
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

func NewMsgS2CAnsSkillCaps(pcId uint32) MsgS2CAnsSkillCaps {
	msg := MsgS2CAnsSkillCaps{
		MsgHead: MsgHead{
			Protocol: S2CAnsSkillCaps,
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

type MsgS2CSkillDelayInfo struct {
	MsgHead
	ByDelay [0x20]uint8
}

func (msg *MsgS2CSkillDelayInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSkillDelayInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSkillDelayInfo(pcId uint32) MsgS2CSkillDelayInfo {
	msg := MsgS2CSkillDelayInfo{
		MsgHead: MsgHead{
			Protocol: S2CSkillDelayInfo,
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

type MsgS2CLearnSkill struct {
	MsgHead
	InvenIdx   uint8
	ScrollID   ZoneServerItemId
	SkillCode  uint8
	SkillLevel uint8
}

func (msg *MsgS2CLearnSkill) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CLearnSkill) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CLearnSkill(pcId uint32) MsgS2CLearnSkill {
	msg := MsgS2CLearnSkill{
		MsgHead: MsgHead{
			Protocol: S2CLearnSkill,
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

type MsgS2CErrSkill struct {
	MsgHead
	BySkillIndex uint8
	ByErrCode    uint8
}

func (msg *MsgS2CErrSkill) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CErrSkill) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CErrSkill(pcId uint32) MsgS2CErrSkill {
	msg := MsgS2CErrSkill{
		MsgHead: MsgHead{
			Protocol: S2CErrSkill,
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

type MsgS2CDoneSkill struct {
	MsgHead
	BySkillIndex uint8
	BySkillLevel uint8
	ByTargetType uint8
	DwCurCell    uint32
	DwBaseCell   uint32
	WHP          uint16
	WMP          uint16
	TargetCnt    uint8
	TargetInfo   [0x19]ZoneServerSkillDamageinfo
}

func (msg *MsgS2CDoneSkill) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CDoneSkill) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CDoneSkill(pcId uint32) MsgS2CDoneSkill {
	msg := MsgS2CDoneSkill{
		MsgHead: MsgHead{
			Protocol: S2CDoneSkill,
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

type MsgS2CGetSkill struct {
	MsgHead
	BySkillIndex uint8
	BySkillLevel uint8
	DwOPCID      uint32
	DwBaseCell   uint32
	WHP          uint16
	WMP          uint16
	BDie         uint8
	ByPetHP      uint8
}

func (msg *MsgS2CGetSkill) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CGetSkill) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CGetSkill(pcId uint32) MsgS2CGetSkill {
	msg := MsgS2CGetSkill{
		MsgHead: MsgHead{
			Protocol: S2CGetSkill,
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

type MsgS2CSeeSkill struct {
	MsgHead
	BySkillIndex uint8
	BySkillLevel uint8
	DwOPCID      uint32
	DwBaseCell   uint32
	ByTargetType uint8
	TargetCnt    uint8
	TargetInfo   [0x19]ZoneServerSkillDamageinfo
}

func (msg *MsgS2CSeeSkill) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSeeSkill) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSeeSkill(pcId uint32) MsgS2CSeeSkill {
	msg := MsgS2CSeeSkill{
		MsgHead: MsgHead{
			Protocol: S2CSeeSkill,
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

type MsgS2CCastSkill struct {
	MsgHead
	BySkillIndex uint8
	BySkillLevel uint8
	WHP          uint16
	WMP          uint16
	DwCurCell    uint32
}

func (msg *MsgS2CCastSkill) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CCastSkill) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CCastSkill(pcId uint32) MsgS2CCastSkill {
	msg := MsgS2CCastSkill{
		MsgHead: MsgHead{
			Protocol: S2CCastSkill,
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

type MsgS2CLastingSkill struct {
	MsgHead
	BySkillClass uint8
	BySkillIndex uint8
	BySkillLevel uint8
	DwOPCID      uint32
	BOnOff       uint8
	StatChange   ZoneServerSkillStatChange
}

func (msg *MsgS2CLastingSkill) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CLastingSkill) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CLastingSkill(pcId uint32) MsgS2CLastingSkill {
	msg := MsgS2CLastingSkill{
		MsgHead: MsgHead{
			Protocol: S2CLastingSkill,
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

type MsgS2CSeeLastingSkill struct {
	MsgHead
	BySkillClass uint8
	BySkillIndex uint8
	BySkillLevel uint8
	ByTargetType uint8
	DwCastPCID   uint32
	DwTargetID   uint32
	BOnOff       int32
}

func (msg *MsgS2CSeeLastingSkill) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSeeLastingSkill) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSeeLastingSkill(pcId uint32) MsgS2CSeeLastingSkill {
	msg := MsgS2CSeeLastingSkill{
		MsgHead: MsgHead{
			Protocol: S2CSeeLastingSkill,
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

type MsgC2SLearnSkill struct {
	MsgHead
	ScrollID ZoneServerItemId
	InvenIdx uint8
}

func (msg *MsgC2SLearnSkill) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SLearnSkill) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SLearnSkill(pcId uint32) MsgC2SLearnSkill {
	msg := MsgC2SLearnSkill{
		MsgHead: MsgHead{
			Protocol: C2SLearnSkill,
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
