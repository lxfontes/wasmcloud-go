// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package network represents the imported interface "wasi:sockets/network@0.2.0".
package network

import (
	"go.bytecodealliance.org/cm"
)

// Network represents the imported resource "wasi:sockets/network@0.2.0#network".
//
//	resource network
type Network cm.Resource

// ResourceDrop represents the imported resource-drop for resource "network".
//
// Drops a resource handle.
//
//go:nosplit
func (self Network) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_NetworkResourceDrop((uint32)(self0))
	return
}

// ErrorCode represents the enum "wasi:sockets/network@0.2.0#error-code".
//
//	enum error-code {
//		unknown,
//		access-denied,
//		not-supported,
//		invalid-argument,
//		out-of-memory,
//		timeout,
//		concurrency-conflict,
//		not-in-progress,
//		would-block,
//		invalid-state,
//		new-socket-limit,
//		address-not-bindable,
//		address-in-use,
//		remote-unreachable,
//		connection-refused,
//		connection-reset,
//		connection-aborted,
//		datagram-too-large,
//		name-unresolvable,
//		temporary-resolver-failure,
//		permanent-resolver-failure
//	}
type ErrorCode uint8

const (
	ErrorCodeUnknown ErrorCode = iota
	ErrorCodeAccessDenied
	ErrorCodeNotSupported
	ErrorCodeInvalidArgument
	ErrorCodeOutOfMemory
	ErrorCodeTimeout
	ErrorCodeConcurrencyConflict
	ErrorCodeNotInProgress
	ErrorCodeWouldBlock
	ErrorCodeInvalidState
	ErrorCodeNewSocketLimit
	ErrorCodeAddressNotBindable
	ErrorCodeAddressInUse
	ErrorCodeRemoteUnreachable
	ErrorCodeConnectionRefused
	ErrorCodeConnectionReset
	ErrorCodeConnectionAborted
	ErrorCodeDatagramTooLarge
	ErrorCodeNameUnresolvable
	ErrorCodeTemporaryResolverFailure
	ErrorCodePermanentResolverFailure
)

var stringsErrorCode = [21]string{
	"unknown",
	"access-denied",
	"not-supported",
	"invalid-argument",
	"out-of-memory",
	"timeout",
	"concurrency-conflict",
	"not-in-progress",
	"would-block",
	"invalid-state",
	"new-socket-limit",
	"address-not-bindable",
	"address-in-use",
	"remote-unreachable",
	"connection-refused",
	"connection-reset",
	"connection-aborted",
	"datagram-too-large",
	"name-unresolvable",
	"temporary-resolver-failure",
	"permanent-resolver-failure",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e ErrorCode) String() string {
	return stringsErrorCode[e]
}

// IPAddressFamily represents the enum "wasi:sockets/network@0.2.0#ip-address-family".
//
//	enum ip-address-family {
//		ipv4,
//		ipv6
//	}
type IPAddressFamily uint8

const (
	IPAddressFamilyIPv4 IPAddressFamily = iota
	IPAddressFamilyIPv6
)

var stringsIPAddressFamily = [2]string{
	"ipv4",
	"ipv6",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e IPAddressFamily) String() string {
	return stringsIPAddressFamily[e]
}

// IPv4Address represents the tuple "wasi:sockets/network@0.2.0#ipv4-address".
//
//	type ipv4-address = tuple<u8, u8, u8, u8>
type IPv4Address [4]uint8

// IPv6Address represents the tuple "wasi:sockets/network@0.2.0#ipv6-address".
//
//	type ipv6-address = tuple<u16, u16, u16, u16, u16, u16, u16, u16>
type IPv6Address [8]uint16

// IPAddress represents the variant "wasi:sockets/network@0.2.0#ip-address".
//
//	variant ip-address {
//		ipv4(ipv4-address),
//		ipv6(ipv6-address),
//	}
type IPAddress cm.Variant[uint8, IPv6Address, IPv6Address]

// IPAddressIPv4 returns a [IPAddress] of case "ipv4".
func IPAddressIPv4(data IPv4Address) IPAddress {
	return cm.New[IPAddress](0, data)
}

// IPv4 returns a non-nil *[IPv4Address] if [IPAddress] represents the variant case "ipv4".
func (self *IPAddress) IPv4() *IPv4Address {
	return cm.Case[IPv4Address](self, 0)
}

// IPAddressIPv6 returns a [IPAddress] of case "ipv6".
func IPAddressIPv6(data IPv6Address) IPAddress {
	return cm.New[IPAddress](1, data)
}

// IPv6 returns a non-nil *[IPv6Address] if [IPAddress] represents the variant case "ipv6".
func (self *IPAddress) IPv6() *IPv6Address {
	return cm.Case[IPv6Address](self, 1)
}

var stringsIPAddress = [2]string{
	"ipv4",
	"ipv6",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v IPAddress) String() string {
	return stringsIPAddress[v.Tag()]
}

// IPv4SocketAddress represents the record "wasi:sockets/network@0.2.0#ipv4-socket-address".
//
//	record ipv4-socket-address {
//		port: u16,
//		address: ipv4-address,
//	}
type IPv4SocketAddress struct {
	_       cm.HostLayout
	Port    uint16
	Address IPv4Address
}

// IPv6SocketAddress represents the record "wasi:sockets/network@0.2.0#ipv6-socket-address".
//
//	record ipv6-socket-address {
//		port: u16,
//		flow-info: u32,
//		address: ipv6-address,
//		scope-id: u32,
//	}
type IPv6SocketAddress struct {
	_        cm.HostLayout
	Port     uint16
	FlowInfo uint32
	Address  IPv6Address
	ScopeID  uint32
}

// IPSocketAddress represents the variant "wasi:sockets/network@0.2.0#ip-socket-address".
//
//	variant ip-socket-address {
//		ipv4(ipv4-socket-address),
//		ipv6(ipv6-socket-address),
//	}
type IPSocketAddress cm.Variant[uint8, IPv6SocketAddressShape, IPv6SocketAddress]

// IPSocketAddressIPv4 returns a [IPSocketAddress] of case "ipv4".
func IPSocketAddressIPv4(data IPv4SocketAddress) IPSocketAddress {
	return cm.New[IPSocketAddress](0, data)
}

// IPv4 returns a non-nil *[IPv4SocketAddress] if [IPSocketAddress] represents the variant case "ipv4".
func (self *IPSocketAddress) IPv4() *IPv4SocketAddress {
	return cm.Case[IPv4SocketAddress](self, 0)
}

// IPSocketAddressIPv6 returns a [IPSocketAddress] of case "ipv6".
func IPSocketAddressIPv6(data IPv6SocketAddress) IPSocketAddress {
	return cm.New[IPSocketAddress](1, data)
}

// IPv6 returns a non-nil *[IPv6SocketAddress] if [IPSocketAddress] represents the variant case "ipv6".
func (self *IPSocketAddress) IPv6() *IPv6SocketAddress {
	return cm.Case[IPv6SocketAddress](self, 1)
}

var stringsIPSocketAddress = [2]string{
	"ipv4",
	"ipv6",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v IPSocketAddress) String() string {
	return stringsIPSocketAddress[v.Tag()]
}
