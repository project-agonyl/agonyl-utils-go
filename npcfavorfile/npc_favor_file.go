// Package npcfavorfile reads and writes Zone Server FavIndex.dat files.
package npcfavorfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

// RecordSize is the encoded byte size of one favor row.
const RecordSize = 8

// Record preserves one favor-index row.
type Record struct {
	Raw [RecordSize]byte
}

// Data is a complete FavIndex.dat file.
type Data []Record

// Read reads all aligned favor rows.
func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("npc favor file: got %d bytes, record size %d", len(b), RecordSize)
	}

	data := make(Data, len(b)/RecordSize)
	for index := range data {
		copy(data[index].Raw[:], b[index*RecordSize:(index+1)*RecordSize])
	}

	return data, nil
}

// Write writes all favor rows.
func Write(w io.Writer, data Data) error {
	for _, record := range data {
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}

	return nil
}

func (r Record) Index() int32    { return int32(binary.LittleEndian.Uint32(r.Raw[0:4])) }
func (r Record) NPCType() uint32 { return binary.LittleEndian.Uint32(r.Raw[4:8]) }

func (r *Record) SetIndex(value int32)    { binary.LittleEndian.PutUint32(r.Raw[0:4], uint32(value)) }
func (r *Record) SetNPCType(value uint32) { binary.LittleEndian.PutUint32(r.Raw[4:8], value) }
