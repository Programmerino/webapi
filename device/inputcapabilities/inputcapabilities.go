// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package inputcapabilities

import js "github.com/gowebapi/webapi/core/js"

// using following types:

// source idl files:
// InputDeviceCapabilities.idl

// transform files:
// InputDeviceCapabilities.go.md

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

// dictionary: InputDeviceCapabilitiesInit
type InputDeviceCapabilitiesInit struct {
	FiresTouchEvents       bool
	PointerMovementScrolls bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *InputDeviceCapabilitiesInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.FiresTouchEvents
	out.Set("firesTouchEvents", value0)
	value1 := _this.PointerMovementScrolls
	out.Set("pointerMovementScrolls", value1)
	return out
}

// InputDeviceCapabilitiesInitFromJS is allocating a new
// InputDeviceCapabilitiesInit object and copy all values from
// input javascript object
func InputDeviceCapabilitiesInitFromJS(value js.Wrapper) *InputDeviceCapabilitiesInit {
	input := value.JSValue()
	var out InputDeviceCapabilitiesInit
	var (
		value0 bool // javascript: boolean {firesTouchEvents FiresTouchEvents firesTouchEvents}
		value1 bool // javascript: boolean {pointerMovementScrolls PointerMovementScrolls pointerMovementScrolls}
	)
	value0 = (input.Get("firesTouchEvents")).Bool()
	out.FiresTouchEvents = value0
	value1 = (input.Get("pointerMovementScrolls")).Bool()
	out.PointerMovementScrolls = value1
	return &out
}

// class: InputDeviceCapabilities
type InputDeviceCapabilities struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *InputDeviceCapabilities) JSValue() js.Value {
	return _this.Value_JS
}

// InputDeviceCapabilitiesFromJS is casting a js.Wrapper into InputDeviceCapabilities.
func InputDeviceCapabilitiesFromJS(value js.Wrapper) *InputDeviceCapabilities {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &InputDeviceCapabilities{}
	ret.Value_JS = input
	return ret
}

func NewInputDeviceCapabilities(deviceInitDict *InputDeviceCapabilitiesInit) (_result *InputDeviceCapabilities) {
	_klass := js.Global().Get("InputDeviceCapabilities")
	var (
		_args [1]interface{}
		_end  int
	)
	if deviceInitDict != nil {
		_p0 := deviceInitDict.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *InputDeviceCapabilities // javascript: InputDeviceCapabilities _what_return_name
	)
	_converted = InputDeviceCapabilitiesFromJS(_returned)
	_result = _converted
	return
}

// FiresTouchEvents returning attribute 'firesTouchEvents' with
// type bool (idl: boolean).
func (_this *InputDeviceCapabilities) FiresTouchEvents() bool {
	var ret bool
	value := _this.Value_JS.Get("firesTouchEvents")
	ret = (value).Bool()
	return ret
}

// PointerMovementScrolls returning attribute 'pointerMovementScrolls' with
// type bool (idl: boolean).
func (_this *InputDeviceCapabilities) PointerMovementScrolls() bool {
	var ret bool
	value := _this.Value_JS.Get("pointerMovementScrolls")
	ret = (value).Bool()
	return ret
}
