// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package patch

import js "github.com/gowebapi/webapi/core/failjs"

// using following types:

// ReleasableApiResource is used to release underlaying
// allocated resources.
type ReleasableApiResource interface {
	Release()
}

type releasableApiResourceList []ReleasableApiResource

func (a releasableApiResourceList) Release() {
	for _, v := range a {
		v.Release()
	}
}

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

// interface: ByteString
type ByteString struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *ByteString) JSValue() js.Value {
	return _this.Value_JS
}

// ByteStringFromJS is casting a js.Wrapper into ByteString.
func ByteStringFromJS(value js.Wrapper) *ByteString {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ByteString{}
	ret.Value_JS = input
	return ret
}

// interface: MediaSource
type MediaSource struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *MediaSource) JSValue() js.Value {
	return _this.Value_JS
}

// MediaSourceFromJS is casting a js.Wrapper into MediaSource.
func MediaSourceFromJS(value js.Wrapper) *MediaSource {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MediaSource{}
	ret.Value_JS = input
	return ret
}

// interface: MediaStream
type MediaStream struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *MediaStream) JSValue() js.Value {
	return _this.Value_JS
}

// MediaStreamFromJS is casting a js.Wrapper into MediaStream.
func MediaStreamFromJS(value js.Wrapper) *MediaStream {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MediaStream{}
	ret.Value_JS = input
	return ret
}

// interface: ReadableStream
type ReadableStream struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *ReadableStream) JSValue() js.Value {
	return _this.Value_JS
}

// ReadableStreamFromJS is casting a js.Wrapper into ReadableStream.
func ReadableStreamFromJS(value js.Wrapper) *ReadableStream {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ReadableStream{}
	ret.Value_JS = input
	return ret
}

// interface: SVGImageElement
type SVGImageElement struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *SVGImageElement) JSValue() js.Value {
	return _this.Value_JS
}

// SVGImageElementFromJS is casting a js.Wrapper into SVGImageElement.
func SVGImageElementFromJS(value js.Wrapper) *SVGImageElement {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &SVGImageElement{}
	ret.Value_JS = input
	return ret
}

// interface: SVGScriptElement
type SVGScriptElement struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *SVGScriptElement) JSValue() js.Value {
	return _this.Value_JS
}

// SVGScriptElementFromJS is casting a js.Wrapper into SVGScriptElement.
func SVGScriptElementFromJS(value js.Wrapper) *SVGScriptElement {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &SVGScriptElement{}
	ret.Value_JS = input
	return ret
}

// interface: Uint8ClampedArray
type Uint8ClampedArray struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Uint8ClampedArray) JSValue() js.Value {
	return _this.Value_JS
}

// Uint8ClampedArrayFromJS is casting a js.Wrapper into Uint8ClampedArray.
func Uint8ClampedArrayFromJS(value js.Wrapper) *Uint8ClampedArray {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Uint8ClampedArray{}
	ret.Value_JS = input
	return ret
}
