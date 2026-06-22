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

func NewMsgC2SPickItem(pcId uint32) MsgC2SPickItem {
	msg := MsgC2SPickItem{
		MsgHead: MsgHead{
			Protocol: C2SPickupItem,
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

func NewMsgC2SDropItem(pcId uint32) MsgC2SDropItem {
	msg := MsgC2SDropItem{
		MsgHead: MsgHead{
			Protocol: C2SDropItem,
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

func NewMsgC2SInvItemMove(pcId uint32) MsgC2SInvItemMove {
	msg := MsgC2SInvItemMove{
		MsgHead: MsgHead{
			Protocol: C2SMoveItem,
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

func NewMsgC2SWearUp(pcId uint32) MsgC2SWearUp {
	msg := MsgC2SWearUp{
		MsgHead: MsgHead{
			Protocol: C2SWearItem,
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

func NewMsgC2SWearDown(pcId uint32) MsgC2SWearDown {
	msg := MsgC2SWearDown{
		MsgHead: MsgHead{
			Protocol: C2SStripItem,
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

func NewMsgC2SBuyItem(pcId uint32) MsgC2SBuyItem {
	msg := MsgC2SBuyItem{
		MsgHead: MsgHead{
			Protocol: C2SBuyItem,
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

func NewMsgC2SUseScroll(pcId uint32) MsgC2SUseScroll {
	msg := MsgC2SUseScroll{
		MsgHead: MsgHead{
			Protocol: C2SUseScroll,
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

func NewMsgS2CWearUpResp(pcId uint32) MsgS2CWearUpResp {
	msg := MsgS2CWearUpResp{
		MsgHead: MsgHead{
			Protocol: S2CWearItemResp,
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

func NewMsgS2CWearDownResp(pcId uint32) MsgS2CWearDownResp {
	msg := MsgS2CWearDownResp{
		MsgHead: MsgHead{
			Protocol: S2CStripItemResp,
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

func NewMsgS2CItemAppear(pcId uint32) MsgS2CItemAppear {
	msg := MsgS2CItemAppear{
		MsgHead: MsgHead{
			Protocol: S2CItemAppear,
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

func NewMsgS2CAddItem(pcId uint32) MsgS2CAddItem {
	msg := MsgS2CAddItem{
		MsgHead: MsgHead{
			Protocol: S2CAddItem,
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

func NewMsgS2CDropItemResp(pcId uint32) MsgS2CDropItemResp {
	msg := MsgS2CDropItemResp{
		MsgHead: MsgHead{
			Protocol: S2CDropItemResp,
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

func NewMsgS2CAnsInvItemMove(pcId uint32) MsgS2CAnsInvItemMove {
	msg := MsgS2CAnsInvItemMove{
		MsgHead: MsgHead{
			Protocol: S2CAnsInvItemMove,
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

func NewMsgS2CBuyItemResp(pcId uint32) MsgS2CBuyItemResp {
	msg := MsgS2CBuyItemResp{
		MsgHead: MsgHead{
			Protocol: S2CBuyItemResp,
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

func NewMsgS2CUpdateItem(pcId uint32) MsgS2CUpdateItem {
	msg := MsgS2CUpdateItem{
		MsgHead: MsgHead{
			Protocol: S2CUpdateItem,
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

type MsgS2CItemDisappear struct {
	MsgHead
	ItemID      ZoneServerItemId
	DwCellIndex uint32
}

func (msg *MsgS2CItemDisappear) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CItemDisappear) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CItemDisappear(pcId uint32) MsgS2CItemDisappear {
	msg := MsgS2CItemDisappear{
		MsgHead: MsgHead{
			Protocol: S2CItemDisappear,
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

type MsgS2CSeeWear struct {
	MsgHead
	DwOPCID    uint32
	ItemAppear ZoneServerItemAppear
}

func (msg *MsgS2CSeeWear) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSeeWear) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSeeWear(pcId uint32) MsgS2CSeeWear {
	msg := MsgS2CSeeWear{
		MsgHead: MsgHead{
			Protocol: S2CSeeWear,
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

type MsgS2CSeeStrip struct {
	MsgHead
	DwOPCID   uint32
	WItemCode uint16
	WearIndex uint8
}

func (msg *MsgS2CSeeStrip) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSeeStrip) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSeeStrip(pcId uint32) MsgS2CSeeStrip {
	msg := MsgS2CSeeStrip{
		MsgHead: MsgHead{
			Protocol: S2CSeeStrip,
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

type MsgS2CSellItem struct {
	MsgHead
	ErrCode uint8
	DwMoney uint32
	ItemID  ZoneServerItemId
}

func (msg *MsgS2CSellItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSellItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSellItem(pcId uint32) MsgS2CSellItem {
	msg := MsgS2CSellItem{
		MsgHead: MsgHead{
			Protocol: S2CSellItem,
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

type MsgS2CGiveItem struct {
	MsgHead
	ItemID   ZoneServerItemId
	ByResult uint8
}

func (msg *MsgS2CGiveItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CGiveItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CGiveItem(pcId uint32) MsgS2CGiveItem {
	msg := MsgS2CGiveItem{
		MsgHead: MsgHead{
			Protocol: S2CGiveItem,
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

type MsgS2CGivenItem struct {
	MsgHead
	SzPCName   [0xd]byte
	InvenIndex uint8
	Item       ZoneServerItem
}

func (msg *MsgS2CGivenItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CGivenItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CGivenItem(pcId uint32) MsgS2CGivenItem {
	msg := MsgS2CGivenItem{
		MsgHead: MsgHead{
			Protocol: S2CGivenItem,
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

type MsgS2CAskDeal struct {
	MsgHead
	DwAskPCID uint32
}

func (msg *MsgS2CAskDeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAskDeal) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAskDeal(pcId uint32) MsgS2CAskDeal {
	msg := MsgS2CAskDeal{
		MsgHead: MsgHead{
			Protocol: S2CAskDeal,
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

type MsgS2CAnsDeal struct {
	MsgHead
	DwOPCID uint32
	ErrCode uint8
}

func (msg *MsgS2CAnsDeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsDeal) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsDeal(pcId uint32) MsgS2CAnsDeal {
	msg := MsgS2CAnsDeal{
		MsgHead: MsgHead{
			Protocol: S2CAnsDeal,
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

type MsgS2CPutInItem struct {
	MsgHead
	DwMoney  uint32
	WHPotion uint16
	WMPotion uint16
	Item     ZoneServerItem
}

func (msg *MsgS2CPutInItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPutInItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPutInItem(pcId uint32) MsgS2CPutInItem {
	msg := MsgS2CPutInItem{
		MsgHead: MsgHead{
			Protocol: S2CPutInItem,
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

type MsgS2CPutOutItem struct {
	MsgHead
	DwMoney  uint32
	WHPotion uint16
	WMPotion uint16
	ItemID   ZoneServerItemId
}

func (msg *MsgS2CPutOutItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPutOutItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPutOutItem(pcId uint32) MsgS2CPutOutItem {
	msg := MsgS2CPutOutItem{
		MsgHead: MsgHead{
			Protocol: S2CPutOutItem,
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

type MsgS2CDecideDeal struct {
	MsgHead
}

func (msg *MsgS2CDecideDeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CDecideDeal) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CDecideDeal(pcId uint32) MsgS2CDecideDeal {
	msg := MsgS2CDecideDeal{
		MsgHead: MsgHead{
			Protocol: S2CDecideDeal,
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

type MsgS2CSuccessDeal struct {
	MsgHead
	DwMoney      uint32
	WHPotion     uint16
	WMPotion     uint16
	TakeItemList [0x14]ZoneServerItemInInven
}

func (msg *MsgS2CSuccessDeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSuccessDeal) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSuccessDeal(pcId uint32) MsgS2CSuccessDeal {
	msg := MsgS2CSuccessDeal{
		MsgHead: MsgHead{
			Protocol: S2CSuccessDeal,
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

type MsgS2CFailDeal struct {
	MsgHead
	ErrCode uint8
}

func (msg *MsgS2CFailDeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CFailDeal) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CFailDeal(pcId uint32) MsgS2CFailDeal {
	msg := MsgS2CFailDeal{
		MsgHead: MsgHead{
			Protocol: S2CFailDeal,
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

type MsgS2CSuccessDealPet struct {
	MsgHead
	TakePetList [0x2]ZoneServerPetInfo
}

func (msg *MsgS2CSuccessDealPet) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSuccessDealPet) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSuccessDealPet(pcId uint32) MsgS2CSuccessDealPet {
	msg := MsgS2CSuccessDealPet{
		MsgHead: MsgHead{
			Protocol: S2CSuccessDealPet,
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

type MsgS2CConfirmItem struct {
	MsgHead
	Item    ZoneServerItem
	ByIndex uint8
	DwMoney uint32
}

func (msg *MsgS2CConfirmItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CConfirmItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CConfirmItem(pcId uint32) MsgS2CConfirmItem {
	msg := MsgS2CConfirmItem{
		MsgHead: MsgHead{
			Protocol: S2CConfirmItem,
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

type MsgS2CRemodelItem struct {
	MsgHead
	ByStoneIndex uint8
	Item         ZoneServerItem
	ByItemIndex  uint8
}

func (msg *MsgS2CRemodelItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CRemodelItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CRemodelItem(pcId uint32) MsgS2CRemodelItem {
	msg := MsgS2CRemodelItem{
		MsgHead: MsgHead{
			Protocol: S2CRemodelItem,
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

type MsgS2CUseScroll struct {
	MsgHead
	ByInvenIndex uint8
	DwItemID     uint32
	Amount       uint8
}

func (msg *MsgS2CUseScroll) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CUseScroll) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CUseScroll(pcId uint32) MsgS2CUseScroll {
	msg := MsgS2CUseScroll{
		MsgHead: MsgHead{
			Protocol: S2CUseScroll,
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

type MsgS2CItemCombination struct {
	MsgHead
	ByResult     uint8
	ByInvenIndex uint8
	NewItem      ZoneServerItem
}

func (msg *MsgS2CItemCombination) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CItemCombination) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CItemCombination(pcId uint32) MsgS2CItemCombination {
	msg := MsgS2CItemCombination{
		MsgHead: MsgHead{
			Protocol: S2CItemCombination,
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

type MsgC2SUseItem struct {
	MsgHead
	DwItemID uint32
}

func (msg *MsgC2SUseItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SUseItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SUseItem(pcId uint32) MsgC2SUseItem {
	msg := MsgC2SUseItem{
		MsgHead: MsgHead{
			Protocol: C2SUseItem,
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

type MsgC2SSellItem struct {
	MsgHead
	DwNPCID uint32
	Item    ZoneServerItem
}

func (msg *MsgC2SSellItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SSellItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SSellItem(pcId uint32) MsgC2SSellItem {
	msg := MsgC2SSellItem{
		MsgHead: MsgHead{
			Protocol: C2SSellItem,
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

type MsgC2SGiveItem struct {
	MsgHead
	DwToPCID uint32
	ItemID   ZoneServerItemId
}

func (msg *MsgC2SGiveItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SGiveItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SGiveItem(pcId uint32) MsgC2SGiveItem {
	msg := MsgC2SGiveItem{
		MsgHead: MsgHead{
			Protocol: C2SGiveItem,
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

type MsgC2SAskDeal struct {
	MsgHead
	DwOPCID uint32
}

func (msg *MsgC2SAskDeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskDeal) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskDeal(pcId uint32) MsgC2SAskDeal {
	msg := MsgC2SAskDeal{
		MsgHead: MsgHead{
			Protocol: C2SAskDeal,
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

type MsgC2SAnsDeal struct {
	MsgHead
	DwAskPCID uint32
	Answer    uint8
}

func (msg *MsgC2SAnsDeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAnsDeal) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAnsDeal(pcId uint32) MsgC2SAnsDeal {
	msg := MsgC2SAnsDeal{
		MsgHead: MsgHead{
			Protocol: C2SAnsDeal,
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

type MsgC2SPutInItem struct {
	MsgHead
	DwMoney  uint32
	WHPotion uint16
	WMPotion uint16
	ItemID   ZoneServerItemId
}

func (msg *MsgC2SPutInItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPutInItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SPutInItem(pcId uint32) MsgC2SPutInItem {
	msg := MsgC2SPutInItem{
		MsgHead: MsgHead{
			Protocol: C2SPutInItem,
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

type MsgC2SPutOutItem struct {
	MsgHead
	DwMoney  uint32
	WHPotion uint16
	WMPotion uint16
	ItemID   ZoneServerItemId
}

func (msg *MsgC2SPutOutItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPutOutItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SPutOutItem(pcId uint32) MsgC2SPutOutItem {
	msg := MsgC2SPutOutItem{
		MsgHead: MsgHead{
			Protocol: C2SPutOutItem,
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

type MsgC2SDecideDeal struct {
	MsgHead
	Decide uint8
}

func (msg *MsgC2SDecideDeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SDecideDeal) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SDecideDeal(pcId uint32) MsgC2SDecideDeal {
	msg := MsgC2SDecideDeal{
		MsgHead: MsgHead{
			Protocol: C2SDecideDeal,
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

type MsgC2SConfirmDeal struct {
	MsgHead
	Confirm uint8
}

func (msg *MsgC2SConfirmDeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SConfirmDeal) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SConfirmDeal(pcId uint32) MsgC2SConfirmDeal {
	msg := MsgC2SConfirmDeal{
		MsgHead: MsgHead{
			Protocol: C2SConfirmDeal,
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

type MsgC2SConfirmItem struct {
	MsgHead
	DwNPCID uint32
	ItemID  ZoneServerItemId
	ByIndex uint8
}

func (msg *MsgC2SConfirmItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SConfirmItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SConfirmItem(pcId uint32) MsgC2SConfirmItem {
	msg := MsgC2SConfirmItem{
		MsgHead: MsgHead{
			Protocol: C2SConfirmItem,
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

type MsgC2SRemodelItem struct {
	MsgHead
	ItemID       ZoneServerItemId
	ByItemIndex  uint8
	StoneID      ZoneServerItemId
	ByStoneIndex uint8
}

func (msg *MsgC2SRemodelItem) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SRemodelItem) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SRemodelItem(pcId uint32) MsgC2SRemodelItem {
	msg := MsgC2SRemodelItem{
		MsgHead: MsgHead{
			Protocol: C2SRemodelItem,
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

type MsgC2SItemCombination struct {
	MsgHead
	ArrItemID [0xa]ZoneServerItemId
	DwNPCID   uint32
}

func (msg *MsgC2SItemCombination) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SItemCombination) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SItemCombination(pcId uint32) MsgC2SItemCombination {
	msg := MsgC2SItemCombination{
		MsgHead: MsgHead{
			Protocol: C2SItemCombination,
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
