package protocol

type MsgS2CLottoPurchase struct {
	MsgHead
	ByResult      uint8
	ByInvenIndex  uint8
	ItemLotto     ZoneServerItem
	UPresentMoney uint32
}

func (msg *MsgS2CLottoPurchase) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CLottoPurchase) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CLottoPurchase(pcId uint32) MsgS2CLottoPurchase {
	msg := MsgS2CLottoPurchase{
		MsgHead: MsgHead{
			Protocol: S2CLottoPurchase,
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

type MsgS2CLottoQueryPrize struct {
	MsgHead
	UNowPrize uint32
}

func (msg *MsgS2CLottoQueryPrize) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CLottoQueryPrize) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CLottoQueryPrize(pcId uint32) MsgS2CLottoQueryPrize {
	msg := MsgS2CLottoQueryPrize{
		MsgHead: MsgHead{
			Protocol: S2CLottoQueryPrize,
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

type MsgS2CLottoQueryHistory struct {
	MsgHead
	NNowLotto        int32
	ArrWinNumHistory [0x19]uint8
}

func (msg *MsgS2CLottoQueryHistory) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CLottoQueryHistory) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CLottoQueryHistory(pcId uint32) MsgS2CLottoQueryHistory {
	msg := MsgS2CLottoQueryHistory{
		MsgHead: MsgHead{
			Protocol: S2CLottoQueryHistory,
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

type MsgS2CLottoSale struct {
	MsgHead
	ByResult      uint8
	UPrizeMoney   uint32
	UPresentMoney uint32
}

func (msg *MsgS2CLottoSale) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CLottoSale) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CLottoSale(pcId uint32) MsgS2CLottoSale {
	msg := MsgS2CLottoSale{
		MsgHead: MsgHead{
			Protocol: S2CLottoSale,
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

type MsgS2CLottoNotify struct {
	MsgHead
	ByNoticeType     uint8
	NNowLotto        int32
	ByArrWinNum      [0x5]uint8
	ByCarriedForward uint8
	U1stPrizeMoney   uint32
}

func (msg *MsgS2CLottoNotify) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CLottoNotify) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CLottoNotify(pcId uint32) MsgS2CLottoNotify {
	msg := MsgS2CLottoNotify{
		MsgHead: MsgHead{
			Protocol: S2CLottoNotify,
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

type MsgS2CDerbyNotice struct {
	MsgHead
	IDerbyIndex  uint16
	ByNoticeCode uint8
}

func (msg *MsgS2CDerbyNotice) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CDerbyNotice) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CDerbyNotice(pcId uint32) MsgS2CDerbyNotice {
	msg := MsgS2CDerbyNotice{
		MsgHead: MsgHead{
			Protocol: S2CDerbyNotice,
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

type MsgS2CDerbyIndexQuery struct {
	MsgHead
	IDerbyIndex uint16
}

func (msg *MsgS2CDerbyIndexQuery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CDerbyIndexQuery) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CDerbyIndexQuery(pcId uint32) MsgS2CDerbyIndexQuery {
	msg := MsgS2CDerbyIndexQuery{
		MsgHead: MsgHead{
			Protocol: S2CDerbyIndexQuery,
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

type MsgS2CDerbyMonsterQuery struct {
	MsgHead
	IDerbyIndex         uint16
	ArrDerbyMonsterInfo [0x5]ZoneServerDerbyMonsterInfo
}

func (msg *MsgS2CDerbyMonsterQuery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CDerbyMonsterQuery) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CDerbyMonsterQuery(pcId uint32) MsgS2CDerbyMonsterQuery {
	msg := MsgS2CDerbyMonsterQuery{
		MsgHead: MsgHead{
			Protocol: S2CDerbyMonsterQuery,
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

type MsgS2CDerbyRatioQuery struct {
	MsgHead
	FDerbyRatio       [0xf]float32
	IDerbyIndex       uint16
	BIsAbleToPurchase int32
}

func (msg *MsgS2CDerbyRatioQuery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CDerbyRatioQuery) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CDerbyRatioQuery(pcId uint32) MsgS2CDerbyRatioQuery {
	msg := MsgS2CDerbyRatioQuery{
		MsgHead: MsgHead{
			Protocol: S2CDerbyRatioQuery,
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

type MsgS2CDerbyPurchase struct {
	MsgHead
	ByResult       uint8
	ByInvenIndex   uint8
	ItemDerby      ZoneServerItem
	DwPresentMoney uint32
}

func (msg *MsgS2CDerbyPurchase) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CDerbyPurchase) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CDerbyPurchase(pcId uint32) MsgS2CDerbyPurchase {
	msg := MsgS2CDerbyPurchase{
		MsgHead: MsgHead{
			Protocol: S2CDerbyPurchase,
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

type MsgS2CDerbyRaceStart struct {
	MsgHead
	IDerbyIndex uint16
	FVelocity   [0x1e]float32
}

func (msg *MsgS2CDerbyRaceStart) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CDerbyRaceStart) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CDerbyRaceStart(pcId uint32) MsgS2CDerbyRaceStart {
	msg := MsgS2CDerbyRaceStart{
		MsgHead: MsgHead{
			Protocol: S2CDerbyRaceStart,
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

type MsgS2CDerbyResultQuery struct {
	MsgHead
	IDerbyIndex uint16
	DhiResult   ZoneServerDerbyHistoryInfo
}

func (msg *MsgS2CDerbyResultQuery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CDerbyResultQuery) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CDerbyResultQuery(pcId uint32) MsgS2CDerbyResultQuery {
	msg := MsgS2CDerbyResultQuery{
		MsgHead: MsgHead{
			Protocol: S2CDerbyResultQuery,
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

type MsgS2CDerbyHistoryQuery struct {
	MsgHead
	IDerbyIndex  uint16
	ILatestIndex uint16
	ByPage       uint8
	ByNumHistory uint8
	DhiHistory   [0xc]ZoneServerDerbyHistoryInfo
}

func (msg *MsgS2CDerbyHistoryQuery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CDerbyHistoryQuery) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CDerbyHistoryQuery(pcId uint32) MsgS2CDerbyHistoryQuery {
	msg := MsgS2CDerbyHistoryQuery{
		MsgHead: MsgHead{
			Protocol: S2CDerbyHistoryQuery,
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

type MsgS2CDerbyExchange struct {
	MsgHead
	IEarnedMoney   uint32
	FEarningRatio  float32
	DwPresentMoney uint32
	ByResult       uint8
}

func (msg *MsgS2CDerbyExchange) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CDerbyExchange) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CDerbyExchange(pcId uint32) MsgS2CDerbyExchange {
	msg := MsgS2CDerbyExchange{
		MsgHead: MsgHead{
			Protocol: S2CDerbyExchange,
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

type MsgC2SLottoPurchase struct {
	MsgHead
	ILottoSellingCounter int32
	ArrLottoNum          [0x5]uint8
	DwNPCID              uint32
}

func (msg *MsgC2SLottoPurchase) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SLottoPurchase) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SLottoPurchase(pcId uint32) MsgC2SLottoPurchase {
	msg := MsgC2SLottoPurchase{
		MsgHead: MsgHead{
			Protocol: C2SLottoPurchase,
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

type MsgC2SLottoQueryPrize struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SLottoQueryPrize) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SLottoQueryPrize) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SLottoQueryPrize(pcId uint32) MsgC2SLottoQueryPrize {
	msg := MsgC2SLottoQueryPrize{
		MsgHead: MsgHead{
			Protocol: C2SLottoQueryPrize,
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

type MsgC2SLottoQueryHistory struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SLottoQueryHistory) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SLottoQueryHistory) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SLottoQueryHistory(pcId uint32) MsgC2SLottoQueryHistory {
	msg := MsgC2SLottoQueryHistory{
		MsgHead: MsgHead{
			Protocol: C2SLottoQueryHistory,
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

type MsgC2SLottoSale struct {
	MsgHead
	ItemLotto ZoneServerItem
	DwNPCID   uint32
}

func (msg *MsgC2SLottoSale) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SLottoSale) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SLottoSale(pcId uint32) MsgC2SLottoSale {
	msg := MsgC2SLottoSale{
		MsgHead: MsgHead{
			Protocol: C2SLottoSale,
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

type MsgC2SDerbyIndexQuery struct {
	MsgHead
}

func (msg *MsgC2SDerbyIndexQuery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SDerbyIndexQuery) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SDerbyIndexQuery(pcId uint32) MsgC2SDerbyIndexQuery {
	msg := MsgC2SDerbyIndexQuery{
		MsgHead: MsgHead{
			Protocol: C2SDerbyIndexQuery,
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

type MsgC2SDerbyMonsterQuery struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SDerbyMonsterQuery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SDerbyMonsterQuery) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SDerbyMonsterQuery(pcId uint32) MsgC2SDerbyMonsterQuery {
	msg := MsgC2SDerbyMonsterQuery{
		MsgHead: MsgHead{
			Protocol: C2SDerbyMonsterQuery,
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

type MsgC2SDerbyRatioQuery struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SDerbyRatioQuery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SDerbyRatioQuery) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SDerbyRatioQuery(pcId uint32) MsgC2SDerbyRatioQuery {
	msg := MsgC2SDerbyRatioQuery{
		MsgHead: MsgHead{
			Protocol: C2SDerbyRatioQuery,
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

type MsgC2SDerbyPurchase struct {
	MsgHead
	IDerbyIndex uint16
	IChosenNum  [0x2]uint8
	IBetMoney   uint16
	DwNPCID     uint32
}

func (msg *MsgC2SDerbyPurchase) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SDerbyPurchase) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SDerbyPurchase(pcId uint32) MsgC2SDerbyPurchase {
	msg := MsgC2SDerbyPurchase{
		MsgHead: MsgHead{
			Protocol: C2SDerbyPurchase,
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

type MsgC2SDerbyResultQuery struct {
	MsgHead
}

func (msg *MsgC2SDerbyResultQuery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SDerbyResultQuery) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SDerbyResultQuery(pcId uint32) MsgC2SDerbyResultQuery {
	msg := MsgC2SDerbyResultQuery{
		MsgHead: MsgHead{
			Protocol: C2SDerbyResultQuery,
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

type MsgC2SDerbyHistoryQuery struct {
	MsgHead
	DwNPCID uint32
	ByPage  uint8
}

func (msg *MsgC2SDerbyHistoryQuery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SDerbyHistoryQuery) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SDerbyHistoryQuery(pcId uint32) MsgC2SDerbyHistoryQuery {
	msg := MsgC2SDerbyHistoryQuery{
		MsgHead: MsgHead{
			Protocol: C2SDerbyHistoryQuery,
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

type MsgC2SDerbyExchange struct {
	MsgHead
	ItemDerby ZoneServerItem
	DwNPCID   uint32
}

func (msg *MsgC2SDerbyExchange) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SDerbyExchange) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SDerbyExchange(pcId uint32) MsgC2SDerbyExchange {
	msg := MsgC2SDerbyExchange{
		MsgHead: MsgHead{
			Protocol: C2SDerbyExchange,
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
