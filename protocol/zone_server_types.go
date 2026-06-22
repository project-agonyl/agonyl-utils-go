package protocol

// Unions use Raw because Go has no packed C union equivalent; bitfields use packed storage fields.
type ZoneServerAgitInfo struct {
	MWAgitID       uint16
	MSzAgitName    [0x20]byte
	MSzOwnClanName [0x20]byte
}

type ZoneServerAuctionInfo struct {
	MWAgitID        uint16
	Pad0            [0x2]uint8
	MDwTimeElapsed  uint32
	MDwBidPrice     uint32
	MDwMaximumPrice uint32
	MSzAgitName     [0x20]byte
}

type ZoneServerBattleWarpData struct {
	ByWarpType  uint8
	Pad0        [0x3]uint8
	DwOtherPCID uint32
}

type ZoneServerDerbyHistoryInfo struct {
	IDerbyIndex         uint16
	ByFirstPlace        uint8
	ByFirstPlaceIndex   uint8
	BySecondPlace       uint8
	BySecondPlaceIndex  uint8
	Pad0                [0x2]uint8
	FSingleEarningRatio float32
	FDoubleEarningRatio float32
}

type ZoneServerDerbyMonsterInfo struct {
	ByMonsterIndex uint8
	ByStamina      uint8
	BySpeed        uint8
	ByAcceleration uint8
	ByEndurance    uint8
	ByCornering    uint8
	WNumGames      uint16
	WNumWins       uint16
}

type ZoneServerGameOption struct {
	PackedBits0 uint8 // bitfields: bDeal:1, bPrivateSay:1, bParty:1, bGiven:1, bWearFlag:1, bReserved:3
}

type ZoneServerItem struct {
	ItemID   ZoneServerItemId
	ItemInfo ZoneServerItemInfo
	ItemKey  uint32
}

type ZoneServerItemAppear struct {
	PackedBits0 uint32 // bitfields: ItemCode:14, wearIndex:4, level:4, bIceAttr:1, bFireAttr:1, bLightAttr:1, reserved:7
}

type ZoneServerItemId struct {
	Id          uint32
	PackedBits0 uint32 // bitfields: code:14, bDefault:1, bBless:1, downOption:2, bItemFlag:1, bCheckFlag:1, reserved:12
}

type ZoneServerItemInfo struct {
	Raw [0x4]byte // union _ITEM_INFO_
	// Alternatives: uint8_t _raw[0x4]; uint32_t confirm; uint32_t fireAttr; uint32_t iceAttr; uint32_t iChosenNum1; uint32_t iChosenNum2; uint32_t iChosenNum3; uint32_t iChosenNum4; uint32_t iChosenNum5; uint32_t iDerbyBetMoney; uint32_t iDerbyChosenNum1; uint32_t iDerbyChosenNum2; uint32_t iDerbyIndex; uint32_t iLottoIndex; uint32_t info; uint32_t level; uint32_t lightAttr; uint16_t money; uint32_t option; uint16_t quantity; uint32_t subtype; uint32_t wearIdx
}

type ZoneServerItemInInven struct {
	Raw [0x14]byte // union _ITEM_IN_INVEN_
	// Alternatives: uint8_t _raw[0x14]; _ITEM_ item; _ITEM_ID_ ItemID; uint8_t _pad0[0x8]; _ITEM_INFO_ ItemInfo; uint8_t _pad1[0xc]; uint32_t ItemKey; uint8_t _pad2[0x10]; uint8_t byInvenIndex
}

type ZoneServerItemInStorage struct {
	ItemID         ZoneServerItemId
	ItemInfo       ZoneServerItemInfo
	ItemKey        uint32
	ByStorageIndex uint8
	Pad0           [0x3]uint8
}

type ZoneServerItemInWear struct {
	Raw [0x14]byte // union _ITEM_IN_WEAR_
	// Alternatives: uint8_t _raw[0x14]; _ITEM_ item; _ITEM_ID_ ItemID; uint8_t _pad0[0x8]; _ITEM_INFO_ ItemInfo; uint8_t _pad1[0xc]; uint32_t ItemKey; uint8_t _pad2[0x10]; uint8_t byWearIndex
}

type ZoneServerKnightInfo struct {
	SzPCName   [0xd]byte
	Pad0       [0x3]uint8
	DwPCID     uint32
	WLevel     uint16
	RankInClan uint8
	Status     uint8
	ByPCType   uint8
	Pad1       [0x3]uint8
}

type ZoneServerNpcSkillDamageinfo struct {
	DwTargetID  uint32
	DwCellIndex uint32
	BDie        uint8
}

type ZoneServerPartyMember struct {
	DwMemberID uint32
	SzName     [0xd]byte
	Pad0       [0x3]uint8
	DwCurCell  uint32
	WCurMap    uint16
	WHPRatio   uint16
}

type ZoneServerPc2Stat struct {
	WHitAttack    uint16
	WMagicAttack  uint16
	WDefense      uint16
	WFireAttack   uint16
	WFireDefence  uint16
	WIceAttack    uint16
	WIceDefense   uint16
	WLightAttack  uint16
	WLightDefense uint16
	WMaxHp        uint16
	WMaxMp        uint16
	WHitAddition  uint16
	WMagAddition  uint16
}

type ZoneServerPetAppear struct {
	Code  uint16
	Level uint16
}

type ZoneServerPetId struct {
	Id          uint32
	PackedBits0 uint32 // bitfields: code:16, subcode:16
}

type ZoneServerPetInfo struct {
	PetID       ZoneServerPetId
	SerialKey   uint32
	PackedBits0 uint32 // bitfields: Level:8, Exp:14, HP:6, FireAttk:1, FireDef:1, IceAttk:1, IceDef:1
	PackedBits1 uint32 // bitfields: LightAttk:1, LightDef:1, AttkRate:1, DefRate:1, MagicAttkRate:1, DownDropRate:1, UpExp:1, Food:14, petIndex:8, Confirm:1, bAlive:1, bCritiRateInc:1
}

type ZoneServerPetStat struct {
	PetID ZoneServerPetId
	Level uint8
	Pad0  [0x1]uint8
	Exp   uint16
	Food  uint16
	Pad1  [0x2]uint8
}

type ZoneServerSkillDamageinfo struct {
	DwTargetID  uint32
	WDamage     uint16
	DwCellIndex uint32
}

type ZoneServerSkillStatChange struct {
	AttkRate      uint8
	MagRate       uint8
	DefRate       uint8
	FireAttkRate  uint8
	FireDefRate   uint8
	IceAttkRate   uint8
	IceDefRate    uint8
	LightAttkRate uint8
	LightDefRate  uint8
	MaxHPRate     uint8
	MaxMPRate     uint8
}

type ZoneServerSocialinfo struct {
	PackedBits0 uint32 // bitfields: nation:8, rank:8, knight_index:16
}

type ZoneServerTyrEntry struct {
	SzPCName   [0xd]byte
	SzClanName [0x20]byte
}
