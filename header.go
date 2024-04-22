package wave

import (
	"fmt"
	"io"

	"github.com/bbars/wave/internal/binary"
)

type Header struct {
	// Assumed "RIFF"
	Magic string

	FileSize uint32

	// Always equals "WAVE"
	FileTypeHeader [4]byte

	ChunkFmt
}

func (h *Header) ReadFrom(r io.Reader) (n int64, err error) {
	if h.Magic, err = binary.ReadString(r, 4); err != nil {
		return
	} else if string(h.Magic) != "RIFF" {
		err = fmt.Errorf("can not find RIFF marker")
		return
	}
	n += 4

	if err = binary.Read(r, binary.LittleEndian, &h.FileSize); err != nil {
		return
	}
	n += 4

	if _, err = r.Read(h.FileTypeHeader[:]); err != nil {
		return
	}
	n += int64(len(h.FileTypeHeader[:]))

	var n2 int64
	if n2, err = h.ChunkFmt.ReadFrom(r); err != nil {
		return
	}
	n += n2

	return n, nil
}

func (h *Header) WriteTo(w io.Writer) (n int64, err error) {
	var n2 int

	n2, err = w.Write([]byte(h.Magic))
	n += int64(n2)
	if err != nil {
		return
	}

	if err = binary.Write(w, binary.LittleEndian, h.FileSize); err != nil {
		return
	}
	n += 4

	n2, err = w.Write(h.FileTypeHeader[:])
	n += int64(n2)
	if err != nil {
		return
	}

	var n3 int64
	n3, err = h.ChunkFmt.WriteTo(w)
	n += n3
	if err != nil {
		return
	}

	return n, nil
}

type FormatTag int16

const (
	FormatUnspecified FormatTag = 0
	FormatPcm                   = 1
)
