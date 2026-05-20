package protocol

const (
	CmdLA2LSSetAgentID      byte = 0xE0
	CmdLA2LSSetClientIP     byte = 0xE1
	CmdLA2LSExit            byte = 0xE2
	CmdLS2LAZoneAgentList   byte = 0xE1
	CmdZA2LSZoneAgentInfo   byte = 0xE0
	CmdZA2LSZoneServerCount byte = 0xE1
	CmdZA2LSPreWorldEntry   byte = 0xE2
	CmdZA2ASPing            byte = 0xE0
	CmdZA2ASLogin           byte = 0xE1
	CmdZA2ZSPing            byte = 0xE0
	CmdS2CZoneAgentList     byte = 0xE1
)

type MsgLA2LSSetAgentID struct {
	MsgHeadNoProtocol
	AgentID byte
}

func (msg *MsgLA2LSSetAgentID) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgLA2LSSetAgentID) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgLA2LSSetClientIP struct {
	MsgHeadNoProtocol
	IPAddress [0x10]byte
}

func (msg *MsgLA2LSSetClientIP) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgLA2LSSetClientIP) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgLA2LSExit struct {
	MsgHeadNoProtocol
}

func (msg *MsgLA2LSExit) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgLA2LSExit) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgZA2LSZoneServerCount struct {
	MsgHeadNoProtocol
	Count   uint32
	Unknown [0x2]byte
}

func (msg *MsgZA2LSZoneServerCount) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgZA2LSZoneServerCount) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgZA2LSPreWorldEntry struct {
	MsgHeadNoProtocol
	Unknown  byte
	Username [0x15]byte
	Day      [0x9]byte
	Time     [0x7]byte
}

func (msg *MsgZA2LSPreWorldEntry) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgZA2LSPreWorldEntry) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgLS2LAZoneAgentList struct {
	MsgHeadNoProtocol
	Count  byte
	Fix    [0x2]byte
	Name   [0x11]byte
	Online [0x51]byte
}

func (msg *MsgLS2LAZoneAgentList) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgLS2LAZoneAgentList) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CZoneAgentList struct {
	MsgHeadNoProtocol
	Unknown [0x15]byte
	Count   byte
	Fix     [0x2]byte
	Name    [0x11]byte
	Online  [0x51]byte
}

func (msg *MsgS2CZoneAgentList) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CZoneAgentList) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgZA2ASPing struct {
	MsgHeadNoProtocol
	Unknown byte
}

func (msg *MsgZA2ASPing) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgZA2ASPing) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgZA2ASLogin struct {
	MsgHeadNoProtocol
	Username [0x15]byte
	Unused   [0x15]byte
	IP       [0x10]byte
	Unknown  [0x4E]byte
}

func (msg *MsgZA2ASLogin) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgZA2ASLogin) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgZA2ZSPing struct {
	MsgHeadNoProtocol
	Unknown byte
}

func (msg *MsgZA2ZSPing) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgZA2ZSPing) SetSize() {
	msg.Size = msg.GetSize()
}
