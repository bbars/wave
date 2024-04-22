package wave

import (
	"fmt"
	"io"
	"reflect"

	"github.com/bbars/wave/internal/binary"
)

type SampleWriter interface {
	WriteSample(Sample) error
}

type sampleWriter struct {
	w        io.Writer
	sf       SampleFormat
	channels uint16
	value    any
}

func NewSampleWriter(w io.Writer, h Header) (res SampleWriter, err error) {
	sw := sampleWriter{
		w:        w,
		channels: h.Channels,
	}
	var ok bool
	if sw.sf, ok = h.SampleFormat(); !ok {
		return nil, fmt.Errorf("unsupported sample format")
	}

	sw.value = reflect.New(reflect.TypeOf(sw.sf.valueSample)).Interface()

	if err != nil {
		return nil, err // TODO: wrap
	}

	return sw, nil
}

var _ SampleWriter = sampleWriter{}

func (s sampleWriter) WriteSample(sample Sample) (err error) {
	i := uint16(0)
	for _, f := range sample.Values {
		i++
		if i > s.channels {
			break
		}

		fromFloat64(f, s.value)

		if err = binary.Write(s.w, s.sf.ByteOrder, s.value); err != nil {
			return
		}
	}

	if i < s.channels {
		fromFloat64(0, s.value)

		for i < s.channels {
			if err = binary.Write(s.w, s.sf.ByteOrder, s.value); err != nil {
				return
			}
		}
	}

	return nil
}
