package wave

import (
	"io"

	"github.com/bbars/wave/internal/binary"
)

type ChunkData struct {
	// Always equals "data"
	ChunkId string

	ChunkSize uint32

	DataReader io.Reader
	DataWriter io.Writer
}

func (c *ChunkData) ReadFrom(r io.Reader) (n int64, err error) {
	if c.ChunkId, err = binary.ReadString(r, 4); err != nil {
		return
	}
	n += 4

	if err = binary.Read(r, binary.LittleEndian, &c.ChunkSize); err != nil {
		return
	}
	n += 4

	c.DataReader = r

	return n, nil
}

func (c *ChunkData) WriteTo(w io.Writer) (n int64, err error) {
	var n2 int

	n2, err = w.Write([]byte(c.ChunkId))
	n += int64(n2)
	if err != nil {
		return
	}
	n += int64(n2)

	if err = binary.Write(w, binary.LittleEndian, c.ChunkSize); err != nil {
		return
	}
	n += 4

	c.DataWriter = w

	return n, nil
}
