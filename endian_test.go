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
/*
 * Integrity
 */
func TestOrder(t *testing.T){

	fmt.Printf("Host byte order: %s\n",HostOrder)
}
/*
 * Consistency
 */
func TestRandomPyramid(t *testing.T){
	var sequence uint8
	var count uint8
	for count = 2; count <= 8; count += 2 {

		for sequence = 1; sequence <= 16; sequence++ {

			var er error

			switch count {
			case 2:
				var value uint16
				value, er = Rand16()

				if nil != er {
					t.Errorf("Test pyramid random: %v",er)
				} else {
					var vector []byte = HostOrder.EncodeUint16(value)

					var check uint16 = HostOrder.DecodeUint16(vector)
					if check != value {
						t.Errorf("Test pyramid random check (0x%x) != value (0x%x)",check,value)
					}
				}

			case 4:
				var value uint32
				value, er = Rand32()

				if nil != er {
					t.Errorf("test pyramid random: %v",er)
				} else {
					var vector []byte = HostOrder.EncodeUint32(value)

					var check uint32 = HostOrder.DecodeUint32(vector)
					if check != value {
						t.Errorf("Test pyramid random check (0x%x) != value (0x%x)",check,value)
					}
				}

			case 8:
				var value uint64
				value, er = Rand64()

				if nil != er {
					t.Errorf("Test pyramid random: %v",er)
				} else {
					var vector []byte = HostOrder.EncodeUint64(value)

					var check uint64 = HostOrder.DecodeUint64(vector)
					if check != value {
						t.Errorf("Test pyramid random check (0x%x) != value (0x%x)",check,value)
					}
				}
			}
		}
	}
}
