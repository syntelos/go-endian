/*
 * RAND
 * Copyright 2023 John Douglas Pritchard, @syntelos
 */
package endian

import (
	"crypto/rand"
	"fmt"
)
/*
 */
func Rand16() (v uint16, e error) {
	var b []byte = make([]byte,2)
	var n int
	n, e = rand.Reader.Read(b)
	if nil != e {
		e = fmt.Errorf("EncodeRand16 [0]: %w",e)
		return 0, e
	} else if 2 != n {
		e = fmt.Errorf("EncodeRand32 expected (2) found (%d).",n)
		return 0, e
	} else {
		v = HostOrder.DecodeUint16(b)
		return v, nil
	}
}
func Rand32() (v uint32, e error) {
	var b []byte = make([]byte,4)
	var n int
	n, e = rand.Read(b)
	if nil != e {
		e = fmt.Errorf("EncodeRand32: %w",e)
		return 0, e
	} else if 4 != n {
		e = fmt.Errorf("EncodeRand32 expected (4) found (%d).",n)
		return 0, e
	} else {
		v = HostOrder.DecodeUint32(b)
		return v, nil
	}
}
func Rand64() (v uint64, e error) {
	var b []byte = make([]byte,8)
	var n int
	n, e = rand.Read(b)
	if nil != e {
		e = fmt.Errorf("EncodeRand64 [0]: %w",e)
		return 0, e
	} else if 8 != n {
		e = fmt.Errorf("EncodeRand32 expected (8) found (%d).",n)
		return 0, e
	} else {
		v = HostOrder.DecodeUint64(b)
		return v, nil
	}
}
