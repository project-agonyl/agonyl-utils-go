package protocol

type MsgC2SRoleMove struct {
	MsgHead
	X       byte
	Y       byte
	Unknown [0x2]byte
	Fix     byte
}

func (msg *MsgC2SRoleMove) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SRoleMove) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SRoleUpdatePos struct {
	MsgHead
	X       byte
	Y       byte
	Unknown [0x2]byte
	Fix     byte
}

func (msg *MsgC2SRoleUpdatePos) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SRoleUpdatePos) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SAskTransport struct {
	MsgHead
	Map     byte
	Unknown byte
}

func (msg *MsgC2SAskTransport) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskTransport) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SReturnToHere struct {
	MsgHead
	ID    uint32
	Index uint16
}

func (msg *MsgC2SReturnToHere) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SReturnToHere) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CRoleMoveResp struct {
	MsgHead
	Flag    byte
	X       byte
	Y       byte
	Unknown [0x2]byte
	Fix     byte
}

func (msg *MsgS2CRoleMoveResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CRoleMoveResp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CRoleMove struct {
	MsgHead
	UID  uint32
	Pos0 uint32
	Pos1 uint32
	Fix  byte
}

func (msg *MsgS2CRoleMove) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CRoleMove) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CAskTransportResp struct {
	MsgHead
	X       byte
	Y       byte
	Unknown uint16
}

func (msg *MsgS2CAskTransportResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAskTransportResp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CUpdateMap struct {
	MsgHead
	Map     byte
	Unknown [0x3]byte
}

func (msg *MsgS2CUpdateMap) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CUpdateMap) SetSize() {
	msg.Size = msg.GetSize()
}
