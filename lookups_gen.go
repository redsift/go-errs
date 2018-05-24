package errs

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z InternalState) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt(o, int(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *InternalState) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zxvk int
		zxvk, bts, err = msgp.ReadIntBytes(bts)
		(*z) = InternalState(zxvk)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z InternalState) Msgsize() (s int) {
	s = msgp.IntSize
	return
}
