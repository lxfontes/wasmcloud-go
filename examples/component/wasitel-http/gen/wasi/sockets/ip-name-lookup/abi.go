// Code generated by wit-bindgen-go. DO NOT EDIT.

package ipnamelookup

import (
	"go.bytecodealliance.org/cm"
	"unsafe"
)

// OptionIPAddressShape is used for storage in variant or result types.
type OptionIPAddressShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(cm.Option[IPAddress]{})]byte
}
