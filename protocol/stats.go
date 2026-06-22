package protocol

type MsgC2SPointInc struct {
	MsgHead
	Strength uint8
	Magic    uint8
	Dex      uint8
	Vit      uint8
	Mana     uint8
}

func (msg *MsgC2SPointInc) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPointInc) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SPointInc(pcId uint32) MsgC2SPointInc {
	msg := MsgC2SPointInc{
		MsgHead: MsgHead{
			Protocol: C2SAllotPoint,
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

type MsgC2SAskHeal struct {
	MsgHead
	ID uint32
}

func (msg *MsgC2SAskHeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskHeal) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskHeal(pcId uint32) MsgC2SAskHeal {
	msg := MsgC2SAskHeal{
		MsgHead: MsgHead{
			Protocol: C2SAskHeal,
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

type MsgC2SPointDec struct {
	MsgHead
	Strength uint8
	Magic    uint8
	Dex      uint8
	Vit      uint8
	Mana     uint8
}

func (msg *MsgC2SPointDec) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SPointDec) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SPointDec(pcId uint32) MsgC2SPointDec {
	msg := MsgC2SPointDec{
		MsgHead: MsgHead{
			Protocol: C2SRetrievePoint,
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

type CharacterPointStat struct {
	Point            uint16
	Strength         uint16
	Magic            uint16
	Dex              uint16
	Vit              uint16
	Mana             uint16
	MaxHPStore       uint32
	MaxMPStore       uint32
	HP               uint16
	MP               uint16
	MinAttack        uint16
	MinMagicAttack   uint16
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
	Unknown          uint16
	MaxMagicAttack   uint16
}

type MsgS2CPointIncResp struct {
	MsgHead
	CharacterPointStat
}

func (msg *MsgS2CPointIncResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPointIncResp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPointIncResp(pcId uint32) MsgS2CPointIncResp {
	msg := MsgS2CPointIncResp{
		MsgHead: MsgHead{
			Protocol: S2CPointIncResp,
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

type MsgS2CUpdateExp struct {
	MsgHead
	Exp     uint32
	Unknown uint16
}

func (msg *MsgS2CUpdateExp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CUpdateExp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CUpdateExp(pcId uint32) MsgS2CUpdateExp {
	msg := MsgS2CUpdateExp{
		MsgHead: MsgHead{
			Protocol: S2CUpdateExp,
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

type MsgS2CAskHeal struct {
	MsgHead
	HP uint16
	MP uint16
}

func (msg *MsgS2CAskHeal) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAskHeal) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAskHeal(pcId uint32) MsgS2CAskHeal {
	msg := MsgS2CAskHeal{
		MsgHead: MsgHead{
			Protocol: S2CAskHeal,
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

type MsgS2CPointDecResp struct {
	MsgHead
	Unknown0 byte
	CharacterPointStat
	LorePoint uint32
}

func (msg *MsgS2CPointDecResp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPointDecResp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPointDecResp(pcId uint32) MsgS2CPointDecResp {
	msg := MsgS2CPointDecResp{
		MsgHead: MsgHead{
			Protocol: S2CPointDecResp,
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

type MsgS2CPayMoney struct {
	MsgHead
	PayMoney uint32
	Unknown  byte
}

func (msg *MsgS2CPayMoney) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPayMoney) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPayMoney(pcId uint32) MsgS2CPayMoney {
	msg := MsgS2CPayMoney{
		MsgHead: MsgHead{
			Protocol: S2CPayMoney,
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

type MsgC2SRecovery struct {
	MsgHead
	Fix  byte
	Type byte
}

func (msg *MsgC2SRecovery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SRecovery) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SRecovery(pcId uint32) MsgC2SRecovery {
	msg := MsgC2SRecovery{
		MsgHead: MsgHead{
			Protocol: C2SRestoreExp,
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

type MsgS2CRecovery struct {
	MsgHead
	Type  byte
	Fix   byte
	Bar   uint32
	Store uint32
}

func (msg *MsgS2CRecovery) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CRecovery) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CRecovery(pcId uint32) MsgS2CRecovery {
	msg := MsgS2CRecovery{
		MsgHead: MsgHead{
			Protocol: S2CRecovery,
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

type MsgS2CPetLevelUp struct {
	MsgHead
	DwOPCID uint32
	ByLevel uint8
}

func (msg *MsgS2CPetLevelUp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CPetLevelUp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CPetLevelUp(pcId uint32) MsgS2CPetLevelUp {
	msg := MsgS2CPetLevelUp{
		MsgHead: MsgHead{
			Protocol: S2CPetLevelUp,
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

type MsgS2CSeeLevelUp struct {
	MsgHead
	WLevelUpPCID uint32
}

func (msg *MsgS2CSeeLevelUp) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CSeeLevelUp) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CSeeLevelUp(pcId uint32) MsgS2CSeeLevelUp {
	msg := MsgS2CSeeLevelUp{
		MsgHead: MsgHead{
			Protocol: S2CSeeLevelUp,
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

type MsgS2CUsePotion struct {
	MsgHead
	ByResult   uint8
	DwRecovery uint16
	DwStored   uint16
}

func (msg *MsgS2CUsePotion) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CUsePotion) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CUsePotion(pcId uint32) MsgS2CUsePotion {
	msg := MsgS2CUsePotion{
		MsgHead: MsgHead{
			Protocol: S2CUsePotion,
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

type MsgS2CAnsDepositeMoney struct {
	MsgHead
	ByErrCode      uint8
	DwStorageMoney uint32
	DwPCMoney      uint32
}

func (msg *MsgS2CAnsDepositeMoney) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsDepositeMoney) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsDepositeMoney(pcId uint32) MsgS2CAnsDepositeMoney {
	msg := MsgS2CAnsDepositeMoney{
		MsgHead: MsgHead{
			Protocol: S2CAnsDepositeMoney,
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

type MsgS2CAnsWithdrawMoney struct {
	MsgHead
	ByErrCode      uint8
	DwStorageMoney uint32
	DwPCMoney      uint32
}

func (msg *MsgS2CAnsWithdrawMoney) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAnsWithdrawMoney) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAnsWithdrawMoney(pcId uint32) MsgS2CAnsWithdrawMoney {
	msg := MsgS2CAnsWithdrawMoney{
		MsgHead: MsgHead{
			Protocol: S2CAnsWithdrawMoney,
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

type MsgS2CCaoMitigation struct {
	MsgHead
	ByResultCode uint8
	WRTime       uint16
	DwLore       uint32
}

func (msg *MsgS2CCaoMitigation) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CCaoMitigation) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CCaoMitigation(pcId uint32) MsgS2CCaoMitigation {
	msg := MsgS2CCaoMitigation{
		MsgHead: MsgHead{
			Protocol: S2CCaoMitigation,
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

type MsgS2CUpdateLore struct {
	MsgHead
	IAddLore int32
	Code     uint8
}

func (msg *MsgS2CUpdateLore) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CUpdateLore) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CUpdateLore(pcId uint32) MsgS2CUpdateLore {
	msg := MsgS2CUpdateLore{
		MsgHead: MsgHead{
			Protocol: S2CUpdateLore,
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

type MsgS2CAgitRepayMoney struct {
	MsgHead
	ByErrorCode uint8
	DwMoney     uint32
}

func (msg *MsgS2CAgitRepayMoney) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAgitRepayMoney) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAgitRepayMoney(pcId uint32) MsgS2CAgitRepayMoney {
	msg := MsgS2CAgitRepayMoney{
		MsgHead: MsgHead{
			Protocol: S2CAgitRepayMoney,
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

type MsgS2CAgitObtainSaleMoney struct {
	MsgHead
	ByErrorCode uint8
	DwMoney     uint32
}

func (msg *MsgS2CAgitObtainSaleMoney) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgS2CAgitObtainSaleMoney) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgS2CAgitObtainSaleMoney(pcId uint32) MsgS2CAgitObtainSaleMoney {
	msg := MsgS2CAgitObtainSaleMoney{
		MsgHead: MsgHead{
			Protocol: S2CAgitObtainSaleMoney,
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

type MsgC2SUsePotion struct {
	MsgHead
	ByPotionType uint8
}

func (msg *MsgC2SUsePotion) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SUsePotion) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SUsePotion(pcId uint32) MsgC2SUsePotion {
	msg := MsgC2SUsePotion{
		MsgHead: MsgHead{
			Protocol: C2SUsePotion,
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

type MsgC2SAskDepositeMoney struct {
	MsgHead
	DwNPCID         uint32
	DwDepositeMoney uint32
}

func (msg *MsgC2SAskDepositeMoney) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskDepositeMoney) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskDepositeMoney(pcId uint32) MsgC2SAskDepositeMoney {
	msg := MsgC2SAskDepositeMoney{
		MsgHead: MsgHead{
			Protocol: C2SAskDepositeMoney,
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

type MsgC2SAskWithdrawMoney struct {
	MsgHead
	DwNPCID         uint32
	DwWithdrawMoney uint32
}

func (msg *MsgC2SAskWithdrawMoney) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAskWithdrawMoney) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAskWithdrawMoney(pcId uint32) MsgC2SAskWithdrawMoney {
	msg := MsgC2SAskWithdrawMoney{
		MsgHead: MsgHead{
			Protocol: C2SAskWithdrawMoney,
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

type MsgC2SCaoMitigation struct {
	MsgHead
	WMinutes uint16
}

func (msg *MsgC2SCaoMitigation) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SCaoMitigation) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SCaoMitigation(pcId uint32) MsgC2SCaoMitigation {
	msg := MsgC2SCaoMitigation{
		MsgHead: MsgHead{
			Protocol: C2SCaoMitigation,
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

type MsgC2SAgitRepayMoney struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SAgitRepayMoney) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAgitRepayMoney) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAgitRepayMoney(pcId uint32) MsgC2SAgitRepayMoney {
	msg := MsgC2SAgitRepayMoney{
		MsgHead: MsgHead{
			Protocol: C2SAgitRepayMoney,
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

type MsgC2SAgitObtainSaleMoney struct {
	MsgHead
	DwNPCID uint32
}

func (msg *MsgC2SAgitObtainSaleMoney) GetSize() uint32 {
	return messageSize(msg)
}

func (msg *MsgC2SAgitObtainSaleMoney) SetSize() {
	msg.Size = msg.GetSize()
}

func NewMsgC2SAgitObtainSaleMoney(pcId uint32) MsgC2SAgitObtainSaleMoney {
	msg := MsgC2SAgitObtainSaleMoney{
		MsgHead: MsgHead{
			Protocol: C2SAgitObtainSaleMoney,
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
