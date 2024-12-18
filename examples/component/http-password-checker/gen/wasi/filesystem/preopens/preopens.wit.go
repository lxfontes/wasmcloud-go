// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package preopens represents the imported interface "wasi:filesystem/preopens@0.2.0".
package preopens

import (
	"go.bytecodealliance.org/cm"
	"http-password-checker/gen/wasi/filesystem/types"
)

// Descriptor represents the imported type alias "wasi:filesystem/preopens@0.2.0#descriptor".
//
// See [types.Descriptor] for more information.
type Descriptor = types.Descriptor

// GetDirectories represents the imported function "get-directories".
//
//	get-directories: func() -> list<tuple<descriptor, string>>
//
//go:nosplit
func GetDirectories() (result cm.List[cm.Tuple[Descriptor, string]]) {
	wasmimport_GetDirectories(&result)
	return
}
