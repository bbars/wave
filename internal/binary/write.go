package binary

import (
	"encoding/binary"
	"fmt"
	"io"
)

func Write(w io.Writer, order binary.ByteOrder, data any) error {
	switch v := data.(type) {
	case Int24:
		return writeUint24(w, order, Uint24(v))
	case Uint24:
		return writeUint24(w, order, v)
	case *Int24:
		return writeUint24(w, order, Uint24(*v))
	case *Uint24:
		return writeUint24(w, order, *v)
	default:
		return binary.Write(w, order, data)
	}
}

func writeUint24(w io.Writer, order binary.ByteOrder, v Uint24) (err error) {
	if order == binary.LittleEndian {
		_, err = w.Write([]byte{
			byte(v),
			byte(v >> 8),
			byte(v >> 16),
		})
	} else if order == binary.BigEndian {
		_, err = w.Write([]byte{
			byte(v >> 16),
			byte(v >> 8),
			byte(v),
		})
	} else {
		err = fmt.Errorf("unknown byte order")
	}

	return err
}
