package protocol

import (
	"encoding/binary"

	"github.com/cyberinferno/go-utils/utils"
)

type MsgC2SLogin struct {
	MsgHeadNoProtocol
	Username [0x15]byte
	Password [0x15]byte
}

func (msg *MsgC2SLogin) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgC2SLogin) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SLogin(username, password string) MsgC2SLogin {
	msg := MsgC2SLogin{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x01, Cmd: 0xE0},
	}
	copy(msg.Username[:], utils.MakeFixedLengthStringBytes(username, 0x15))
	copy(msg.Password[:], utils.MakeFixedLengthStringBytes(password, 0x15))
	msg.SetSize()
	return msg
}

type MsgC2SGateLogin struct {
	MsgHeadNoProtocol
	PcId     uint32
	Account  [0x15]byte
	Password [0x15]byte
}

func (msg *MsgC2SGateLogin) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgC2SGateLogin) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SGateLogin(pcId uint32, account string, password string) *MsgC2SGateLogin {
	msg := MsgC2SGateLogin{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x01, Cmd: 0xE2, PcId: pcId},
		PcId:              pcId,
	}

	copy(msg.Account[:], utils.MakeFixedLengthStringBytes(account, 0x15))
	copy(msg.Password[:], utils.MakeFixedLengthStringBytes(password, 0x15))
	msg.SetSize()
	return &msg
}

type MsgLs2ClSay struct {
	MsgHeadNoProtocol
	Type  byte
	Words [0x51]byte
}

func (msg *MsgLs2ClSay) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgLs2ClSay) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgLs2ClSay(words string) MsgLs2ClSay {
	msg := MsgLs2ClSay{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x01, Cmd: 0xE0},
		Type:              0x00,
	}
	copy(msg.Words[:], utils.MakeFixedLengthStringBytes(words, 0x51))
	msg.SetSize()
	return msg
}

type GateServerInfo struct {
	ServerID     byte
	ServerName   [0x11]byte
	ServerStatus [0x51]byte
}

type MsgLs2GateLogin struct {
	MsgHeadNoProtocol
	Account [0x15]byte
	Unknown [0x09]byte
}

func (msg *MsgLs2GateLogin) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgLs2GateLogin) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgLs2GateLogin(account string, pcId uint32) MsgLs2GateLogin {
	msg := MsgLs2GateLogin{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x01, Cmd: 0xE1, PcId: pcId},
	}
	copy(msg.Account[:], utils.MakeFixedLengthStringBytes(account, 0x15))
	msg.SetSize()
	return msg
}

type MsgS2CGateInfo struct {
	MsgHeadNoProtocol
	PcId   uint32
	ZaIP   [0x10]byte
	ZaPort uint32
}

func (msg *MsgS2CGateInfo) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgS2CGateInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CGateInfo(pcId uint32, zaIP string, zaPort uint32) MsgS2CGateInfo {
	msg := MsgS2CGateInfo{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x01, Cmd: 0xE2, PcId: pcId},
		PcId:              pcId,
		ZaPort:            zaPort,
	}
	copy(msg.ZaIP[:], utils.MakeFixedLengthStringBytes(zaIP, 0x10))
	msg.SetSize()
	return msg
}

type MsgLs2ZaDisconnect struct {
	MsgHeadNoProtocol
	Reason  byte
	Account [0x15]byte
	Unknown [0x10]byte
}

func (msg *MsgLs2ZaDisconnect) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgLs2ZaDisconnect) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgLs2ZaDisconnect(reason byte, account string, pcId uint32) MsgLs2ZaDisconnect {
	msg := MsgLs2ZaDisconnect{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x01, Cmd: 0xE3, PcId: pcId},
		Reason:            reason,
	}
	copy(msg.Account[:], utils.MakeFixedLengthStringBytes(account, 0x15))
	msg.SetSize()
	return msg
}

type MsgC2SSelectServer struct {
	MsgHeadNoProtocol
	ServerID byte
}

func (msg *MsgC2SSelectServer) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgC2SSelectServer) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SSelectServer(serverID byte) MsgC2SSelectServer {
	msg := MsgC2SSelectServer{
		MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x01, Cmd: 0xE1},
		ServerID:          serverID,
	}
	msg.SetSize()
	return msg
}

type MsgC2SCharacterLogout struct {
	MsgHead
}

func (msg *MsgC2SCharacterLogout) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgC2SCharacterLogout) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SCharacterLogout(pcId uint32) MsgC2SCharacterLogout {
	msg := MsgC2SCharacterLogout{
		MsgHead: MsgHead{Protocol: C2SCharacterLogout, MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x03, Cmd: 0xFF, PcId: pcId}},
	}
	msg.SetSize()
	return msg
}

type MsgC2SCharacterLogin struct {
	MsgHead
	CharacterName [0x15]byte
	ClientVersion uint32
}

func (msg *MsgC2SCharacterLogin) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgC2SCharacterLogin) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SCharacterLogin(pcId uint32, characterName string, clientVersion uint32) MsgC2SCharacterLogin {
	msg := MsgC2SCharacterLogin{
		MsgHead: MsgHead{
			Protocol: C2SCharacterLogin,
			MsgHeadNoProtocol: MsgHeadNoProtocol{
				Ctrl: 0x03,
				Cmd:  0xFF,
				PcId: pcId,
			},
		},
	}
	copy(msg.CharacterName[:], utils.MakeFixedLengthStringBytes(characterName, 0x15))
	msg.ClientVersion = clientVersion
	msg.SetSize()
	return msg
}

type MsgC2SWorldLogin struct {
	MsgHead
	CharacterName [0x15]byte
}

func (msg *MsgC2SWorldLogin) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgC2SWorldLogin) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SWorldLogin(pcId uint32, characterName string) MsgC2SWorldLogin {
	msg := MsgC2SWorldLogin{
		MsgHead: MsgHead{
			Protocol: C2SWorldLogin,
			MsgHeadNoProtocol: MsgHeadNoProtocol{
				Ctrl: 0x03,
				Cmd:  0xFF,
				PcId: pcId,
			},
		},
	}
	copy(msg.CharacterName[:], utils.MakeFixedLengthStringBytes(characterName, 0x15))
	msg.SetSize()
	return msg
}

type MsgS2CWorldLogin struct {
	MsgHead
	CharacterName   [0x15]byte
	Class           byte
	Level           uint16
	Exp             uint32
	MapNum          uint32
	XY              uint32
	SkillInfo       [0x1C]byte
	Town            byte
	Unknown1        byte
	Unknown2        uint16
	Woonz           uint32
	HPPot           uint32
	MPPot           uint32
	Lore            uint32
	RemainingPoints uint16
	Strength        uint16
	Intelligence    uint16
}

func (msg *MsgS2CWorldLogin) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgS2CWorldLogin) SetSize() {
	msg.Size = msg.GetSize()
}

type MsgS2CCharacterLogin struct {
	MsgHead
	CharacterName [0x15]byte
	Unknown       uint32
	MapNum        uint16
}

func (msg *MsgS2CCharacterLogin) GetSize() uint32 {
	return uint32(binary.Size(msg))
}

func (msg *MsgS2CCharacterLogin) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CCharacterLogin(pcId uint32, characterName string, unknown uint32, mapNum uint16) MsgS2CCharacterLogin {
	msg := MsgS2CCharacterLogin{
		MsgHead: MsgHead{Protocol: S2CCharacterLoginOk, MsgHeadNoProtocol: MsgHeadNoProtocol{Ctrl: 0x03, Cmd: 0xFF, PcId: pcId}},
	}
	copy(msg.CharacterName[:], utils.MakeFixedLengthStringBytes(characterName, 0x15))
	msg.Unknown = unknown
	msg.MapNum = mapNum
	msg.SetSize()
	return msg
}
