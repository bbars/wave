package wave

import (
	"fmt"

	"github.com/bbars/wave/internal/binary"
)

func size[T number]() int {
	var t T
	var a any = t

	switch a.(type) {
	case int8:
		return 8 / 8
	case uint8:
		return 8 / 8
	case int16:
		return 16 / 8
	case uint16:
		return 16 / 8
	case binary.Int24:
		return 24 / 8
	case binary.Uint24:
		return 24 / 8
	case int32:
		return 32 / 8
	case uint32:
		return 32 / 8
	case int64:
		return 64 / 8
	case uint64:
		return 64 / 8
	case float32:
		return 32 / 8
	case float64:
		return 64 / 8
	default:
		panic(fmt.Errorf("unsupported type %T", t))
	}
}
