/*
 * Host byte order for GOPL coders
 * Copyright 2023 John Douglas Pritchard, Syntelos
 *
 *
 * References
 *
 * https://pkg.go.dev/encoding/binary
 */
package endian

import (
	"fmt"
	"io"
)
/*
 */
type ByteOrder interface {


	EncodeUint16([2]byte, uint16)
	DecodeUint16([2]byte) (uint16)

	EncodeUint32([4]byte, uint32)
	DecodeUint32([4]byte) (uint32)

	EncodeUint64([8]byte, uint64)
	DecodeUint64([8]byte) (uint64)


	WriteUint16(io.Writer, uint16) (int, error)
	ReadUint16(io.ByteReader) (uint16, error)

	WriteUint32(io.Writer, uint32) (int, error)
	ReadUint32(io.ByteReader) (uint32, error)

	WriteUint64(io.Writer, uint64) (int, error)
	ReadUint64(io.ByteReader) (uint64, error)

	String() (string)
}
/*
 */
type ByteOrderSmall struct{}

type ByteOrderLarge struct{}

const ByteOrderSmallName string = "little-endian"

const ByteOrderLargeName string = "big-endian"
/*
 */
var LilEndian ByteOrderSmall

var BigEndian ByteOrderLarge

var HostOrder ByteOrder = HostByteOrder()
/*
 */
func HostByteOrder() (ByteOrder) {

	var hbo ByteOrder
	{
		const check uint16 = 0x00FF
		var vector [2]byte

		BigEndian.EncodeUint16(vector,check)

		var value uint16 = BigEndian.DecodeUint16(vector)

		if check == value {

			return BigEndian
		} else {

			return LilEndian
		}
	}
	return hbo
}
/*
 */
func (ByteOrderSmall) EncodeUint16(w [2]byte, v uint16){
	var a byte = byte(v & 0xFF)
	var b byte = byte((v >> 8) & 0xFF)

	w[0] = a
	w[1] = b
}
func (ByteOrderSmall) DecodeUint16(r [2]byte) (v uint16){
	var a, b byte

	a = r[0]
	b = r[1]

	var a16, b16 uint16 = uint16(a), uint16(b)
	v = ((b16 << 8) | a16)
	return v
}
func (ByteOrderSmall) EncodeUint32(w [4]byte, v uint32){
	var a byte = byte(v & 0xFF)
	var b byte = byte((v >>  8) & 0xFF)
	var c byte = byte((v >> 16) & 0xFF)
	var d byte = byte((v >> 24) & 0xFF)

	w[0] = a
	w[1] = b
	w[2] = c
	w[3] = d
}
func (ByteOrderSmall) DecodeUint32(r [4]byte) (v uint32){
	var a, b, c, d byte

	a = r[0]
	b = r[1]
	c = r[2]
	d = r[3]

	var d32, c32, b32, a32 uint32 = uint32(d), uint32(c), uint32(b), uint32(a)
	v = ((d32 << 24) | (c32 << 16) | (b32 << 8) | a32)
	return v
}
func (ByteOrderSmall) EncodeUint64(w [8]byte, v uint64){
	var a byte = byte(v & 0xFF)
	var b byte = byte((v >>  8) & 0xFF)
	var c byte = byte((v >> 16) & 0xFF)
	var d byte = byte((v >> 24) & 0xFF)

	var e byte = byte((v >> 32) & 0xFF)
	var f byte = byte((v >> 40) & 0xFF)
	var g byte = byte((v >> 48) & 0xFF)
	var h byte = byte((v >> 56) & 0xFF)

	w[0] = a
	w[1] = b
	w[2] = c
	w[3] = d
	w[4] = e
	w[5] = f
	w[6] = g
	w[7] = h
}
func (ByteOrderSmall) DecodeUint64(r [8]byte) (v uint64){
	var a, b, c, d, e, f, g, h byte

	a = r[0]
	b = r[1]
	c = r[2]
	d = r[3]
	e = r[4]
	f = r[5]
	g = r[6]
	h = r[7]

	var h64, g64, f64, e64, d64, c64, b64, a64 = uint64(h), uint64(g), uint64(f), uint64(e), uint64(d), uint64(c), uint64(b), uint64(a)
	v = ((h64 << 56) | (g64 << 48) | (f64 << 40) | (e64 << 32) | (d64 << 24) | (c64 << 16) | (b64 << 8) | a64)
	return v
}
func (ByteOrderSmall) WriteUint16(w io.Writer, v uint16) (c int, e error){
	var a byte = byte(v & 0xFF)
	var b byte = byte((v >> 8) & 0xFF)

	var out []byte
	{
		out[0] = a
		out[1] = b
	}
	return w.Write(out)
}
func (ByteOrderSmall) ReadUint16(r io.ByteReader) (v uint16, e error){
	v = 0

	var a, b byte

	a, e = r.ReadByte()
	if nil != e {
		fmt.Errorf("ReadUint16 [0]: %w",e)
		return v, e
	} else {
		b, e = r.ReadByte()
		if nil != e {
			fmt.Errorf("ReadUint16 [1]: %w",e)
			return v, e
		} else {
			var a16, b16 uint16 = uint16(a), uint16(b)
			v = ((b16 << 8) | a16)
			return v, nil
		}
	}
}
func (ByteOrderSmall) WriteUint32(w io.Writer, v uint32) (c int, e error){
	var a byte = byte(v & 0xFF)
	var b byte = byte((v >>  8) & 0xFF)
	var d byte = byte((v >> 16) & 0xFF)
	var f byte = byte((v >> 24) & 0xFF)

	var out []byte
	{
		out[0] = a
		out[1] = b
		out[2] = d
		out[3] = f
	}
	return w.Write(out)
}
func (ByteOrderSmall) ReadUint32(r io.ByteReader) (v uint32, e error){
	v = 0

	var a, b, c, d byte

	a, e = r.ReadByte()
	if nil != e {
		fmt.Errorf("ReadUint32 [0]: %w",e)
		return v, e
	} else {
		b, e = r.ReadByte()
		if nil != e {
			fmt.Errorf("ReadUint32 [1]: %w",e)
			return v, e
		} else {
			c, e = r.ReadByte()
			if nil != e {
				fmt.Errorf("ReadUint32 [1]: %w",e)
				return v, e
			} else {
				d, e = r.ReadByte()
				if nil != e {
					fmt.Errorf("ReadUint32 [1]: %w",e)
					return v, e
				} else {
					var d32, c32, b32, a32 uint32 = uint32(d), uint32(c), uint32(b), uint32(a)
					v = ((d32 << 24) | (c32 << 16) | (b32 << 8) | a32)
					return v, nil
				}
			}
		}
	}
}
func (ByteOrderSmall) WriteUint64(w io.Writer, v uint64) (c int, e error){
	var a byte = byte(v & 0xFF)
	var b byte = byte((v >>  8) & 0xFF)
	var d byte = byte((v >> 16) & 0xFF)
	var f byte = byte((v >> 24) & 0xFF)

	var g byte = byte((v >> 32) & 0xFF)
	var h byte = byte((v >> 40) & 0xFF)
	var i byte = byte((v >> 48) & 0xFF)
	var j byte = byte((v >> 56) & 0xFF)

	var out []byte
	{
		out[0] = a
		out[1] = b
		out[2] = d
		out[3] = f
		out[4] = g
		out[5] = h
		out[6] = i
		out[7] = j
	}
	return w.Write(out)
}
func (ByteOrderSmall) ReadUint64(r io.ByteReader) (v uint64, e error){
	v = 0

	var a, b, c, d, f, g, h, i byte

	a, e = r.ReadByte()
	if nil != e {
		fmt.Errorf("ReadUint64 [0]: %w",e)
		return v, e
	} else {
		b, e = r.ReadByte()
		if nil != e {
			fmt.Errorf("ReadUint64 [1]: %w",e)
			return v, e
		} else {
			c, e = r.ReadByte()
			if nil != e {
				fmt.Errorf("ReadUint64 [1]: %w",e)
				return v, e
			} else {
				d, e = r.ReadByte()
				if nil != e {
					fmt.Errorf("ReadUint64 [1]: %w",e)
					return v, e
				} else {
					f, e = r.ReadByte()
					if nil != e {
						fmt.Errorf("ReadUint64 [0]: %w",e)
						return v, e
					} else {
						g, e = r.ReadByte()
						if nil != e {
							fmt.Errorf("ReadUint64 [1]: %w",e)
							return v, e
						} else {
							h, e = r.ReadByte()
							if nil != e {
								fmt.Errorf("ReadUint64 [1]: %w",e)
								return v, e
							} else {
								i, e = r.ReadByte()
								if nil != e {
									fmt.Errorf("ReadUint64 [1]: %w",e)
									return v, e
								} else {
									var i64, h64, g64, f64, d64, c64, b64, a64 = uint64(i), uint64(h), uint64(g), uint64(f), uint64(d), uint64(c), uint64(b), uint64(a)
									v = ((i64 << 56) | (h64 << 48) | (g64 << 40) | (f64 << 32) | (d64 << 24) | (c64 << 16) | (b64 << 8) | a64)
									return v, nil
								}
							}
						}
					}
				}
			}
		}
	}
}
func (ByteOrderLarge) String() (string) {
	return ByteOrderLargeName
}
/*
 */
func (ByteOrderLarge) EncodeUint16(w [2]byte, v uint16){
	var a byte = byte((v >> 8) & 0xFF)
	var b byte = byte(v & 0xFF)

	w[0] = a
	w[1] = b
}
func (ByteOrderLarge) DecodeUint16(r [2]byte) (v uint16){
	var a, b byte

	a = r[0]
	b = r[1]

	var a16, b16 uint16 = uint16(a), uint16(b)
	v = ((a16 << 8) | b16)
	return v
}
func (ByteOrderLarge) EncodeUint32(w [4]byte, v uint32){
	var a byte = byte((v >> 24) & 0xFF)
	var b byte = byte((v >> 16) & 0xFF)
	var c byte = byte((v >>  8) & 0xFF)
	var d byte = byte(v & 0xFF)

	w[0] = a
	w[1] = b
	w[2] = c
	w[3] = d
}
func (ByteOrderLarge) DecodeUint32(r [4]byte) (v uint32){
	var a, b, c, d byte

	a = r[0]
	b = r[1]
	c = r[2]
	d = r[3]

	var a32, b32, c32, d32 uint32 = uint32(a), uint32(b), uint32(c), uint32(d)
	v = ((a32 << 24) | (b32 << 16) | (c32 << 8) | d32)
	return v
}
func (ByteOrderLarge) EncodeUint64(w [8]byte, v uint64){
	var a byte = byte((v >> 56) & 0xFF)
	var b byte = byte((v >> 48) & 0xFF)
	var c byte = byte((v >> 40) & 0xFF)
	var d byte = byte((v >> 32) & 0xFF)

	var e byte = byte((v >> 24) & 0xFF)
	var f byte = byte((v >> 16) & 0xFF)
	var g byte = byte((v >>  8) & 0xFF)
	var h byte = byte(v & 0xFF)

	w[0] = a
	w[1] = b
	w[2] = c
	w[3] = d
	w[4] = e
	w[5] = f
	w[6] = g
	w[7] = h
}
func (ByteOrderLarge) DecodeUint64(r [8]byte) (v uint64){
	var a, b, c, d, e, f, g, h byte

	a = r[0]
	b = r[1]
	c = r[2]
	d = r[3]
	e = r[4]
	f = r[5]
	g = r[6]
	h = r[7]

	var a64, b64, c64, d64, e64, f64, g64, h64 = uint64(a), uint64(b), uint64(c), uint64(d), uint64(e), uint64(f), uint64(g), uint64(h)
	v = ((a64 << 56) | (b64 << 48) | (c64 << 40) | (d64 << 32) | (e64 << 24) | (f64 << 16) | (g64 << 8) | h64)
	return v
}
func (ByteOrderLarge) WriteUint16(w io.Writer, v uint16) (c int, e error){
	var a byte = byte((v >> 8) & 0xFF)
	var b byte = byte(v & 0xFF)

	var out []byte
	{
		out[0] = a
		out[1] = b
	}
	return w.Write(out)
}
func (ByteOrderLarge) ReadUint16(r io.ByteReader) (v uint16, e error){
	v = 0

	var a, b byte

	a, e = r.ReadByte()
	if nil != e {
		fmt.Errorf("ReadUint16 [0]: %w",e)
		return v, e
	} else {
		b, e = r.ReadByte()
		if nil != e {
			fmt.Errorf("ReadUint16 [1]: %w",e)
			return v, e
		} else {
			var a16, b16 uint16 = uint16(a), uint16(b)
			v = ((a16 << 8) | b16)
			return v, nil
		}
	}
}
func (ByteOrderLarge) WriteUint32(w io.Writer, v uint32) (c int, e error){
	var a byte = byte((v >> 24) & 0xFF)
	var b byte = byte((v >> 16) & 0xFF)
	var d byte = byte((v >>  8) & 0xFF)
	var f byte = byte(v & 0xFF)

	var out []byte
	{
		out[0] = a
		out[1] = b
		out[2] = d
		out[3] = f
	}
	return w.Write(out)
}
func (ByteOrderLarge) ReadUint32(r io.ByteReader) (v uint32, e error){
	v = 0

	var a, b, c, d byte

	a, e = r.ReadByte()
	if nil != e {
		fmt.Errorf("ReadUint32 [0]: %w",e)
		return v, e
	} else {
		b, e = r.ReadByte()
		if nil != e {
			fmt.Errorf("ReadUint32 [1]: %w",e)
			return v, e
		} else {
			c, e = r.ReadByte()
			if nil != e {
				fmt.Errorf("ReadUint32 [1]: %w",e)
				return v, e
			} else {
				d, e = r.ReadByte()
				if nil != e {
					fmt.Errorf("ReadUint32 [1]: %w",e)
					return v, e
				} else {
					var a32, b32, c32, d32 uint32 = uint32(a), uint32(b), uint32(c), uint32(d)
					v = ((a32 << 24) | (b32 << 16) | (c32 << 8) | d32)
					return v, nil
				}
			}
		}
	}
}
func (ByteOrderLarge) WriteUint64(w io.Writer, v uint64) (c int, e error){
	var a byte = byte((v >> 56) & 0xFF)
	var b byte = byte((v >> 48) & 0xFF)
	var d byte = byte((v >> 40) & 0xFF)
	var f byte = byte((v >> 32) & 0xFF)

	var g byte = byte((v >> 24) & 0xFF)
	var h byte = byte((v >> 16) & 0xFF)
	var i byte = byte((v >>  8) & 0xFF)	
	var j byte = byte(v & 0xFF)

	var out []byte
	{
		out[0] = a
		out[1] = b
		out[2] = d
		out[3] = f
		out[4] = g
		out[5] = h
		out[6] = i
		out[7] = j
	}
	return w.Write(out)
}
func (ByteOrderLarge) ReadUint64(r io.ByteReader) (v uint64, e error){
	v = 0

	var a, b, c, d, f, g, h, i byte

	a, e = r.ReadByte()
	if nil != e {
		fmt.Errorf("ReadUint64 [0]: %w",e)
		return v, e
	} else {
		b, e = r.ReadByte()
		if nil != e {
			fmt.Errorf("ReadUint64 [1]: %w",e)
			return v, e
		} else {
			c, e = r.ReadByte()
			if nil != e {
				fmt.Errorf("ReadUint64 [1]: %w",e)
				return v, e
			} else {
				d, e = r.ReadByte()
				if nil != e {
					fmt.Errorf("ReadUint64 [1]: %w",e)
					return v, e
				} else {
					f, e = r.ReadByte()
					if nil != e {
						fmt.Errorf("ReadUint64 [0]: %w",e)
						return v, e
					} else {
						g, e = r.ReadByte()
						if nil != e {
							fmt.Errorf("ReadUint64 [1]: %w",e)
							return v, e
						} else {
							h, e = r.ReadByte()
							if nil != e {
								fmt.Errorf("ReadUint64 [1]: %w",e)
								return v, e
							} else {
								i, e = r.ReadByte()
								if nil != e {
									fmt.Errorf("ReadUint64 [1]: %w",e)
									return v, e
								} else {
									var a64, b64, c64, d64, f64, g64, h64, i64 uint64 = uint64(a), uint64(b), uint64(c), uint64(d), uint64(f), uint64(g), uint64(h), uint64(i)
									v = ((a64 << 56) | (b64 << 48) | (c64 << 40) | (d64 << 32) | (f64 << 24) | (g64 << 16) | (h64 << 8) | i64)
									return v, nil
								}
							}
						}
					}
				}
			}
		}
	}
}
func (ByteOrderSmall) String() (string) {
	return ByteOrderSmallName
}
