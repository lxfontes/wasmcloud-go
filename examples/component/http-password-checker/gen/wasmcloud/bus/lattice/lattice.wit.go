// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package lattice represents the imported interface "wasmcloud:bus/lattice@1.0.0".
package lattice

import (
	"go.bytecodealliance.org/cm"
)

// CallTargetInterface represents the imported resource "wasmcloud:bus/lattice@1.0.0#call-target-interface".
//
//	resource call-target-interface
type CallTargetInterface cm.Resource

// ResourceDrop represents the imported resource-drop for resource "call-target-interface".
//
// Drops a resource handle.
//
//go:nosplit
func (self CallTargetInterface) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_CallTargetInterfaceResourceDrop((uint32)(self0))
	return
}

// NewCallTargetInterface represents the imported constructor for resource "call-target-interface".
//
//	constructor(namespace: string, %package: string, %interface: string)
//
//go:nosplit
func NewCallTargetInterface(namespace string, package_ string, interface_ string) (result CallTargetInterface) {
	namespace0, namespace1 := cm.LowerString(namespace)
	package0, package1 := cm.LowerString(package_)
	interface0, interface1 := cm.LowerString(interface_)
	result0 := wasmimport_NewCallTargetInterface((*uint8)(namespace0), (uint32)(namespace1), (*uint8)(package0), (uint32)(package1), (*uint8)(interface0), (uint32)(interface1))
	result = cm.Reinterpret[CallTargetInterface]((uint32)(result0))
	return
}

// SetLinkName represents the imported function "set-link-name".
//
//	set-link-name: func(name: string, interfaces: list<call-target-interface>)
//
//go:nosplit
func SetLinkName(name string, interfaces cm.List[CallTargetInterface]) {
	name0, name1 := cm.LowerString(name)
	interfaces0, interfaces1 := cm.LowerList(interfaces)
	wasmimport_SetLinkName((*uint8)(name0), (uint32)(name1), (*CallTargetInterface)(interfaces0), (uint32)(interfaces1))
	return
}