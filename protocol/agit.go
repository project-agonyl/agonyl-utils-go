package protocol

type MsgS2CAgitInfo struct {
	MsgHead
	ByAgitCount uint8
	ArrAgitInfo [0x19]ZoneServerAgitInfo
}

func (msg *MsgS2CAgitInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAgitInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAgitInfo(pcId uint32) MsgS2CAgitInfo {
	msg := MsgS2CAgitInfo{
		MsgHead: MsgHead{
			Protocol: S2CAgitInfo,
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

type MsgS2CAuctionInfo struct {
	MsgHead
	ByAuctionCount uint8
	ArrAuctionInfo [0x14]ZoneServerAuctionInfo
}

func (msg *MsgS2CAuctionInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAuctionInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAuctionInfo(pcId uint32) MsgS2CAuctionInfo {
	msg := MsgS2CAuctionInfo{
		MsgHead: MsgHead{
			Protocol: S2CAuctionInfo,
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

type MsgS2CAgitPutupAuction struct {
	MsgHead
	ByErrorCode uint8
}

func (msg *MsgS2CAgitPutupAuction) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAgitPutupAuction) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAgitPutupAuction(pcId uint32) MsgS2CAgitPutupAuction {
	msg := MsgS2CAgitPutupAuction{
		MsgHead: MsgHead{
			Protocol: S2CAgitPutUpAuction,
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

type MsgS2CAgitBidOn struct {
	MsgHead
	ByErrorCode uint8
	DwMoney     uint32
}

func (msg *MsgS2CAgitBidOn) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAgitBidOn) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAgitBidOn(pcId uint32) MsgS2CAgitBidOn {
	msg := MsgS2CAgitBidOn{
		MsgHead: MsgHead{
			Protocol: S2CAgitBidOn,
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

type MsgS2CAgitPayExpense struct {
	MsgHead
	ByErrorCode uint8
	DwMoney     uint32
}

func (msg *MsgS2CAgitPayExpense) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAgitPayExpense) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAgitPayExpense(pcId uint32) MsgS2CAgitPayExpense {
	msg := MsgS2CAgitPayExpense{
		MsgHead: MsgHead{
			Protocol: S2CAgitPayExpense,
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

type MsgS2CAgitChangeName struct {
	MsgHead
	ByErrorCode uint8
	DwLore      uint32
}

func (msg *MsgS2CAgitChangeName) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAgitChangeName) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAgitChangeName(pcId uint32) MsgS2CAgitChangeName {
	msg := MsgS2CAgitChangeName{
		MsgHead: MsgHead{
			Protocol: S2CAgitChangeName,
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

type MsgC2SAgitInfo struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SAgitInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAgitInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAgitInfo(pcId uint32) MsgC2SAgitInfo {
	msg := MsgC2SAgitInfo{
		MsgHead: MsgHead{
			Protocol: C2SAgitInfo,
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

type MsgC2SAuctionInfo struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SAuctionInfo) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAuctionInfo) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAuctionInfo(pcId uint32) MsgC2SAuctionInfo {
	msg := MsgC2SAuctionInfo{
		MsgHead: MsgHead{
			Protocol: C2SAuctionInfo,
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

type MsgC2SAgitPutUpAuction struct {
	MsgHead
	DwNPCID        uint32
	DwAuctionTerm  uint32
	DwMaximumPrice uint32
	DwMinimumPrice uint32
}

func (msg *MsgC2SAgitPutUpAuction) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAgitPutUpAuction) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAgitPutUpAuction(pcId uint32) MsgC2SAgitPutUpAuction {
	msg := MsgC2SAgitPutUpAuction{
		MsgHead: MsgHead{
			Protocol: C2SAgitPutUpAuction,
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

type MsgC2SAgitBidOn struct {
	MsgHead
	DwNPCID    uint32
	WAgitID    uint16
	DwBidPrice uint32
}

func (msg *MsgC2SAgitBidOn) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAgitBidOn) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAgitBidOn(pcId uint32) MsgC2SAgitBidOn {
	msg := MsgC2SAgitBidOn{
		MsgHead: MsgHead{
			Protocol: C2SAgitBidOn,
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

type MsgC2SAgitPayExpense struct {
	MsgHead
}

func (msg *MsgC2SAgitPayExpense) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAgitPayExpense) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAgitPayExpense(pcId uint32) MsgC2SAgitPayExpense {
	msg := MsgC2SAgitPayExpense{
		MsgHead: MsgHead{
			Protocol: C2SAgitPayExpense,
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

type MsgC2SAgitChangeName struct {
	MsgHead
	SzAgitName [0x20]byte
}

func (msg *MsgC2SAgitChangeName) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAgitChangeName) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAgitChangeName(pcId uint32) MsgC2SAgitChangeName {
	msg := MsgC2SAgitChangeName{
		MsgHead: MsgHead{
			Protocol: C2SAgitChangeName,
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
