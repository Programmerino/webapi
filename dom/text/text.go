// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package text

import js "github.com/Programmerino/webapi/core/js"

import (
	"github.com/Programmerino/webapi/javascript"
	"github.com/Programmerino/webapi/javascript/missingtypes"
	"github.com/Programmerino/webapi/patch"
)

// using following types:
// javascript.Uint8Array
// missingtypes.WritableStream
// patch.ReadableStream

// source idl files:
// encoding.idl

// transform files:
// encoding.go.md

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

// dictionary: TextDecodeOptions
type DecodeOptions struct {
	Stream bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *DecodeOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Stream
	out.Set("stream", value0)
	return out
}

// DecodeOptionsFromJS is allocating a new
// DecodeOptions object and copy all values from
// input javascript object
func DecodeOptionsFromJS(value js.Wrapper) *DecodeOptions {
	input := value.JSValue()
	var out DecodeOptions
	var (
		value0 bool // javascript: boolean {stream Stream stream}
	)
	value0 = (input.Get("stream")).Bool()
	out.Stream = value0
	return &out
}

// dictionary: TextDecoderOptions
type DecoderOptions struct {
	Fatal     bool
	IgnoreBOM bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *DecoderOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Fatal
	out.Set("fatal", value0)
	value1 := _this.IgnoreBOM
	out.Set("ignoreBOM", value1)
	return out
}

// DecoderOptionsFromJS is allocating a new
// DecoderOptions object and copy all values from
// input javascript object
func DecoderOptionsFromJS(value js.Wrapper) *DecoderOptions {
	input := value.JSValue()
	var out DecoderOptions
	var (
		value0 bool // javascript: boolean {fatal Fatal fatal}
		value1 bool // javascript: boolean {ignoreBOM IgnoreBOM ignoreBOM}
	)
	value0 = (input.Get("fatal")).Bool()
	out.Fatal = value0
	value1 = (input.Get("ignoreBOM")).Bool()
	out.IgnoreBOM = value1
	return &out
}

// dictionary: TextEncoderEncodeIntoResult
type EncoderEncodeIntoResult struct {
	Read    int
	Written int
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *EncoderEncodeIntoResult) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Read
	out.Set("read", value0)
	value1 := _this.Written
	out.Set("written", value1)
	return out
}

// EncoderEncodeIntoResultFromJS is allocating a new
// EncoderEncodeIntoResult object and copy all values from
// input javascript object
func EncoderEncodeIntoResultFromJS(value js.Wrapper) *EncoderEncodeIntoResult {
	input := value.JSValue()
	var out EncoderEncodeIntoResult
	var (
		value0 int // javascript: unsigned long long {read Read read}
		value1 int // javascript: unsigned long long {written Written written}
	)
	value0 = (input.Get("read")).Int()
	out.Read = value0
	value1 = (input.Get("written")).Int()
	out.Written = value1
	return &out
}

// class: TextDecoder
type Decoder struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Decoder) JSValue() js.Value {
	return _this.Value_JS
}

// DecoderFromJS is casting a js.Wrapper into Decoder.
func DecoderFromJS(value js.Wrapper) *Decoder {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Decoder{}
	ret.Value_JS = input
	return ret
}

func NewTextDecoder(label *string, options *DecoderOptions) (_result *Decoder) {
	_klass := js.Global().Get("TextDecoder")
	var (
		_args [2]interface{}
		_end  int
	)
	if label != nil {

		var _p0 interface{}
		if label != nil {
			_p0 = *(label)
		} else {
			_p0 = nil
		}
		_args[0] = _p0
		_end++
	}
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Decoder // javascript: TextDecoder _what_return_name
	)
	_converted = DecoderFromJS(_returned)
	_result = _converted
	return
}

// Encoding returning attribute 'encoding' with
// type string (idl: DOMString).
func (_this *Decoder) Encoding() string {
	var ret string
	value := _this.Value_JS.Get("encoding")
	ret = (value).String()
	return ret
}

// Fatal returning attribute 'fatal' with
// type bool (idl: boolean).
func (_this *Decoder) Fatal() bool {
	var ret bool
	value := _this.Value_JS.Get("fatal")
	ret = (value).Bool()
	return ret
}

// IgnoreBOM returning attribute 'ignoreBOM' with
// type bool (idl: boolean).
func (_this *Decoder) IgnoreBOM() bool {
	var ret bool
	value := _this.Value_JS.Get("ignoreBOM")
	ret = (value).Bool()
	return ret
}

func (_this *Decoder) Decode(input *Union, options *DecodeOptions) (_result string) {
	var (
		_args [2]interface{}
		_end  int
	)
	if input != nil {
		_p0 := input.JSValue()
		_args[0] = _p0
		_end++
	}
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("decode", _args[0:_end]...)
	var (
		_converted string // javascript: USVString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

// class: TextDecoderStream
type DecoderStream struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DecoderStream) JSValue() js.Value {
	return _this.Value_JS
}

// DecoderStreamFromJS is casting a js.Wrapper into DecoderStream.
func DecoderStreamFromJS(value js.Wrapper) *DecoderStream {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &DecoderStream{}
	ret.Value_JS = input
	return ret
}

func NewTextDecoderStream(label *string, options *DecoderOptions) (_result *DecoderStream) {
	_klass := js.Global().Get("TextDecoderStream")
	var (
		_args [2]interface{}
		_end  int
	)
	if label != nil {

		var _p0 interface{}
		if label != nil {
			_p0 = *(label)
		} else {
			_p0 = nil
		}
		_args[0] = _p0
		_end++
	}
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *DecoderStream // javascript: TextDecoderStream _what_return_name
	)
	_converted = DecoderStreamFromJS(_returned)
	_result = _converted
	return
}

// Encoding returning attribute 'encoding' with
// type string (idl: DOMString).
func (_this *DecoderStream) Encoding() string {
	var ret string
	value := _this.Value_JS.Get("encoding")
	ret = (value).String()
	return ret
}

// Fatal returning attribute 'fatal' with
// type bool (idl: boolean).
func (_this *DecoderStream) Fatal() bool {
	var ret bool
	value := _this.Value_JS.Get("fatal")
	ret = (value).Bool()
	return ret
}

// IgnoreBOM returning attribute 'ignoreBOM' with
// type bool (idl: boolean).
func (_this *DecoderStream) IgnoreBOM() bool {
	var ret bool
	value := _this.Value_JS.Get("ignoreBOM")
	ret = (value).Bool()
	return ret
}

// Readable returning attribute 'readable' with
// type patch.ReadableStream (idl: ReadableStream).
func (_this *DecoderStream) Readable() *patch.ReadableStream {
	var ret *patch.ReadableStream
	value := _this.Value_JS.Get("readable")
	ret = patch.ReadableStreamFromJS(value)
	return ret
}

// Writable returning attribute 'writable' with
// type missingtypes.WritableStream (idl: WritableStream).
func (_this *DecoderStream) Writable() *missingtypes.WritableStream {
	var ret *missingtypes.WritableStream
	value := _this.Value_JS.Get("writable")
	ret = missingtypes.WritableStreamFromJS(value)
	return ret
}

// class: TextEncoder
type Encoder struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Encoder) JSValue() js.Value {
	return _this.Value_JS
}

// EncoderFromJS is casting a js.Wrapper into Encoder.
func EncoderFromJS(value js.Wrapper) *Encoder {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Encoder{}
	ret.Value_JS = input
	return ret
}

func NewTextEncoder() (_result *Encoder) {
	_klass := js.Global().Get("TextEncoder")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Encoder // javascript: TextEncoder _what_return_name
	)
	_converted = EncoderFromJS(_returned)
	_result = _converted
	return
}

// Encoding returning attribute 'encoding' with
// type string (idl: DOMString).
func (_this *Encoder) Encoding() string {
	var ret string
	value := _this.Value_JS.Get("encoding")
	ret = (value).String()
	return ret
}

func (_this *Encoder) Encode(input *string) (_result *javascript.Uint8Array) {
	var (
		_args [1]interface{}
		_end  int
	)
	if input != nil {

		var _p0 interface{}
		if input != nil {
			_p0 = *(input)
		} else {
			_p0 = nil
		}
		_args[0] = _p0
		_end++
	}
	_returned := _this.Value_JS.Call("encode", _args[0:_end]...)
	var (
		_converted *javascript.Uint8Array // javascript: Uint8Array _what_return_name
	)
	_converted = javascript.Uint8ArrayFromJS(_returned)
	_result = _converted
	return
}

func (_this *Encoder) EncodeInto(source string, destination *javascript.Uint8Array) (_result *EncoderEncodeIntoResult) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := source
	_args[0] = _p0
	_end++
	_p1 := destination.JSValue()
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("encodeInto", _args[0:_end]...)
	var (
		_converted *EncoderEncodeIntoResult // javascript: TextEncoderEncodeIntoResult _what_return_name
	)
	_converted = EncoderEncodeIntoResultFromJS(_returned)
	_result = _converted
	return
}

// class: TextEncoderStream
type EncoderStream struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *EncoderStream) JSValue() js.Value {
	return _this.Value_JS
}

// EncoderStreamFromJS is casting a js.Wrapper into EncoderStream.
func EncoderStreamFromJS(value js.Wrapper) *EncoderStream {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &EncoderStream{}
	ret.Value_JS = input
	return ret
}

func NewTextEncoderStream() (_result *EncoderStream) {
	_klass := js.Global().Get("TextEncoderStream")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *EncoderStream // javascript: TextEncoderStream _what_return_name
	)
	_converted = EncoderStreamFromJS(_returned)
	_result = _converted
	return
}

// Encoding returning attribute 'encoding' with
// type string (idl: DOMString).
func (_this *EncoderStream) Encoding() string {
	var ret string
	value := _this.Value_JS.Get("encoding")
	ret = (value).String()
	return ret
}

// Readable returning attribute 'readable' with
// type patch.ReadableStream (idl: ReadableStream).
func (_this *EncoderStream) Readable() *patch.ReadableStream {
	var ret *patch.ReadableStream
	value := _this.Value_JS.Get("readable")
	ret = patch.ReadableStreamFromJS(value)
	return ret
}

// Writable returning attribute 'writable' with
// type missingtypes.WritableStream (idl: WritableStream).
func (_this *EncoderStream) Writable() *missingtypes.WritableStream {
	var ret *missingtypes.WritableStream
	value := _this.Value_JS.Get("writable")
	ret = missingtypes.WritableStreamFromJS(value)
	return ret
}
