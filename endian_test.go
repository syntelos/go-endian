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

	fmt.Println(HostOrder)
}
