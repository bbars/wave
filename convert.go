package wave

import (
	"fmt"

	"github.com/bbars/wave/internal/binary"
)

func toFloat64(a any) float64 {
	switch v := a.(type) {
	case int8:
		return float64(v) / float64(0xff/2)
	case uint8:
		return (float64(v) - float64(0xff/2)) / float64(0xff/2)
	case int16:
		return float64(v) / float64(0xffff/2)
	case uint16:
		return (float64(v) - float64(0xffff/2)) / float64(0xffff/2)
	case binary.Int24:
		return float64(v) / float64(0xffffff/2)
	case binary.Uint24:
		return (float64(v) - float64(0xffffff/2)) / float64(0xffffff/2)
	case int32:
		return float64(v) / float64(0xffffffff/2)
	case uint32:
		return (float64(v) - float64(0xffffffff/2)) / float64(0xffffffff/2)
	case int64:
		return float64(v) / float64(0xffffffffffffffff/2)
	case uint64:
		return (float64(v) - float64(0xffffffffffffffff/2)) / float64(0xffffffffffffffff/2)
	case float32:
		return float64(v)
	case float64:
		return v
	case *int8:
		return float64(*v) / float64(0xff/2)
	case *uint8:
		return (float64(*v) - float64(0xff/2)) / float64(0xff/2)
	case *int16:
		return float64(*v) / float64(0xffff/2)
	case *uint16:
		return (float64(*v) - float64(0xffff/2)) / float64(0xffff/2)
	case *binary.Int24:
		return float64(*v) / float64(0xffffff/2)
	case *binary.Uint24:
		return (float64(*v) - float64(0xffffff/2)) / float64(0xffffff/2)
	case *int32:
		return float64(*v) / float64(0xffffffff/2)
	case *uint32:
		return (float64(*v) - float64(0xffffffff/2)) / float64(0xffffffff/2)
	case *int64:
		return float64(*v) / float64(0xffffffffffffffff/2)
	case *uint64:
		return (float64(*v) - float64(0xffffffffffffffff/2)) / float64(0xffffffffffffffff/2)
	case *float32:
		return float64(*v)
	case *float64:
		return *v
	default:
		panic(fmt.Errorf("type %T is not supported by binary reader", v))
	}
}

func fromFloat64(v float64, p any) {
	switch out := p.(type) {
	case *int8:
		*out = int8(float64(0xff) * v)
	case *uint8:
		*out = uint8((0xff / 2) + float64(0xff/2)*v)
	case *int16:
		*out = int16(float64(0xffff/2) * v)
	case *uint16:
		*out = uint16((0xffff / 2) + float64(0xffff/2)*v)
	case *binary.Int24:
		*out = binary.Int24(float64(0xffffff/2)*v) & 0xffffff
	case *binary.Uint24:
		*out = binary.Uint24((0xffffff/2)+float64(0xffffff/2)*v) & 0xffffff
	case *int32:
		*out = int32(float64(0xffffffff/2) * v)
	case *uint32:
		*out = uint32((0xffffffff / 2) + float64(0xffffffff/2)*v)
	case *int64:
		*out = int64(float64(0xffffffffffffffff/2) * v)
	case *uint64:
		*out = uint64((0xffffffffffffffff / 2) + float64(0xffffffffffffffff/2)*v)
	case *float32:
		*out = float32(v)
	case *float64:
		*out = v
	default:
		panic(fmt.Errorf("type %T is not supported by binary reader", out))
	}
}
