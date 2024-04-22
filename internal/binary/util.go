package binary

import (
	"io"
)

func ReadString(r io.Reader, numBytes int) (string, error) {
	b := make([]byte, numBytes)
	_, err := r.Read(b)
	return string(b), err
}
