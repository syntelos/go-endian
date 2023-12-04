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
	"testing"
)

func TestOrder(t *testing.T){

	fmt.Printf("Host byte order: %s\n",HostOrder)
}
func TestRandom(t *testing.T){
	var e error

	var vector16 uint16
	vector16, e = Rand16()
	if nil != e {
		t.Error(e)
	} else {
		fmt.Printf("Random [16]: 0x%04x\n",vector16)
	}

	var vector32 uint32
	vector32, e = Rand32()
	if nil != e {
		t.Error(e)
	} else {
		fmt.Printf("Random [32]: 0x%08x\n",vector32)
	}

	var vector64 uint64
	vector64, e = Rand64()
	if nil != e {
		t.Error(e)
	} else {
		fmt.Printf("Random [64]: 0x%016x\n",vector64)
	}
}
