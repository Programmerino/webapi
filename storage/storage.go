// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package storage

import js "github.com/Programmerino/webapi/core/js"

import (
	"github.com/Programmerino/webapi/javascript"
)

// using following types:
// javascript.PromiseBool
// javascript.PromiseFinally

// source idl files:
// promises.idl
// storage.idl

// transform files:
// promises.go.md
// storage.go.md

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
type PromiseStorageEstimateOnFulfilledFunc func(value *StorageEstimate)

// PromiseStorageEstimateOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseStorageEstimateOnFulfilled js.Func

func PromiseStorageEstimateOnFulfilledToJS(callback PromiseStorageEstimateOnFulfilledFunc) *PromiseStorageEstimateOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseStorageEstimateOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *StorageEstimate // javascript: StorageEstimate value
		)
		_p0 = StorageEstimateFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseStorageEstimateOnFulfilledFromJS(_value js.Value) PromiseStorageEstimateOnFulfilledFunc {
	return func(value *StorageEstimate) {
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
type PromiseStorageEstimateOnRejectedFunc func(reason js.Value)

// PromiseStorageEstimateOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseStorageEstimateOnRejected js.Func

func PromiseStorageEstimateOnRejectedToJS(callback PromiseStorageEstimateOnRejectedFunc) *PromiseStorageEstimateOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseStorageEstimateOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromiseStorageEstimateOnRejectedFromJS(_value js.Value) PromiseStorageEstimateOnRejectedFunc {
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

// dictionary: StorageEstimate
type StorageEstimate struct {
	Usage int
	Quota int
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *StorageEstimate) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Usage
	out.Set("usage", value0)
	value1 := _this.Quota
	out.Set("quota", value1)
	return out
}

// StorageEstimateFromJS is allocating a new
// StorageEstimate object and copy all values from
// input javascript object
func StorageEstimateFromJS(value js.Wrapper) *StorageEstimate {
	input := value.JSValue()
	var out StorageEstimate
	var (
		value0 int // javascript: unsigned long long {usage Usage usage}
		value1 int // javascript: unsigned long long {quota Quota quota}
	)
	value0 = (input.Get("usage")).Int()
	out.Usage = value0
	value1 = (input.Get("quota")).Int()
	out.Quota = value1
	return &out
}

// class: Promise
type PromiseStorageEstimate struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseStorageEstimate) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseStorageEstimateFromJS is casting a js.Wrapper into PromiseStorageEstimate.
func PromiseStorageEstimateFromJS(value js.Wrapper) *PromiseStorageEstimate {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromiseStorageEstimate{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseStorageEstimate) Then(onFulfilled *PromiseStorageEstimateOnFulfilled, onRejected *PromiseStorageEstimateOnRejected) (_result *PromiseStorageEstimate) {
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
		_converted *PromiseStorageEstimate // javascript: Promise _what_return_name
	)
	_converted = PromiseStorageEstimateFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseStorageEstimate) Catch(onRejected *PromiseStorageEstimateOnRejected) (_result *PromiseStorageEstimate) {
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
		_converted *PromiseStorageEstimate // javascript: Promise _what_return_name
	)
	_converted = PromiseStorageEstimateFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseStorageEstimate) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseStorageEstimate) {
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
		_converted *PromiseStorageEstimate // javascript: Promise _what_return_name
	)
	_converted = PromiseStorageEstimateFromJS(_returned)
	_result = _converted
	return
}

// class: StorageManager
type StorageManager struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *StorageManager) JSValue() js.Value {
	return _this.Value_JS
}

// StorageManagerFromJS is casting a js.Wrapper into StorageManager.
func StorageManagerFromJS(value js.Wrapper) *StorageManager {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &StorageManager{}
	ret.Value_JS = input
	return ret
}

func (_this *StorageManager) Persisted() (_result *javascript.PromiseBool) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("persisted", _args[0:_end]...)
	var (
		_converted *javascript.PromiseBool // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseBoolFromJS(_returned)
	_result = _converted
	return
}

func (_this *StorageManager) Persist() (_result *javascript.PromiseBool) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("persist", _args[0:_end]...)
	var (
		_converted *javascript.PromiseBool // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseBoolFromJS(_returned)
	_result = _converted
	return
}

func (_this *StorageManager) Estimate() (_result *PromiseStorageEstimate) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("estimate", _args[0:_end]...)
	var (
		_converted *PromiseStorageEstimate // javascript: Promise _what_return_name
	)
	_converted = PromiseStorageEstimateFromJS(_returned)
	_result = _converted
	return
}
