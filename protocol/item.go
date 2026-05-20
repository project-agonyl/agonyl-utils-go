package protocol

type ItemID struct {
	ID        uint32
	CodeParts uint32
}

type ItemInfo struct {
	Info uint32
}

type AccountItem struct {
	ItemID    ItemID
	ItemInfo  ItemInfo
	WearIndex uint32
}

type CharacterInventoryItem struct {
	ID    ItemID
	Info  ItemInfo
	Magic uint32
	Index uint32
}
