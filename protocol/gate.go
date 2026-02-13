package protocol

import (
	"encoding/binary"

	"github.com/cyberinferno/go-utils/utils"
)

type MsgGate2LsConnect struct {
	MsgHeadNoProtocol
	ServerId  byte
	AgentId   byte
	IpAddress [0x10]byte
	Port      uint32
	Name      [0x11]byte
}

func (msg *MsgGate2LsConnect) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgGate2LsConnect) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgGate2LsConnect(serverId byte, agentId byte, ipAddress string, port uint32, name string) MsgGate2LsConnect {
	msg := MsgGate2LsConnect{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x02, Cmd: 0xE0},
		ServerId:          serverId,
		AgentId:           agentId,
		Port:              port,
	}
	copy(msg.IpAddress[:], utils.MakeFixedLengthStringBytes(ipAddress, 0x10))
	copy(msg.Name[:], utils.MakeFixedLengthStringBytes(name, 0x11))
	msg.SetSize()
	return msg
}

type MsgGate2LsAccLogout struct {
	MsgHeadNoProtocol
	Reason     byte
	Account    [0x15]byte
	LogoutDate [0x09]byte
	LogoutTime [0x07]byte
}

func (msg *MsgGate2LsAccLogout) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgGate2LsAccLogout) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgGate2LsAccLogout(reason byte, account string) MsgGate2LsAccLogout {
	msg := MsgGate2LsAccLogout{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x02, Cmd: 0xE2},
		Reason:            reason,
	}
	copy(msg.Account[:], utils.MakeFixedLengthStringBytes(account, 0x15))
	msg.SetSize()
	return msg
}

type MsgGate2LsPreparedAccLogin struct {
	MsgHeadNoProtocol
	Account [0x15]byte
}

func (msg *MsgGate2LsPreparedAccLogin) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgGate2LsPreparedAccLogin) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgGate2LsPreparedAccLogin(account string) MsgGate2LsPreparedAccLogin {
	msg := MsgGate2LsPreparedAccLogin{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x02, Cmd: 0xE3},
	}
	copy(msg.Account[:], utils.MakeFixedLengthStringBytes(account, 0x15))
	msg.SetSize()
	return msg
}

type MsgGate2ZsConnect struct {
	MsgHeadNoProtocol
	AgentID byte
}

func (msg *MsgGate2ZsConnect) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgGate2ZsConnect) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgGate2ZsConnect(agentID byte) MsgGate2ZsConnect {
	msg := MsgGate2ZsConnect{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x01, Cmd: 0xE0},
		AgentID:           agentID,
	}
	msg.SetSize()
	return msg
}

type MsgZa2ZsAccLogout struct {
	MsgHeadNoProtocol
	Reason byte
}

func (msg *MsgZa2ZsAccLogout) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgZa2ZsAccLogout) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgZa2ZsAccLogout(pcId uint32, reason byte) *MsgZa2ZsAccLogout {
	msg := MsgZa2ZsAccLogout{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x01, Cmd: 0xE2, PcId: pcId},
		Reason:            reason,
	}
	msg.SetSize()
	return &msg
}
