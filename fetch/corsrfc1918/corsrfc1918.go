// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package corsrfc1918

import js "github.com/Programmerino/webapi/core/js"

// using following types:

// source idl files:
// cors-rfc1918.idl

// transform files:
// cors-rfc1918.go.md

// workaround for compiler error
func unused(value interface{}) {
	// TODO remove this method
}

type Union struct {
	Value js.Value
}

func (u *Union) JSValue() js.Value {
	return u.Value
}

func UnionFromJS(value js.Value) *Union {
	return &Union{Value: value}
}

// enum: AddressSpace
type AddressSpace int

const (
	LocalAddressSpace AddressSpace = iota
	PrivateAddressSpace
	PublicAddressSpace
)

var addressSpaceToWasmTable = []string{
	"local", "private", "public",
}

var addressSpaceFromWasmTable = map[string]AddressSpace{
	"local": LocalAddressSpace, "private": PrivateAddressSpace, "public": PublicAddressSpace,
}

// JSValue is converting this enum into a javascript object
func (this *AddressSpace) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this AddressSpace) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(addressSpaceToWasmTable) {
		return addressSpaceToWasmTable[idx]
	}
	panic("unknown input value")
}

// AddressSpaceFromJS is converting a javascript value into
// a AddressSpace enum value.
func AddressSpaceFromJS(value js.Value) AddressSpace {
	key := value.String()
	conv, ok := addressSpaceFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}
