// Code generated by wit-bindgen-go. DO NOT EDIT.

package query

import (
	"go.bytecodealliance.org/cm"
	"test-harness/gen/wasmcloud/postgres/types"
	"unsafe"
)

// QueryErrorShape is used for storage in variant or result types.
type QueryErrorShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(types.QueryError{})]byte
}
