package protocol

type MsgS2CStatAutoUpdate struct {
	MsgHead
	WRTime    uint16
	PetActive ZoneServerPetStat
	PetInven  [0x5]ZoneServerPetStat
}

func (msg *MsgS2CStatAutoUpdate) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CStatAutoUpdate) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CStatAutoUpdate(pcId uint32) MsgS2CStatAutoUpdate {
	msg := MsgS2CStatAutoUpdate{
		MsgHead: MsgHead{
			Protocol: S2CStatAutoUpdate,
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

type MsgS2CPearDecide struct {
	MsgHead
}

func (msg *MsgS2CPearDecide) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPearDecide) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPearDecide(pcId uint32) MsgS2CPearDecide {
	msg := MsgS2CPearDecide{
		MsgHead: MsgHead{
			Protocol: S2CPearDecide,
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

type MsgS2CSetPc2Stat struct {
	MsgHead
	PC2ndStat  ZoneServerPc2Stat
	ByStatCode uint8
}

func (msg *MsgS2CSetPc2Stat) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSetPc2Stat) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSetPc2Stat(pcId uint32) MsgS2CSetPc2Stat {
	msg := MsgS2CSetPc2Stat{
		MsgHead: MsgHead{
			Protocol: S2CSetPc2Stat,
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

type MsgS2CSvrTime struct {
	MsgHead
	Hour uint8
	Min  uint8
}

func (msg *MsgS2CSvrTime) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSvrTime) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSvrTime(pcId uint32) MsgS2CSvrTime {
	msg := MsgS2CSvrTime{
		MsgHead: MsgHead{
			Protocol: S2CSvrTime,
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

type MsgS2CUpdateKnight struct {
	MsgHead
	ByCode     uint8
	KnightInfo ZoneServerKnightInfo
}

func (msg *MsgS2CUpdateKnight) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CUpdateKnight) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CUpdateKnight(pcId uint32) MsgS2CUpdateKnight {
	msg := MsgS2CUpdateKnight{
		MsgHead: MsgHead{
			Protocol: S2CUpdateKnight,
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
