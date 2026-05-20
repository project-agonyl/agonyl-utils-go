package protocol

type MsgC2SPickItem struct {
	MsgHead
	ID   uint32
	Code uint32
	Cell uint32
}

func (msg *MsgC2SPickItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPickItem) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SDropItem struct {
	MsgHead
	ID    uint32
	Code  uint32
	Index byte
}

func (msg *MsgC2SDropItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SDropItem) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SInvItemMove struct {
	MsgHead
	ID           uint32
	Code         uint32
	CurrentIndex byte
	TargetIndex  byte
}

func (msg *MsgC2SInvItemMove) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SInvItemMove) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SWearUp struct {
	MsgHead
	ID   uint32
	Code uint32
}

func (msg *MsgC2SWearUp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SWearUp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SWearDown struct {
	MsgHead
	ID   uint32
	Code uint32
}

func (msg *MsgC2SWearDown) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SWearDown) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SBuyItem struct {
	MsgHead
	ID    uint32
	Code  uint16
	Count uint32
}

func (msg *MsgC2SBuyItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SBuyItem) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgC2SUseScroll struct {
	MsgHead
	Index byte
	ID    uint32
}

func (msg *MsgC2SUseScroll) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SUseScroll) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CWearUpResp struct {
	MsgHead
	Slot             byte
	ID               uint32
	Code             uint32
	Unknown0         byte
	Unknown1         uint16
	Unknown2         uint16
	Unknown3         uint16
	Unknown4         uint16
	MinAttack        uint16
	Unknown5         uint16
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
	Unknown7         uint16
}

func (msg *MsgS2CWearUpResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CWearUpResp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CWearDownResp struct {
	MsgHead
	ID               uint32
	Code             uint32
	Index            byte
	MinAttack        uint16
	Unknown1         uint16
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
	Unknown5         uint16
}

func (msg *MsgS2CWearDownResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CWearDownResp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CItemAppear struct {
	MsgHead
	ID      uint32
	Code    uint32
	Unknown [0x9]byte
	Cell    uint32
}

func (msg *MsgS2CItemAppear) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CItemAppear) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CAddItem struct {
	MsgHead
	ID      uint32
	Code    uint32
	Attr    uint32
	Unknown uint32
	Index   byte
}

func (msg *MsgS2CAddItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAddItem) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CDropItemResp struct {
	MsgHead
	ID    uint32
	Code  uint32
	Index byte
	Cell  uint32
}

func (msg *MsgS2CDropItemResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CDropItemResp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CAnsInvItemMove struct {
	MsgHead
	Index0 byte
	ID0    uint32
	Code0  uint32
	Index1 byte
	ID1    uint32
	Code1  uint32
}

func (msg *MsgS2CAnsInvItemMove) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsInvItemMove) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CBuyItemResp struct {
	MsgHead
	Index      byte
	ID         uint32
	Code       uint32
	Count      uint32
	Attr       uint32
	Money      uint32
	PointCount uint32
}

func (msg *MsgS2CBuyItemResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CBuyItemResp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CUpdateItem struct {
	MsgHead
	Index byte
	ID    uint32
	Count byte
}

func (msg *MsgS2CUpdateItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CUpdateItem) SetSize() {
	msg.Size = msg.GetSize()
}
