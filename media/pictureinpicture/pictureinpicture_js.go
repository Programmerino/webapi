// Code generated by webidl-bind. DO NOT EDIT.

package pictureinpicture

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget
// javascript.PromiseFinally

// source idl files:
// picture-in-picture.idl
// promises.idl

// transform files:
// picture-in-picture.go.md
// promises.go.md

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

// callback: PromiseTemplateOnFulfilled
type PromisePictureInPictureWindowOnFulfilledFunc func(value *PictureInPictureWindow)

// PromisePictureInPictureWindowOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromisePictureInPictureWindowOnFulfilled js.Func

func PromisePictureInPictureWindowOnFulfilledToJS(callback PromisePictureInPictureWindowOnFulfilledFunc) *PromisePictureInPictureWindowOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromisePictureInPictureWindowOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *PictureInPictureWindow // javascript: PictureInPictureWindow value
		)
		_p0 = PictureInPictureWindowFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromisePictureInPictureWindowOnFulfilledFromJS(_value js.Value) PromisePictureInPictureWindowOnFulfilledFunc {
	return func(value *PictureInPictureWindow) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := value.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnRejected
type PromisePictureInPictureWindowOnRejectedFunc func(reason js.Value)

// PromisePictureInPictureWindowOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromisePictureInPictureWindowOnRejected js.Func

func PromisePictureInPictureWindowOnRejectedToJS(callback PromisePictureInPictureWindowOnRejectedFunc) *PromisePictureInPictureWindowOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromisePictureInPictureWindowOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 js.Value // javascript: any reason
		)
		_p0 = args[0]
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromisePictureInPictureWindowOnRejectedFromJS(_value js.Value) PromisePictureInPictureWindowOnRejectedFunc {
	return func(reason js.Value) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := reason
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: EnterPictureInPictureEventInit
type EnterPictureInPictureEventInit struct {
	Bubbles                bool
	Cancelable             bool
	Composed               bool
	PictureInPictureWindow *PictureInPictureWindow
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *EnterPictureInPictureEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.PictureInPictureWindow.JSValue()
	out.Set("pictureInPictureWindow", value3)
	return out
}

// EnterPictureInPictureEventInitFromJS is allocating a new
// EnterPictureInPictureEventInit object and copy all values from
// input javascript object
func EnterPictureInPictureEventInitFromJS(value js.Wrapper) *EnterPictureInPictureEventInit {
	input := value.JSValue()
	var out EnterPictureInPictureEventInit
	var (
		value0 bool                    // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool                    // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool                    // javascript: boolean {composed Composed composed}
		value3 *PictureInPictureWindow // javascript: PictureInPictureWindow {pictureInPictureWindow PictureInPictureWindow pictureInPictureWindow}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = PictureInPictureWindowFromJS(input.Get("pictureInPictureWindow"))
	out.PictureInPictureWindow = value3
	return &out
}

// class: EnterPictureInPictureEvent
type EnterPictureInPictureEvent struct {
	domcore.Event
}

// EnterPictureInPictureEventFromJS is casting a js.Wrapper into EnterPictureInPictureEvent.
func EnterPictureInPictureEventFromJS(value js.Wrapper) *EnterPictureInPictureEvent {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &EnterPictureInPictureEvent{}
	ret.Value_JS = input
	return ret
}

func NewEnterPictureInPictureEvent(_type string, eventInitDict *EnterPictureInPictureEventInit) (_result *EnterPictureInPictureEvent) {
	_klass := js.Global().Get("EnterPictureInPictureEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_p1 := eventInitDict.JSValue()
	_args[1] = _p1
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *EnterPictureInPictureEvent // javascript: EnterPictureInPictureEvent _what_return_name
	)
	_converted = EnterPictureInPictureEventFromJS(_returned)
	_result = _converted
	return
}

// PictureInPictureWindow returning attribute 'pictureInPictureWindow' with
// type PictureInPictureWindow (idl: PictureInPictureWindow).
func (_this *EnterPictureInPictureEvent) PictureInPictureWindow() *PictureInPictureWindow {
	var ret *PictureInPictureWindow
	value := _this.Value_JS.Get("pictureInPictureWindow")
	ret = PictureInPictureWindowFromJS(value)
	return ret
}

// class: PictureInPictureWindow
type PictureInPictureWindow struct {
	domcore.EventTarget
}

// PictureInPictureWindowFromJS is casting a js.Wrapper into PictureInPictureWindow.
func PictureInPictureWindowFromJS(value js.Wrapper) *PictureInPictureWindow {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PictureInPictureWindow{}
	ret.Value_JS = input
	return ret
}

// Width returning attribute 'width' with
// type int (idl: long).
func (_this *PictureInPictureWindow) Width() int {
	var ret int
	value := _this.Value_JS.Get("width")
	ret = (value).Int()
	return ret
}

// Height returning attribute 'height' with
// type int (idl: long).
func (_this *PictureInPictureWindow) Height() int {
	var ret int
	value := _this.Value_JS.Get("height")
	ret = (value).Int()
	return ret
}

// OnResize returning attribute 'onresize' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *PictureInPictureWindow) OnResize() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onresize")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// event attribute: domcore.Event
func eventFuncPictureInPictureWindow_domcore_Event(listener func(event *domcore.Event, target *PictureInPictureWindow)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *domcore.Event
		value := args[0]
		incoming := value.Get("target")
		ret = domcore.EventFromJS(value)
		src := PictureInPictureWindowFromJS(incoming)
		listener(ret, src)
		return js.Undefined
	}
	return js.FuncOf(fn)
}

// AddResize is adding doing AddEventListener for 'Resize' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *PictureInPictureWindow) AddEventResize(listener func(event *domcore.Event, currentTarget *PictureInPictureWindow)) js.Func {
	cb := eventFuncPictureInPictureWindow_domcore_Event(listener)
	_this.Value_JS.Call("addEventListener", "resize", cb)
	return cb
}

// SetOnResize is assigning a function to 'onresize'. This
// This method is returning allocated javascript function that need to be released.
func (_this *PictureInPictureWindow) SetOnResize(listener func(event *domcore.Event, currentTarget *PictureInPictureWindow)) js.Func {
	cb := eventFuncPictureInPictureWindow_domcore_Event(listener)
	_this.Value_JS.Set("onresize", cb)
	return cb
}

// class: Promise
type PromisePictureInPictureWindow struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromisePictureInPictureWindow) JSValue() js.Value {
	return _this.Value_JS
}

// PromisePictureInPictureWindowFromJS is casting a js.Wrapper into PromisePictureInPictureWindow.
func PromisePictureInPictureWindowFromJS(value js.Wrapper) *PromisePictureInPictureWindow {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromisePictureInPictureWindow{}
	ret.Value_JS = input
	return ret
}

func (_this *PromisePictureInPictureWindow) Then(onFulfilled *PromisePictureInPictureWindowOnFulfilled, onRejected *PromisePictureInPictureWindowOnRejected) (_result *PromisePictureInPictureWindow) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFulfilled != nil {
		__callback0 = (*onFulfilled).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if onRejected != nil {

		var __callback1 js.Value
		if onRejected != nil {
			__callback1 = (*onRejected).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("then", _args[0:_end]...)
	var (
		_converted *PromisePictureInPictureWindow // javascript: Promise _what_return_name
	)
	_converted = PromisePictureInPictureWindowFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromisePictureInPictureWindow) Catch(onRejected *PromisePictureInPictureWindowOnRejected) (_result *PromisePictureInPictureWindow) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onRejected != nil {
		__callback0 = (*onRejected).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("catch", _args[0:_end]...)
	var (
		_converted *PromisePictureInPictureWindow // javascript: Promise _what_return_name
	)
	_converted = PromisePictureInPictureWindowFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromisePictureInPictureWindow) Finally(onFinally *javascript.PromiseFinally) (_result *PromisePictureInPictureWindow) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFinally != nil {
		__callback0 = (*onFinally).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("finally", _args[0:_end]...)
	var (
		_converted *PromisePictureInPictureWindow // javascript: Promise _what_return_name
	)
	_converted = PromisePictureInPictureWindowFromJS(_returned)
	_result = _converted
	return
}
