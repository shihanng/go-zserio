package ztype

import (
	"io"
	"encoding/binary"

	"github.com/x448/float16"
)

// WriteFloat16 writes a float value in half precision (FP16) format.
func WriteFloat16(w io.Writer, v float32) error {
	f := float16.Fromfloat32(v).Bits()
	return binary.Write(w, binary.BigEndian, f)
}

// WriteFloat32 writes a float32 to the bitstream.
func WriteFloat32(w io.Writer, v float32) error {
	return binary.Write(w, binary.BigEndian, v)
}

// WriteFloat64 writes a float64 to the bitstream.
func WriteFloat64(w io.Writer, v float64) error {
	return binary.Write(w, binary.BigEndian, v)
}
