// Code generated by wit-bindgen-go. DO NOT EDIT.

package logging

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:logging@0.1.0-draft".

//go:wasmimport wasi:logging/logging@0.1.0-draft log
//go:noescape
func wasmimport_Log(level0 uint32, context0 *uint8, context1 uint32, message0 *uint8, message1 uint32)

//go:wasmexport wasi:logging/logging@0.1.0-draft#log
//export wasi:logging/logging@0.1.0-draft#log
func wasmexport_Log(level0 uint32, context0 *uint8, context1 uint32, message0 *uint8, message1 uint32) {
	level := (Level)((uint32)(level0))
	context := cm.LiftString[string]((*uint8)(context0), (uint32)(context1))
	message := cm.LiftString[string]((*uint8)(message0), (uint32)(message1))
	Exports.Log(level, context, message)
	return
}
