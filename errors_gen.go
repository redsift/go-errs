package errs

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z ErrorSource) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "pointer"
	o = append(o, 0x82, 0xa7, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72)
	o = msgp.AppendString(o, z.Pointer)
	// string "parameter"
	o = append(o, 0xa9, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72)
	o = msgp.AppendString(o, z.Parameter)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ErrorSource) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "pointer":
			z.Pointer, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "parameter":
			z.Parameter, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ErrorSource) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Pointer) + 10 + msgp.StringPrefixSize + len(z.Parameter)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PropagatedError) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "id"
	o = append(o, 0x86, 0xa2, 0x69, 0x64)
	o = msgp.AppendString(o, z.Id)
	// string "code"
	o = append(o, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	o, err = z.Code.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "title"
	o = append(o, 0xa5, 0x74, 0x69, 0x74, 0x6c, 0x65)
	o = msgp.AppendString(o, z.Title)
	// string "detail"
	o = append(o, 0xa6, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c)
	o = msgp.AppendString(o, z.Detail)
	// string "link"
	o = append(o, 0xa4, 0x6c, 0x69, 0x6e, 0x6b)
	o = msgp.AppendString(o, z.Link)
	// string "source"
	o = append(o, 0xa6, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65)
	if z.Source == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 2
		// string "pointer"
		o = append(o, 0x82, 0xa7, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72)
		o = msgp.AppendString(o, z.Source.Pointer)
		// string "parameter"
		o = append(o, 0xa9, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72)
		o = msgp.AppendString(o, z.Source.Parameter)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PropagatedError) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "code":
			bts, err = z.Code.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "title":
			z.Title, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "detail":
			z.Detail, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "link":
			z.Link, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "source":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Source = nil
			} else {
				if z.Source == nil {
					z.Source = new(ErrorSource)
				}
				var zbai uint32
				zbai, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zbai > 0 {
					zbai--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "pointer":
						z.Source.Pointer, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "parameter":
						z.Source.Parameter, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PropagatedError) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.Id) + 5 + z.Code.Msgsize() + 6 + msgp.StringPrefixSize + len(z.Title) + 7 + msgp.StringPrefixSize + len(z.Detail) + 5 + msgp.StringPrefixSize + len(z.Link) + 7
	if z.Source == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 8 + msgp.StringPrefixSize + len(z.Source.Pointer) + 10 + msgp.StringPrefixSize + len(z.Source.Parameter)
	}
	return
}
