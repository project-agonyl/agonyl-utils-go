package protocol

import "encoding/binary"

func messageSize(v any) uint32 {
	return uint32(binary.Size(v))
}
