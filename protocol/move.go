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

func NewMsgC2SRoleMove(pcId uint32) MsgC2SRoleMove {
	msg := MsgC2SRoleMove{
		MsgHead: MsgHead{
			Protocol: C2SAskMove,
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

func NewMsgC2SRoleUpdatePos(pcId uint32) MsgC2SRoleUpdatePos {
	msg := MsgC2SRoleUpdatePos{
		MsgHead: MsgHead{
			Protocol: C2SPcMove,
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

func NewMsgC2SAskTransport(pcId uint32) MsgC2SAskTransport {
	msg := MsgC2SAskTransport{
		MsgHead: MsgHead{
			Protocol: C2SWarp,
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

func NewMsgC2SReturnToHere(pcId uint32) MsgC2SReturnToHere {
	msg := MsgC2SReturnToHere{
		MsgHead: MsgHead{
			Protocol: C2SReturn2Here,
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

func NewMsgS2CRoleMoveResp(pcId uint32) MsgS2CRoleMoveResp {
	msg := MsgS2CRoleMoveResp{
		MsgHead: MsgHead{
			Protocol: S2CAnsMove,
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

func NewMsgS2CRoleMove(pcId uint32) MsgS2CRoleMove {
	msg := MsgS2CRoleMove{
		MsgHead: MsgHead{
			Protocol: S2CSeeMove,
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

func NewMsgS2CAskTransportResp(pcId uint32) MsgS2CAskTransportResp {
	msg := MsgS2CAskTransportResp{
		MsgHead: MsgHead{
			Protocol: S2CAskTransportResp,
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

func NewMsgS2CUpdateMap(pcId uint32) MsgS2CUpdateMap {
	msg := MsgS2CUpdateMap{
		MsgHead: MsgHead{
			Protocol: S2CUpdateMap,
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

type MsgS2CFixMove struct {
	MsgHead
	DwFixCell uint32
}

func (msg *MsgS2CFixMove) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CFixMove) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CFixMove(pcId uint32) MsgS2CFixMove {
	msg := MsgS2CFixMove{
		MsgHead: MsgHead{
			Protocol: S2CFixMove,
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

type MsgS2CSeeStop struct {
	MsgHead
	DwStopPCID uint32
	DwFixCell  uint32
}

func (msg *MsgS2CSeeStop) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSeeStop) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSeeStop(pcId uint32) MsgS2CSeeStop {
	msg := MsgS2CSeeStop{
		MsgHead: MsgHead{
			Protocol: S2CSeeStop,
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

type MsgS2CAnsMoveItemInStorage struct {
	MsgHead
	ByErrCode     uint8
	ItemID        ZoneServerItemId
	BySrcIndex    uint8
	ByTargetIndex uint8
	SwapitemID    ZoneServerItemId
}

func (msg *MsgS2CAnsMoveItemInStorage) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsMoveItemInStorage) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsMoveItemInStorage(pcId uint32) MsgS2CAnsMoveItemInStorage {
	msg := MsgS2CAnsMoveItemInStorage{
		MsgHead: MsgHead{
			Protocol: S2CAnsMoveItemInStorage,
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

type MsgC2SAskMoveItemInStorage struct {
	MsgHead
	DwNPCID     uint32
	ItemID      ZoneServerItemId
	ByFromIndex uint8
	ByToIndex   uint8
}

func (msg *MsgC2SAskMoveItemInStorage) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskMoveItemInStorage) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskMoveItemInStorage(pcId uint32) MsgC2SAskMoveItemInStorage {
	msg := MsgC2SAskMoveItemInStorage{
		MsgHead: MsgHead{
			Protocol: C2SAskMoveItemInStorage,
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
