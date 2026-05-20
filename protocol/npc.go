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
