// Package itemfile reads and writes A3 item binary files and converts raw item
// records into the parsed item view used by Agonyl server tooling.
package itemfile

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cyberinferno/go-utils/utils"
)

// IT0Level is a parsed item level property for IT0 equipment items.
type IT0Level struct {
	Level               byte
	AttributeRange      uint16
	Attribute           uint16
	Strength            uint16
	Intelligence        uint16
	Dexterity           uint16
	AdditionalAttribute uint16
	RedOption           uint16
	GreyOption          uint16
	BlueOption          uint16
}

// IT0Property contains parsed per-level equipment properties.
type IT0Property struct {
	Levels []IT0Level
}

// IT1Property contains parsed properties for IT1 items.
type IT1Property struct {
	RequiredLevel uint16
	Attribute     uint16
	RedOption     uint16
	GreyOption    uint16
	BlueOption    uint16
}

// IT2Property contains parsed properties for IT2 items.
type IT2Property struct {
	RequiredLevel uint16
	SkillLevel    uint16
	Class         uint16
}

// Item is the parsed item representation shared by IT0, IT1, IT2, and IT3.
type Item struct {
	ItemCode    uint32
	SlotIndex   byte
	ItemName    string
	Itemtype    byte
	NPCPrice    uint32
	IT0Property *IT0Property
	IT1Property *IT1Property
	IT2Property *IT2Property
}

// IT0RawLevelProperties is one raw level property block inside IT0/IT0Ex.
type IT0RawLevelProperties struct {
	AdditionalAttribute uint16
	Strength            uint16
	Dexterity           uint16
	Intelligence        uint16
	Attribute           uint16
	Range               uint16
	BlueOption          uint16
	RedOption           uint16
	GreyOption          uint16
}

// IT0Raw is one raw IT0 record.
type IT0Raw struct {
	ItemCodeBase uint16
	Row          uint16
	Slot         uint16
	Type         uint16
	Name         [32]byte
	NPCPrice     uint32
	Unknown2     [9]uint16
	Levels       [10]IT0RawLevelProperties
}

// IT0ExRaw is one raw IT0 extension record.
type IT0ExRaw struct {
	Row    uint16
	Levels [5]IT0RawLevelProperties
}

// IT1Raw is one raw IT1 record.
type IT1Raw struct {
	Type          uint16
	Row           uint16
	Name          [32]byte
	NPCPrice      uint32
	Unknown1      uint16
	RequiredLevel uint16
	Attribute     uint16
	BlueOption    uint16
	RedOption     uint16
	GreyOption    uint16
}

// IT2Raw is one raw IT2 record.
type IT2Raw struct {
	Type          uint16
	Row           uint16
	Name          [32]byte
	NPCPrice      uint32
	Class         uint16
	RequiredLevel uint16
	Unknown1      uint16
	SkillLevel    uint16
}

// IT3Raw is one raw IT3 record.
type IT3Raw struct {
	Type     uint16
	Row      uint16
	Name     [32]byte
	NPCPrice uint32
	Unknown1 uint16
	Unknown2 uint16
	Unknown3 uint16
	Unknown4 uint16
}

// IT0File is a raw IT0 file: contiguous IT0Raw records with no count prefix.
type IT0File []IT0Raw

// IT0ExFile is a raw IT0 extension file: contiguous IT0ExRaw records.
type IT0ExFile []IT0ExRaw

// IT1File is a raw IT1 file: contiguous IT1Raw records.
type IT1File []IT1Raw

// IT2File is a raw IT2 file: contiguous IT2Raw records.
type IT2File []IT2Raw

// IT3File is a raw IT3 file: contiguous IT3Raw records.
type IT3File []IT3Raw

// ReadIT0 reads a raw IT0 file from r.
func ReadIT0(r io.Reader) (IT0File, error) {
	return readRaw[IT0Raw, IT0File](r)
}

// WriteIT0 writes a raw IT0 file to w.
func WriteIT0(w io.Writer, data IT0File) error {
	return writeRaw(w, data)
}

// ReadIT0Ex reads a raw IT0 extension file from r.
func ReadIT0Ex(r io.Reader) (IT0ExFile, error) {
	return readRaw[IT0ExRaw, IT0ExFile](r)
}

// WriteIT0Ex writes a raw IT0 extension file to w.
func WriteIT0Ex(w io.Writer, data IT0ExFile) error {
	return writeRaw(w, data)
}

// ReadIT1 reads a raw IT1 file from r.
func ReadIT1(r io.Reader) (IT1File, error) {
	return readRaw[IT1Raw, IT1File](r)
}

// WriteIT1 writes a raw IT1 file to w.
func WriteIT1(w io.Writer, data IT1File) error {
	return writeRaw(w, data)
}

// ReadIT2 reads a raw IT2 file from r.
func ReadIT2(r io.Reader) (IT2File, error) {
	return readRaw[IT2Raw, IT2File](r)
}

// WriteIT2 writes a raw IT2 file to w.
func WriteIT2(w io.Writer, data IT2File) error {
	return writeRaw(w, data)
}

// ReadIT3 reads a raw IT3 file from r.
func ReadIT3(r io.Reader) (IT3File, error) {
	return readRaw[IT3Raw, IT3File](r)
}

// WriteIT3 writes a raw IT3 file to w.
func WriteIT3(w io.Writer, data IT3File) error {
	return writeRaw(w, data)
}

// ParseIT0Items converts raw IT0 and IT0Ex data into parsed items.
func ParseIT0Items(it0 IT0File, it0Ex IT0ExFile) ([]Item, error) {
	items := make([]Item, len(it0))
	for i := range it0 {
		raw := it0[i]
		row := int(raw.Row)
		if row >= len(items) {
			return nil, fmt.Errorf("itemfile: IT0 row %d out of range for %d items", raw.Row, len(items))
		}

		property := &IT0Property{Levels: make([]IT0Level, len(raw.Levels))}
		for j, level := range raw.Levels {
			property.Levels[j] = parseIT0Level(byte(j+1), level)
		}

		items[row] = Item{
			ItemCode:    uint32((raw.ItemCodeBase << 10) + raw.Row),
			SlotIndex:   byte(raw.Slot),
			ItemName:    utils.ReadStringFromBytes(raw.Name[:]),
			Itemtype:    byte(raw.Type),
			NPCPrice:    raw.NPCPrice,
			IT0Property: property,
		}
	}

	for i := range it0Ex {
		raw := it0Ex[i]
		row := int(raw.Row)
		if row >= len(items) {
			return nil, fmt.Errorf("itemfile: IT0Ex row %d out of range for %d items", raw.Row, len(items))
		}
		if items[row].IT0Property == nil {
			return nil, fmt.Errorf("itemfile: IT0Ex row %d references missing IT0 base item", raw.Row)
		}

		for j, level := range raw.Levels {
			items[row].IT0Property.Levels = append(items[row].IT0Property.Levels, parseIT0Level(byte(j+1), level))
		}
	}

	return items, nil
}

// ParseIT1Items converts raw IT1 data into parsed items.
func ParseIT1Items(it1 IT1File) ([]Item, error) {
	items := make([]Item, len(it1))
	for i := range it1 {
		raw := it1[i]
		row := int(raw.Row)
		if row >= len(items) {
			return nil, fmt.Errorf("itemfile: IT1 row %d out of range for %d items", raw.Row, len(items))
		}

		slotIndex := byte(9)
		if raw.Type == 4 {
			slotIndex = 8
		}

		items[row] = Item{
			ItemCode:  uint32((raw.Type << 10) + raw.Row),
			SlotIndex: slotIndex,
			ItemName:  utils.ReadStringFromBytes(raw.Name[:]),
			Itemtype:  byte(raw.Type),
			NPCPrice:  raw.NPCPrice,
			IT1Property: &IT1Property{
				RequiredLevel: raw.RequiredLevel,
				Attribute:     raw.Attribute,
				RedOption:     raw.RedOption,
				GreyOption:    raw.GreyOption,
				BlueOption:    raw.BlueOption,
			},
		}
	}

	return items, nil
}

// ParseIT2Items converts raw IT2 data into parsed items.
func ParseIT2Items(it2 IT2File) ([]Item, error) {
	items := make([]Item, len(it2))
	for i := range it2 {
		raw := it2[i]
		row := int(raw.Row)
		if row >= len(items) {
			return nil, fmt.Errorf("itemfile: IT2 row %d out of range for %d items", raw.Row, len(items))
		}

		items[row] = Item{
			ItemCode: uint32((raw.Type << 10) + raw.Row),
			ItemName: utils.ReadStringFromBytes(raw.Name[:]),
			Itemtype: byte(raw.Type),
			NPCPrice: raw.NPCPrice,
			IT2Property: &IT2Property{
				RequiredLevel: raw.RequiredLevel,
				SkillLevel:    raw.SkillLevel,
				Class:         raw.Class,
			},
		}
	}

	return items, nil
}

// ParseIT3Items converts raw IT3 data into parsed items.
func ParseIT3Items(it3 IT3File) ([]Item, error) {
	items := make([]Item, len(it3))
	for i := range it3 {
		raw := it3[i]
		row := int(raw.Row)
		if row >= len(items) {
			return nil, fmt.Errorf("itemfile: IT3 row %d out of range for %d items", raw.Row, len(items))
		}

		items[row] = Item{
			ItemCode: uint32((raw.Type << 10) + raw.Row),
			ItemName: utils.ReadStringFromBytes(raw.Name[:]),
			Itemtype: byte(raw.Type),
			NPCPrice: raw.NPCPrice,
		}
	}

	return items, nil
}

// ReadIT0Items reads raw IT0/IT0Ex data and returns parsed items.
func ReadIT0Items(it0 io.Reader, it0Ex io.Reader) ([]Item, error) {
	it0Data, err := ReadIT0(it0)
	if err != nil {
		return nil, err
	}

	it0ExData, err := ReadIT0Ex(it0Ex)
	if err != nil {
		return nil, err
	}

	return ParseIT0Items(it0Data, it0ExData)
}

// ReadIT1Items reads raw IT1 data and returns parsed items.
func ReadIT1Items(r io.Reader) ([]Item, error) {
	data, err := ReadIT1(r)
	if err != nil {
		return nil, err
	}

	return ParseIT1Items(data)
}

// ReadIT2Items reads raw IT2 data and returns parsed items.
func ReadIT2Items(r io.Reader) ([]Item, error) {
	data, err := ReadIT2(r)
	if err != nil {
		return nil, err
	}

	return ParseIT2Items(data)
}

// ReadIT3Items reads raw IT3 data and returns parsed items.
func ReadIT3Items(r io.Reader) ([]Item, error) {
	data, err := ReadIT3(r)
	if err != nil {
		return nil, err
	}

	return ParseIT3Items(data)
}

// GetName returns the IT0 item name as a string.
func (i *IT0Raw) GetName() string {
	return utils.ReadStringFromBytes(i.Name[:])
}

// GetName returns the IT1 item name as a string.
func (i *IT1Raw) GetName() string {
	return utils.ReadStringFromBytes(i.Name[:])
}

// GetName returns the IT2 item name as a string.
func (i *IT2Raw) GetName() string {
	return utils.ReadStringFromBytes(i.Name[:])
}

// GetName returns the IT3 item name as a string.
func (i *IT3Raw) GetName() string {
	return utils.ReadStringFromBytes(i.Name[:])
}

func readRaw[T any, S ~[]T](r io.Reader) (S, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	itemSize := binary.Size(*new(T))
	if itemSize <= 0 {
		return nil, fmt.Errorf("itemfile: invalid binary item size %d", itemSize)
	}
	if len(b)%itemSize != 0 {
		return nil, io.ErrUnexpectedEOF
	}

	data := make(S, len(b)/itemSize)
	if len(data) == 0 {
		return data, nil
	}

	if err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func writeRaw[T any, S ~[]T](w io.Writer, data S) error {
	return binary.Write(w, binary.LittleEndian, data)
}

func parseIT0Level(levelNumber byte, raw IT0RawLevelProperties) IT0Level {
	return IT0Level{
		Level:               levelNumber,
		AttributeRange:      raw.Range,
		Attribute:           raw.Attribute,
		Strength:            raw.Strength,
		Intelligence:        raw.Intelligence,
		Dexterity:           raw.Dexterity,
		AdditionalAttribute: raw.AdditionalAttribute,
		RedOption:           raw.RedOption,
		GreyOption:          raw.GreyOption,
		BlueOption:          raw.BlueOption,
	}
}
