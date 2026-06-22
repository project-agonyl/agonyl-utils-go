package protocol

type MsgC2SNpcHover struct {
	MsgHead
	Code uint16
}

func (msg *MsgC2SNpcHover) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SNpcHover) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SNpcHover(pcId uint32) MsgC2SNpcHover {
	msg := MsgC2SNpcHover{
		MsgHead: MsgHead{
			Protocol: C2SAskNpcFavor,
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

type MsgS2CRoleInfo struct {
	MsgHead
	UID         uint32
	RoleName    [0x15]byte
	Career      byte
	Unknown0    byte
	State       byte
	X           byte
	Y           byte
	Constant0   uint16
	Nation      byte
	Rank        byte
	KnightIndex uint16
	Head        uint32
	Armour      uint32
	Trousers    uint32
	Groove      uint32
	Boot        uint32
	Neck        uint32
	Weapon      uint32
	Pet         uint32
	Skill       [0x1E]byte
	Unknown     [0x2B]byte
}

func (msg *MsgS2CRoleInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CRoleInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CRoleInfo(pcId uint32) MsgS2CRoleInfo {
	msg := MsgS2CRoleInfo{
		MsgHead: MsgHead{
			Protocol: S2CPcAppear,
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

type MsgS2CPcDisappear struct {
	MsgHead
	UID uint32
}

func (msg *MsgS2CPcDisappear) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPcDisappear) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPcDisappear(pcId uint32) MsgS2CPcDisappear {
	msg := MsgS2CPcDisappear{
		MsgHead: MsgHead{
			Protocol: S2CPcDisappear,
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

type MsgS2CNpcInfo struct {
	MsgHead
	Code     uint16
	SerialNo uint16
	ID       uint16
	HP       uint32
	Fix      uint32
	X        byte
	Y        byte
	Unknown0 uint16
	A        uint16
	B        uint16
	C        uint16
	Unknown2 [0x1E]byte
	Unknown3 [0xA]byte
}

func (msg *MsgS2CNpcInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CNpcInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CNpcInfo(pcId uint32) MsgS2CNpcInfo {
	msg := MsgS2CNpcInfo{
		MsgHead: MsgHead{
			Protocol: S2CNpcAppear,
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

type MsgS2CNpcMove struct {
	MsgHead
	UID uint32
	Pos [0x9]uint32
}

func (msg *MsgS2CNpcMove) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CNpcMove) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CNpcMove(pcId uint32) MsgS2CNpcMove {
	msg := MsgS2CNpcMove{
		MsgHead: MsgHead{
			Protocol: S2CNpcMove,
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

type MsgS2CMonsterAttack struct {
	MsgHead
	ID         uint32
	MonsterPos uint32
	Unknown0   uint32
	PCPos      uint32
	HP         uint16
	MP         uint16
	Unknown1   [0x7]byte
}

func (msg *MsgS2CMonsterAttack) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CMonsterAttack) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CMonsterAttack(pcId uint32) MsgS2CMonsterAttack {
	msg := MsgS2CMonsterAttack{
		MsgHead: MsgHead{
			Protocol: S2CMonsterAttack,
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

type MsgS2CNpcHoverResp struct {
	MsgHead
	Code    uint16
	Unknown [0x2]byte
}

func (msg *MsgS2CNpcHoverResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CNpcHoverResp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CNpcHoverResp(pcId uint32) MsgS2CNpcHoverResp {
	msg := MsgS2CNpcHoverResp{
		MsgHead: MsgHead{
			Protocol: S2CNpcHoverResp,
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

type MsgS2CNpcDisappear struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgS2CNpcDisappear) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CNpcDisappear) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CNpcDisappear(pcId uint32) MsgS2CNpcDisappear {
	msg := MsgS2CNpcDisappear{
		MsgHead: MsgHead{
			Protocol: S2CNpcDisappear,
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

type MsgS2CNpcState struct {
	MsgHead
	DwNPCID    uint32
	ByNPCState uint8
	WReserved  uint16
}

func (msg *MsgS2CNpcState) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CNpcState) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CNpcState(pcId uint32) MsgS2CNpcState {
	msg := MsgS2CNpcState{
		MsgHead: MsgHead{
			Protocol: S2CNpcState,
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

type MsgS2CSeeNpcAttack struct {
	MsgHead
	DwNPCID     uint32
	DwNPCCell   uint32
	DwPCID      uint32
	DwPCCell    uint32
	WAttackType uint16
	BPCDie      int32
}

func (msg *MsgS2CSeeNpcAttack) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSeeNpcAttack) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSeeNpcAttack(pcId uint32) MsgS2CSeeNpcAttack {
	msg := MsgS2CSeeNpcAttack{
		MsgHead: MsgHead{
			Protocol: S2CSeeNpcAttack,
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

type MsgS2CSeeNpcSkillAttack struct {
	MsgHead
	DwNPCID      uint32
	DwNPCCell    uint32
	DwBaseCell   uint32
	ByAttackType uint8
	TargetCnt    uint8
	TargetInfo   [0x19]ZoneServerNpcSkillDamageinfo
}

func (msg *MsgS2CSeeNpcSkillAttack) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSeeNpcSkillAttack) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSeeNpcSkillAttack(pcId uint32) MsgS2CSeeNpcSkillAttack {
	msg := MsgS2CSeeNpcSkillAttack{
		MsgHead: MsgHead{
			Protocol: S2CSeeNpcSkillAttack,
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

type MsgS2CSubmapInfo struct {
	MsgHead
	SzSubmapName   [0x14]byte
	WSubmapIndex   uint16
	PackedBits0    uint8 // bitfields: bOpen:1, bPublic:1, bKeyword:1, Reserved:5
	ConditionLevel uint8
	ConditionType  uint8
}

func (msg *MsgS2CSubmapInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSubmapInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSubmapInfo(pcId uint32) MsgS2CSubmapInfo {
	msg := MsgS2CSubmapInfo{
		MsgHead: MsgHead{
			Protocol: S2CSubmapInfo,
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

type MsgS2CEnter struct {
	MsgHead
	ByErrorCode  uint8
	WSubmapIndex uint16
	DwCellIndex  uint32
}

func (msg *MsgS2CEnter) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CEnter) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CEnter(pcId uint32) MsgS2CEnter {
	msg := MsgS2CEnter{
		MsgHead: MsgHead{
			Protocol: S2CEnter,
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

type MsgS2COtherUpdate struct {
	MsgHead
	DwOPCID    uint32
	ByCaoLevel uint8
	ByPose     uint8
}

func (msg *MsgS2COtherUpdate) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2COtherUpdate) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2COtherUpdate(pcId uint32) MsgS2COtherUpdate {
	msg := MsgS2COtherUpdate{
		MsgHead: MsgHead{
			Protocol: S2CPcUpdate,
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

type MsgS2CFortifier struct {
	MsgHead
	ByGreen  uint8
	ByBlue   uint8
	ByRed    uint8
	ByYellow uint8
	ByBlack  uint8
}

func (msg *MsgS2CFortifier) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CFortifier) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CFortifier(pcId uint32) MsgS2CFortifier {
	msg := MsgS2CFortifier{
		MsgHead: MsgHead{
			Protocol: S2CFortifier,
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

type MsgS2CAgitEnter struct {
	MsgHead
	ByErrorCode  uint8
	WSubmapIndex uint16
	DwCellIndex  uint32
}

func (msg *MsgS2CAgitEnter) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAgitEnter) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAgitEnter(pcId uint32) MsgS2CAgitEnter {
	msg := MsgS2CAgitEnter{
		MsgHead: MsgHead{
			Protocol: S2CAgitEnter,
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

type MsgC2SSubmapInfo struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SSubmapInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SSubmapInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SSubmapInfo(pcId uint32) MsgC2SSubmapInfo {
	msg := MsgC2SSubmapInfo{
		MsgHead: MsgHead{
			Protocol: C2SSubmapInfo,
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

type MsgC2SEnter struct {
	MsgHead
	DwNPCID uint32
	Keyword [0x40]byte
}

func (msg *MsgC2SEnter) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SEnter) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SEnter(pcId uint32) MsgC2SEnter {
	msg := MsgC2SEnter{
		MsgHead: MsgHead{
			Protocol: C2SEnter,
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

type MsgC2SObjectNpc struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SObjectNpc) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SObjectNpc) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SObjectNpc(pcId uint32) MsgC2SObjectNpc {
	msg := MsgC2SObjectNpc{
		MsgHead: MsgHead{
			Protocol: C2SObjectNpc,
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

type MsgC2SAgitEnter struct {
	MsgHead
	DwNPCID    uint32
	BySelfAgit uint8
	WAgitID    uint16
}

func (msg *MsgC2SAgitEnter) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAgitEnter) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAgitEnter(pcId uint32) MsgC2SAgitEnter {
	msg := MsgC2SAgitEnter{
		MsgHead: MsgHead{
			Protocol: C2SAgitEnter,
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
