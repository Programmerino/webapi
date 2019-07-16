// Code generated by webidl-bind. DO NOT EDIT.

package css

import "syscall/js"

// using following types:

// source idl files:
// css-properties-values-api.idl

// transform files:
// css-properties-values-api.go.md

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

// dictionary: PropertyDescriptor
type PropertyDescriptor struct {
	Name         string
	Syntax       string
	Inherits     bool
	InitialValue string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *PropertyDescriptor) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Name
	out.Set("name", value0)
	value1 := _this.Syntax
	out.Set("syntax", value1)
	value2 := _this.Inherits
	out.Set("inherits", value2)
	value3 := _this.InitialValue
	out.Set("initialValue", value3)
	return out
}

// PropertyDescriptorFromJS is allocating a new
// PropertyDescriptor object and copy all values from
// input javascript object
func PropertyDescriptorFromJS(value js.Wrapper) *PropertyDescriptor {
	input := value.JSValue()
	var out PropertyDescriptor
	var (
		value0 string // javascript: DOMString {name Name name}
		value1 string // javascript: DOMString {syntax Syntax syntax}
		value2 bool   // javascript: boolean {inherits Inherits inherits}
		value3 string // javascript: DOMString {initialValue InitialValue initialValue}
	)
	value0 = (input.Get("name")).String()
	out.Name = value0
	value1 = (input.Get("syntax")).String()
	out.Syntax = value1
	value2 = (input.Get("inherits")).Bool()
	out.Inherits = value2
	value3 = (input.Get("initialValue")).String()
	out.InitialValue = value3
	return &out
}