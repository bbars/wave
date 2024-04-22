package binary

import (
	"encoding/binary"
)

type ByteOrder binary.ByteOrder

var (
	LittleEndian = binary.LittleEndian
	BigEndian    = binary.BigEndian
)
