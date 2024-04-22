package wave

import (
	"io"

	"github.com/bbars/wave/internal/binary"
)

type ChunkFmt struct {
	// Always equals "fmt "
	ChunkId string

	ChunkSize uint32

	// Type of format (1 is PCM)
	FormatTag FormatTag

	// Number of channels
	Channels uint16

	// Number of Samples per second, or Hertz. Common values are 44100 (CD), 48000 (DAT).
	SamplesPerSec uint32

	// (SamplesPerSec * BitsPerSample * Channels) / 8
	AvgBytesPerSec uint32

	// (BitsPerSample * Channels) / 8.
	//  1 - 8 bit mono
	//  2 - 8 bit stereo/16 bit mono
	//  4 - 16 bit stereo
	BlockAlign uint16

	//
	BitsPerSample uint16
}

func (c *ChunkFmt) ReadFrom(r io.Reader) (n int64, err error) {
	if c.ChunkId, err = binary.ReadString(r, 4); err != nil {
		return
	}
	n += 4

	if err = binary.Read(r, binary.LittleEndian, &c.ChunkSize); err != nil {
		return
	}
	n += 4

	var format int16
	if err = binary.Read(r, binary.LittleEndian, &format); err != nil {
		return
	} else {
		c.FormatTag = FormatTag(format)
	}
	n += 2

	if err = binary.Read(r, binary.LittleEndian, &c.Channels); err != nil {
		return
	}
	n += 2

	if err = binary.Read(r, binary.LittleEndian, &c.SamplesPerSec); err != nil {
		return
	}
	n += 4

	if err = binary.Read(r, binary.LittleEndian, &c.AvgBytesPerSec); err != nil {
		return
	}
	n += 4

	if err = binary.Read(r, binary.LittleEndian, &c.BlockAlign); err != nil {
		return
	}
	n += 2

	if err = binary.Read(r, binary.LittleEndian, &c.BitsPerSample); err != nil {
		return
	}
	n += 2

	return n, nil
}

func (c *ChunkFmt) WriteTo(w io.Writer) (n int64, err error) {
	var n2 int

	n2, err = w.Write([]byte(c.ChunkId))
	n += int64(n2)
	if err != nil {
		return
	}
	n += int64(n2)

	err = binary.Write(w, binary.LittleEndian, c.ChunkSize)
	if err != nil {
		return
	}
	n += 4

	err = binary.Write(w, binary.LittleEndian, int16(c.FormatTag))
	if err != nil {
		return
	}
	n += 2

	err = binary.Write(w, binary.LittleEndian, int16(c.Channels))
	if err != nil {
		return
	}
	n += 2

	err = binary.Write(w, binary.LittleEndian, c.SamplesPerSec)
	if err != nil {
		return
	}
	n += 4

	err = binary.Write(w, binary.LittleEndian, c.AvgBytesPerSec)
	if err != nil {
		return
	}
	n += 4

	err = binary.Write(w, binary.LittleEndian, c.BlockAlign)
	if err != nil {
		return
	}
	n += 2

	err = binary.Write(w, binary.LittleEndian, c.BitsPerSample)
	if err != nil {
		return
	}
	n += 2

	return n, nil
}

func (c ChunkFmt) SampleFormat() (SampleFormat, bool) {
	for _, sf := range KnownSampleFormats {
		if sf.FormatTag == c.FormatTag && sf.BitsPerSample == c.BitsPerSample {
			return sf, true
		}
	}

	return SampleFormat{}, false
}
