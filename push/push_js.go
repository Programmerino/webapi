// Code generated by webidl-bind. DO NOT EDIT.

package push

import "syscall/js"

import (
	"github.com/Programmerino/webapi/dom/domcore"
	"github.com/Programmerino/webapi/file"
	"github.com/Programmerino/webapi/javascript"
)

// using following types:
// domcore.ExtendableEvent
// file.Blob
// javascript.ArrayBuffer
// javascript.FrozenArray
// javascript.PromiseBool
// javascript.PromiseFinally

// source idl files:
// promises.idl
// push-api.idl

// transform files:
// promises.go.md
// push-api.go.md

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

// enum: PushEncryptionKeyName
type EncryptionKeyName int

const (
	P256dh EncryptionKeyName = iota
	Auth
)

var pushEncryptionKeyNameToWasmTable = []string{
	"p256dh", "auth",
}

var pushEncryptionKeyNameFromWasmTable = map[string]EncryptionKeyName{
	"p256dh": P256dh, "auth": Auth,
}

// JSValue is converting this enum into a javascript object
func (this *EncryptionKeyName) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this EncryptionKeyName) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(pushEncryptionKeyNameToWasmTable) {
		return pushEncryptionKeyNameToWasmTable[idx]
	}
	panic("unknown input value")
}

// EncryptionKeyNameFromJS is converting a javascript value into
// a EncryptionKeyName enum value.
func EncryptionKeyNameFromJS(value js.Value) EncryptionKeyName {
	key := value.String()
	conv, ok := pushEncryptionKeyNameFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: PushPermissionState
type PermissionState int

const (
	Denied PermissionState = iota
	Granted
	Prompt
)

var pushPermissionStateToWasmTable = []string{
	"denied", "granted", "prompt",
}

var pushPermissionStateFromWasmTable = map[string]PermissionState{
	"denied": Denied, "granted": Granted, "prompt": Prompt,
}

// JSValue is converting this enum into a javascript object
func (this *PermissionState) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this PermissionState) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(pushPermissionStateToWasmTable) {
		return pushPermissionStateToWasmTable[idx]
	}
	panic("unknown input value")
}

// PermissionStateFromJS is converting a javascript value into
// a PermissionState enum value.
func PermissionStateFromJS(value js.Value) PermissionState {
	key := value.String()
	conv, ok := pushPermissionStateFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: PromiseTemplateOnFulfilled
type PromiseNilSubscriptionOnFulfilledFunc func(value *Subscription)

// PromiseNilSubscriptionOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseNilSubscriptionOnFulfilled js.Func

func PromiseNilSubscriptionOnFulfilledToJS(callback PromiseNilSubscriptionOnFulfilledFunc) *PromiseNilSubscriptionOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseNilSubscriptionOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Subscription // javascript: PushSubscription value
		)
		_p0 = SubscriptionFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseNilSubscriptionOnFulfilledFromJS(_value js.Value) PromiseNilSubscriptionOnFulfilledFunc {
	return func(value *Subscription) {
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
type PromiseNilSubscriptionOnRejectedFunc func(reason js.Value)

// PromiseNilSubscriptionOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseNilSubscriptionOnRejected js.Func

func PromiseNilSubscriptionOnRejectedToJS(callback PromiseNilSubscriptionOnRejectedFunc) *PromiseNilSubscriptionOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseNilSubscriptionOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromiseNilSubscriptionOnRejectedFromJS(_value js.Value) PromiseNilSubscriptionOnRejectedFunc {
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

// callback: PromiseTemplateOnFulfilled
type PromisePermissionStateOnFulfilledFunc func(value PermissionState)

// PromisePermissionStateOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromisePermissionStateOnFulfilled js.Func

func PromisePermissionStateOnFulfilledToJS(callback PromisePermissionStateOnFulfilledFunc) *PromisePermissionStateOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromisePermissionStateOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 PermissionState // javascript: PushPermissionState value
		)
		_p0 = PermissionStateFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromisePermissionStateOnFulfilledFromJS(_value js.Value) PromisePermissionStateOnFulfilledFunc {
	return func(value PermissionState) {
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
type PromisePermissionStateOnRejectedFunc func(reason js.Value)

// PromisePermissionStateOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromisePermissionStateOnRejected js.Func

func PromisePermissionStateOnRejectedToJS(callback PromisePermissionStateOnRejectedFunc) *PromisePermissionStateOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromisePermissionStateOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromisePermissionStateOnRejectedFromJS(_value js.Value) PromisePermissionStateOnRejectedFunc {
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

// callback: PromiseTemplateOnFulfilled
type PromiseSubscriptionOnFulfilledFunc func(value *Subscription)

// PromiseSubscriptionOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseSubscriptionOnFulfilled js.Func

func PromiseSubscriptionOnFulfilledToJS(callback PromiseSubscriptionOnFulfilledFunc) *PromiseSubscriptionOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseSubscriptionOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Subscription // javascript: PushSubscription value
		)
		_p0 = SubscriptionFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseSubscriptionOnFulfilledFromJS(_value js.Value) PromiseSubscriptionOnFulfilledFunc {
	return func(value *Subscription) {
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
type PromiseSubscriptionOnRejectedFunc func(reason js.Value)

// PromiseSubscriptionOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseSubscriptionOnRejected js.Func

func PromiseSubscriptionOnRejectedToJS(callback PromiseSubscriptionOnRejectedFunc) *PromiseSubscriptionOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseSubscriptionOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromiseSubscriptionOnRejectedFromJS(_value js.Value) PromiseSubscriptionOnRejectedFunc {
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

// dictionary: PushEventInit
type EventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Data       *Union
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *EventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Data.JSValue()
	out.Set("data", value3)
	return out
}

// EventInitFromJS is allocating a new
// EventInit object and copy all values from
// input javascript object
func EventInitFromJS(value js.Wrapper) *EventInit {
	input := value.JSValue()
	var out EventInit
	var (
		value0 bool   // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool   // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool   // javascript: boolean {composed Composed composed}
		value3 *Union // javascript: Union {data Data data}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = UnionFromJS(input.Get("data"))
	out.Data = value3
	return &out
}

// dictionary: PushSubscriptionChangeInit
type SubscriptionChangeInit struct {
	Bubbles         bool
	Cancelable      bool
	Composed        bool
	NewSubscription *Subscription
	OldSubscription *Subscription
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *SubscriptionChangeInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.NewSubscription.JSValue()
	out.Set("newSubscription", value3)
	value4 := _this.OldSubscription.JSValue()
	out.Set("oldSubscription", value4)
	return out
}

// SubscriptionChangeInitFromJS is allocating a new
// SubscriptionChangeInit object and copy all values from
// input javascript object
func SubscriptionChangeInitFromJS(value js.Wrapper) *SubscriptionChangeInit {
	input := value.JSValue()
	var out SubscriptionChangeInit
	var (
		value0 bool          // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool          // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool          // javascript: boolean {composed Composed composed}
		value3 *Subscription // javascript: PushSubscription {newSubscription NewSubscription newSubscription}
		value4 *Subscription // javascript: PushSubscription {oldSubscription OldSubscription oldSubscription}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = SubscriptionFromJS(input.Get("newSubscription"))
	out.NewSubscription = value3
	value4 = SubscriptionFromJS(input.Get("oldSubscription"))
	out.OldSubscription = value4
	return &out
}

// dictionary: PushSubscriptionJSON
type SubscriptionJSON struct {
	Endpoint       string
	ExpirationTime *int
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *SubscriptionJSON) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Endpoint
	out.Set("endpoint", value0)

	var value1 interface{}
	if _this.ExpirationTime != nil {
		value1 = *(_this.ExpirationTime)
	} else {
		value1 = nil
	}
	out.Set("expirationTime", value1)
	return out
}

// SubscriptionJSONFromJS is allocating a new
// SubscriptionJSON object and copy all values from
// input javascript object
func SubscriptionJSONFromJS(value js.Wrapper) *SubscriptionJSON {
	input := value.JSValue()
	var out SubscriptionJSON
	var (
		value0 string // javascript: USVString {endpoint Endpoint endpoint}
		value1 *int   // javascript: unsigned long long {expirationTime ExpirationTime expirationTime}
	)
	value0 = (input.Get("endpoint")).String()
	out.Endpoint = value0
	if input.Get("expirationTime").Type() != js.TypeNull && input.Get("expirationTime").Type() != js.TypeUndefined {
		__tmp := (input.Get("expirationTime")).Int()
		value1 = &__tmp
	}
	out.ExpirationTime = value1
	return &out
}

// dictionary: PushSubscriptionOptionsInit
type SubscriptionOptionsInit struct {
	UserVisibleOnly      bool
	ApplicationServerKey *Union
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *SubscriptionOptionsInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.UserVisibleOnly
	out.Set("userVisibleOnly", value0)
	value1 := _this.ApplicationServerKey.JSValue()
	out.Set("applicationServerKey", value1)
	return out
}

// SubscriptionOptionsInitFromJS is allocating a new
// SubscriptionOptionsInit object and copy all values from
// input javascript object
func SubscriptionOptionsInitFromJS(value js.Wrapper) *SubscriptionOptionsInit {
	input := value.JSValue()
	var out SubscriptionOptionsInit
	var (
		value0 bool   // javascript: boolean {userVisibleOnly UserVisibleOnly userVisibleOnly}
		value1 *Union // javascript: Union {applicationServerKey ApplicationServerKey applicationServerKey}
	)
	value0 = (input.Get("userVisibleOnly")).Bool()
	out.UserVisibleOnly = value0
	if input.Get("applicationServerKey").Type() != js.TypeNull && input.Get("applicationServerKey").Type() != js.TypeUndefined {
		value1 = UnionFromJS(input.Get("applicationServerKey"))
	}
	out.ApplicationServerKey = value1
	return &out
}

// class: PushEvent
type Event struct {
	domcore.ExtendableEvent
}

// EventFromJS is casting a js.Wrapper into Event.
func EventFromJS(value js.Wrapper) *Event {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Event{}
	ret.Value_JS = input
	return ret
}

func NewPushEvent(_type string, eventInitDict *EventInit) (_result *Event) {
	_klass := js.Global().Get("PushEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Event // javascript: PushEvent _what_return_name
	)
	_converted = EventFromJS(_returned)
	_result = _converted
	return
}

// Data returning attribute 'data' with
// type MessageData (idl: PushMessageData).
func (_this *Event) Data() *MessageData {
	var ret *MessageData
	value := _this.Value_JS.Get("data")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = MessageDataFromJS(value)
	}
	return ret
}

// class: PushManager
type Manager struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Manager) JSValue() js.Value {
	return _this.Value_JS
}

// ManagerFromJS is casting a js.Wrapper into Manager.
func ManagerFromJS(value js.Wrapper) *Manager {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Manager{}
	ret.Value_JS = input
	return ret
}

// SupportedContentEncodings returning attribute 'supportedContentEncodings' with
// type javascript.FrozenArray (idl: FrozenArray).
func SupportedContentEncodings() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	_klass := js.Global().Get("PushManager")
	value := _klass.Get("supportedContentEncodings")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

func (_this *Manager) Subscribe(options *SubscriptionOptionsInit) (_result *PromiseSubscription) {
	var (
		_args [1]interface{}
		_end  int
	)
	if options != nil {
		_p0 := options.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _this.Value_JS.Call("subscribe", _args[0:_end]...)
	var (
		_converted *PromiseSubscription // javascript: Promise _what_return_name
	)
	_converted = PromiseSubscriptionFromJS(_returned)
	_result = _converted
	return
}

func (_this *Manager) GetSubscription() (_result *PromiseNilSubscription) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getSubscription", _args[0:_end]...)
	var (
		_converted *PromiseNilSubscription // javascript: Promise _what_return_name
	)
	_converted = PromiseNilSubscriptionFromJS(_returned)
	_result = _converted
	return
}

func (_this *Manager) PermissionState(options *SubscriptionOptionsInit) (_result *PromisePermissionState) {
	var (
		_args [1]interface{}
		_end  int
	)
	if options != nil {
		_p0 := options.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _this.Value_JS.Call("permissionState", _args[0:_end]...)
	var (
		_converted *PromisePermissionState // javascript: Promise _what_return_name
	)
	_converted = PromisePermissionStateFromJS(_returned)
	_result = _converted
	return
}

// class: PushMessageData
type MessageData struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *MessageData) JSValue() js.Value {
	return _this.Value_JS
}

// MessageDataFromJS is casting a js.Wrapper into MessageData.
func MessageDataFromJS(value js.Wrapper) *MessageData {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &MessageData{}
	ret.Value_JS = input
	return ret
}

func (_this *MessageData) ArrayBuffer() (_result *javascript.ArrayBuffer) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("arrayBuffer", _args[0:_end]...)
	var (
		_converted *javascript.ArrayBuffer // javascript: ArrayBuffer _what_return_name
	)
	_converted = javascript.ArrayBufferFromJS(_returned)
	_result = _converted
	return
}

func (_this *MessageData) Blob() (_result *file.Blob) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("blob", _args[0:_end]...)
	var (
		_converted *file.Blob // javascript: Blob _what_return_name
	)
	_converted = file.BlobFromJS(_returned)
	_result = _converted
	return
}

func (_this *MessageData) Json() (_result js.Value) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("json", _args[0:_end]...)
	var (
		_converted js.Value // javascript: any _what_return_name
	)
	_converted = _returned
	_result = _converted
	return
}

func (_this *MessageData) Text() (_result string) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("text", _args[0:_end]...)
	var (
		_converted string // javascript: USVString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

// class: Promise
type PromiseNilSubscription struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseNilSubscription) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseNilSubscriptionFromJS is casting a js.Wrapper into PromiseNilSubscription.
func PromiseNilSubscriptionFromJS(value js.Wrapper) *PromiseNilSubscription {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromiseNilSubscription{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseNilSubscription) Then(onFulfilled *PromiseNilSubscriptionOnFulfilled, onRejected *PromiseNilSubscriptionOnRejected) (_result *PromiseNilSubscription) {
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
		_converted *PromiseNilSubscription // javascript: Promise _what_return_name
	)
	_converted = PromiseNilSubscriptionFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseNilSubscription) Catch(onRejected *PromiseNilSubscriptionOnRejected) (_result *PromiseNilSubscription) {
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
		_converted *PromiseNilSubscription // javascript: Promise _what_return_name
	)
	_converted = PromiseNilSubscriptionFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseNilSubscription) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseNilSubscription) {
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
		_converted *PromiseNilSubscription // javascript: Promise _what_return_name
	)
	_converted = PromiseNilSubscriptionFromJS(_returned)
	_result = _converted
	return
}

// class: Promise
type PromisePermissionState struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromisePermissionState) JSValue() js.Value {
	return _this.Value_JS
}

// PromisePermissionStateFromJS is casting a js.Wrapper into PromisePermissionState.
func PromisePermissionStateFromJS(value js.Wrapper) *PromisePermissionState {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromisePermissionState{}
	ret.Value_JS = input
	return ret
}

func (_this *PromisePermissionState) Then(onFulfilled *PromisePermissionStateOnFulfilled, onRejected *PromisePermissionStateOnRejected) (_result *PromisePermissionState) {
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
		_converted *PromisePermissionState // javascript: Promise _what_return_name
	)
	_converted = PromisePermissionStateFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromisePermissionState) Catch(onRejected *PromisePermissionStateOnRejected) (_result *PromisePermissionState) {
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
		_converted *PromisePermissionState // javascript: Promise _what_return_name
	)
	_converted = PromisePermissionStateFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromisePermissionState) Finally(onFinally *javascript.PromiseFinally) (_result *PromisePermissionState) {
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
		_converted *PromisePermissionState // javascript: Promise _what_return_name
	)
	_converted = PromisePermissionStateFromJS(_returned)
	_result = _converted
	return
}

// class: Promise
type PromiseSubscription struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseSubscription) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseSubscriptionFromJS is casting a js.Wrapper into PromiseSubscription.
func PromiseSubscriptionFromJS(value js.Wrapper) *PromiseSubscription {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromiseSubscription{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseSubscription) Then(onFulfilled *PromiseSubscriptionOnFulfilled, onRejected *PromiseSubscriptionOnRejected) (_result *PromiseSubscription) {
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
		_converted *PromiseSubscription // javascript: Promise _what_return_name
	)
	_converted = PromiseSubscriptionFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseSubscription) Catch(onRejected *PromiseSubscriptionOnRejected) (_result *PromiseSubscription) {
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
		_converted *PromiseSubscription // javascript: Promise _what_return_name
	)
	_converted = PromiseSubscriptionFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseSubscription) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseSubscription) {
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
		_converted *PromiseSubscription // javascript: Promise _what_return_name
	)
	_converted = PromiseSubscriptionFromJS(_returned)
	_result = _converted
	return
}

// class: PushSubscription
type Subscription struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Subscription) JSValue() js.Value {
	return _this.Value_JS
}

// SubscriptionFromJS is casting a js.Wrapper into Subscription.
func SubscriptionFromJS(value js.Wrapper) *Subscription {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Subscription{}
	ret.Value_JS = input
	return ret
}

// Endpoint returning attribute 'endpoint' with
// type string (idl: USVString).
func (_this *Subscription) Endpoint() string {
	var ret string
	value := _this.Value_JS.Get("endpoint")
	ret = (value).String()
	return ret
}

// ExpirationTime returning attribute 'expirationTime' with
// type int (idl: unsigned long long).
func (_this *Subscription) ExpirationTime() *int {
	var ret *int
	value := _this.Value_JS.Get("expirationTime")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).Int()
		ret = &__tmp
	}
	return ret
}

// Options returning attribute 'options' with
// type SubscriptionOptions (idl: PushSubscriptionOptions).
func (_this *Subscription) Options() *SubscriptionOptions {
	var ret *SubscriptionOptions
	value := _this.Value_JS.Get("options")
	ret = SubscriptionOptionsFromJS(value)
	return ret
}

func (_this *Subscription) GetKey(name EncryptionKeyName) (_result *javascript.ArrayBuffer) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getKey", _args[0:_end]...)
	var (
		_converted *javascript.ArrayBuffer // javascript: ArrayBuffer _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = javascript.ArrayBufferFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *Subscription) Unsubscribe() (_result *javascript.PromiseBool) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("unsubscribe", _args[0:_end]...)
	var (
		_converted *javascript.PromiseBool // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseBoolFromJS(_returned)
	_result = _converted
	return
}

func (_this *Subscription) ToJSON() (_result *SubscriptionJSON) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("toJSON", _args[0:_end]...)
	var (
		_converted *SubscriptionJSON // javascript: PushSubscriptionJSON _what_return_name
	)
	_converted = SubscriptionJSONFromJS(_returned)
	_result = _converted
	return
}

// class: PushSubscriptionChangeEvent
type SubscriptionChangeEvent struct {
	domcore.ExtendableEvent
}

// SubscriptionChangeEventFromJS is casting a js.Wrapper into SubscriptionChangeEvent.
func SubscriptionChangeEventFromJS(value js.Wrapper) *SubscriptionChangeEvent {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &SubscriptionChangeEvent{}
	ret.Value_JS = input
	return ret
}

func NewPushSubscriptionChangeEvent(_type string, eventInitDict *SubscriptionChangeInit) (_result *SubscriptionChangeEvent) {
	_klass := js.Global().Get("PushSubscriptionChangeEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *SubscriptionChangeEvent // javascript: PushSubscriptionChangeEvent _what_return_name
	)
	_converted = SubscriptionChangeEventFromJS(_returned)
	_result = _converted
	return
}

// NewSubscription returning attribute 'newSubscription' with
// type Subscription (idl: PushSubscription).
func (_this *SubscriptionChangeEvent) NewSubscription() *Subscription {
	var ret *Subscription
	value := _this.Value_JS.Get("newSubscription")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = SubscriptionFromJS(value)
	}
	return ret
}

// OldSubscription returning attribute 'oldSubscription' with
// type Subscription (idl: PushSubscription).
func (_this *SubscriptionChangeEvent) OldSubscription() *Subscription {
	var ret *Subscription
	value := _this.Value_JS.Get("oldSubscription")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = SubscriptionFromJS(value)
	}
	return ret
}

// class: PushSubscriptionOptions
type SubscriptionOptions struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *SubscriptionOptions) JSValue() js.Value {
	return _this.Value_JS
}

// SubscriptionOptionsFromJS is casting a js.Wrapper into SubscriptionOptions.
func SubscriptionOptionsFromJS(value js.Wrapper) *SubscriptionOptions {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &SubscriptionOptions{}
	ret.Value_JS = input
	return ret
}

// UserVisibleOnly returning attribute 'userVisibleOnly' with
// type bool (idl: boolean).
func (_this *SubscriptionOptions) UserVisibleOnly() bool {
	var ret bool
	value := _this.Value_JS.Get("userVisibleOnly")
	ret = (value).Bool()
	return ret
}

// ApplicationServerKey returning attribute 'applicationServerKey' with
// type javascript.ArrayBuffer (idl: ArrayBuffer).
func (_this *SubscriptionOptions) ApplicationServerKey() *javascript.ArrayBuffer {
	var ret *javascript.ArrayBuffer
	value := _this.Value_JS.Get("applicationServerKey")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = javascript.ArrayBufferFromJS(value)
	}
	return ret
}
