package wave

import (
	"fmt"
	"io"
	"reflect"

	"github.com/bbars/wave/internal/binary"
)

type SampleReader interface {
	ReadSample() (Sample, error)
}

type sampleReader struct {
	r        io.Reader
	sf       SampleFormat
	channels uint16
	value    any
}

func NewAutoSampleReader(r io.Reader) (res SampleReader, err error) {
	h := Header{}
	if _, err = h.ReadFrom(r); err != nil {
		return nil, err // TODO: wrap
	}

	return NewSampleReader(r, h)
}

func NewSampleReader(r io.Reader, h Header) (res SampleReader, err error) {
	sr := sampleReader{
		r:        r,
		channels: h.Channels,
	}
	var ok bool
	if sr.sf, ok = h.SampleFormat(); !ok {
		return nil, fmt.Errorf("unsupported sample format")
	}

	sr.value = reflect.New(reflect.TypeOf(sr.sf.valueSample)).Interface()

	return sr, nil
}

var _ SampleReader = sampleReader{}

func (s sampleReader) ReadSample() (res Sample, err error) {
	if s.channels == 0 {
		return
	}

	res.Values = make([]float64, s.channels)
	for i := uint16(0); i < s.channels; i++ {
		err = binary.Read(s.r, s.sf.ByteOrder, s.value)
		if err != nil {
			return
		}
		res.Values[i] = toFloat64(s.value)
	}

	return
}
