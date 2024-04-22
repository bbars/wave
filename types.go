package wave

import (
	"github.com/bbars/wave/internal/binary"
)

type number interface {
	int8 | uint8 | int16 | uint16 | binary.Int24 | binary.Uint24 | int32 | uint32 | int64 | uint64 | float32 | float64
}
