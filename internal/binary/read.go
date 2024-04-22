package binary

import (
	"encoding/binary"
	"fmt"
	"io"
)

func Read(r io.Reader, order binary.ByteOrder, data any) error {
	switch v := data.(type) {
	case *Int24:
		return readInt24(r, order, v)
	case *Uint24:
		return readUint24(r, order, v)
	default:
		return binary.Read(r, order, data)
	}
}

func readInt24(r io.Reader, order binary.ByteOrder, v *Int24) error {
	b := make([]byte, 3)
	if _, err := r.Read(b); err != nil {
		return err
	}

	if order == binary.LittleEndian {
		*v = Int24(int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16)
	} else if order == binary.BigEndian {
		*v = Int24(int32(b[2]) | int32(b[1])<<8 | int32(b[0])<<16)
	} else {
		return fmt.Errorf("unknown byte order")
	}

	return nil
}

func readUint24(r io.Reader, order binary.ByteOrder, v *Uint24) error {
	b := make([]byte, 3)
	if _, err := r.Read(b); err != nil {
		return err
	}

	if order == binary.LittleEndian {
		*v = Uint24(int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16)
	} else if order == binary.BigEndian {
		*v = Uint24(int32(b[2]) | int32(b[1])<<8 | int32(b[0])<<16)
	} else {
		return fmt.Errorf("unknown byte order")
	}

	return nil
}
