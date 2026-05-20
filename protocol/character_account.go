package protocol

type MsgZA2ASCreateCharacter struct {
	MsgHead
	Career byte
	Nation byte
	Name   [0x15]byte
}

func (msg *MsgZA2ASCreateCharacter) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgZA2ASCreateCharacter) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgZA2ASDeleteCharacter struct {
	MsgHead
	Name [0x15]byte
}

func (msg *MsgZA2ASDeleteCharacter) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgZA2ASDeleteCharacter) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgAS2ZACreateCharacter struct {
	MsgHead
	Career   byte
	Name     [0x15]byte
	WearList [0xA]AccountItem
}

func (msg *MsgAS2ZACreateCharacter) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgAS2ZACreateCharacter) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgAS2ZADeleteCharacter struct {
	MsgHead
	Name [0x15]byte
}

func (msg *MsgAS2ZADeleteCharacter) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgAS2ZADeleteCharacter) SetSize() {
	msg.Size = msg.GetSize()
}

type RoleBrief struct {
	Name      [0x15]byte
	ClassType byte
	Town      byte
	Unknown   byte
	Level     uint32
	WearList  [0xA]AccountItem
}

type MsgAS2ZALoginResp struct {
	MsgHead
	RoleBrief [0x5]RoleBrief
}

func (msg *MsgAS2ZALoginResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgAS2ZALoginResp) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgZA2ASSelectRole struct {
	MsgHead
	Name    [0x15]byte
	Unknown uint32
}

func (msg *MsgZA2ASSelectRole) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgZA2ASSelectRole) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgAS2ZAConfirmRole struct {
	MsgHead
	Name     [0x15]byte
	UID      uint32
	MapIndex uint16
}

func (msg *MsgAS2ZAConfirmRole) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgAS2ZAConfirmRole) SetSize() {
	msg.Size = msg.GetSize()
}
