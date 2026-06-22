package protocol

type MsgS2CAnsOpenStorage struct {
	MsgHead
	ByErrCode uint8
	DwMoney   uint32
	ByNumItem uint8
	Storage   [0x50]ZoneServerItemInStorage
}

func (msg *MsgS2CAnsOpenStorage) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsOpenStorage) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsOpenStorage(pcId uint32) MsgS2CAnsOpenStorage {
	msg := MsgS2CAnsOpenStorage{
		MsgHead: MsgHead{
			Protocol: S2CAnsOpenStorage,
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

type MsgS2CAnsInven2Storage struct {
	MsgHead
	ByErrCode      uint8
	ItemID         ZoneServerItemId
	ByInvenIndex   uint8
	ByStorageIndex uint8
	DwMoney        uint32
}

func (msg *MsgS2CAnsInven2Storage) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsInven2Storage) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsInven2Storage(pcId uint32) MsgS2CAnsInven2Storage {
	msg := MsgS2CAnsInven2Storage{
		MsgHead: MsgHead{
			Protocol: S2CAnsInven2Storage,
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

type MsgS2CAnsStorage2Inven struct {
	MsgHead
	ByErrCode      uint8
	ItemID         ZoneServerItemId
	ByStorageIndex uint8
	ByInvenIndex   uint8
}

func (msg *MsgS2CAnsStorage2Inven) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsStorage2Inven) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsStorage2Inven(pcId uint32) MsgS2CAnsStorage2Inven {
	msg := MsgS2CAnsStorage2Inven{
		MsgHead: MsgHead{
			Protocol: S2CAnsStorage2Inven,
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

type MsgC2SAskOpenStorage struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SAskOpenStorage) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskOpenStorage) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskOpenStorage(pcId uint32) MsgC2SAskOpenStorage {
	msg := MsgC2SAskOpenStorage{
		MsgHead: MsgHead{
			Protocol: C2SAskOpenStorage,
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

type MsgC2SAskInven2Storage struct {
	MsgHead
	DwNPCID        uint32
	ItemID         ZoneServerItemId
	ByInvenIndex   uint8
	ByStorageIndex uint8
}

func (msg *MsgC2SAskInven2Storage) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskInven2Storage) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskInven2Storage(pcId uint32) MsgC2SAskInven2Storage {
	msg := MsgC2SAskInven2Storage{
		MsgHead: MsgHead{
			Protocol: C2SAskInven2Storage,
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

type MsgC2SAskStorage2Inven struct {
	MsgHead
	DwNPCID        uint32
	ItemID         ZoneServerItemId
	ByStorageIndex uint8
	ByInvenIndex   uint8
}

func (msg *MsgC2SAskStorage2Inven) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskStorage2Inven) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskStorage2Inven(pcId uint32) MsgC2SAskStorage2Inven {
	msg := MsgC2SAskStorage2Inven{
		MsgHead: MsgHead{
			Protocol: C2SAskStorage2Inven,
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

type MsgC2SAskCloseStorage struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SAskCloseStorage) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskCloseStorage) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskCloseStorage(pcId uint32) MsgC2SAskCloseStorage {
	msg := MsgC2SAskCloseStorage{
		MsgHead: MsgHead{
			Protocol: C2SAskCloseStorage,
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
