// Code generated by webidl-bind. DO NOT EDIT.

package webvr

import "syscall/js"

import (
	"github.com/Programmerino/webapi/core/jsarray"
	"github.com/Programmerino/webapi/dom/domcore"
	"github.com/Programmerino/webapi/html/htmlcommon"
	"github.com/Programmerino/webapi/javascript"
)

// using following types:
// domcore.Event
// domcore.EventTarget
// htmlcommon.FrameRequestCallback
// javascript.Float32Array
// javascript.PromiseFinally
// javascript.PromiseVoid

// source idl files:
// promises.idl
// webvr.idl

// transform files:
// promises.go.md
// webvr.go.md

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

// enum: VRDisplayEventReason
type DisplayEventReason int

const (
	Mounted DisplayEventReason = iota
	Navigation
	Requested
	Unmounted
)

var vRDisplayEventReasonToWasmTable = []string{
	"mounted", "navigation", "requested", "unmounted",
}

var vRDisplayEventReasonFromWasmTable = map[string]DisplayEventReason{
	"mounted": Mounted, "navigation": Navigation, "requested": Requested, "unmounted": Unmounted,
}

// JSValue is converting this enum into a javascript object
func (this *DisplayEventReason) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this DisplayEventReason) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(vRDisplayEventReasonToWasmTable) {
		return vRDisplayEventReasonToWasmTable[idx]
	}
	panic("unknown input value")
}

// DisplayEventReasonFromJS is converting a javascript value into
// a DisplayEventReason enum value.
func DisplayEventReasonFromJS(value js.Value) DisplayEventReason {
	key := value.String()
	conv, ok := vRDisplayEventReasonFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: VREye
type Eye int

const (
	Left Eye = iota
	Right
)

var vREyeToWasmTable = []string{
	"left", "right",
}

var vREyeFromWasmTable = map[string]Eye{
	"left": Left, "right": Right,
}

// JSValue is converting this enum into a javascript object
func (this *Eye) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this Eye) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(vREyeToWasmTable) {
		return vREyeToWasmTable[idx]
	}
	panic("unknown input value")
}

// EyeFromJS is converting a javascript value into
// a Eye enum value.
func EyeFromJS(value js.Value) Eye {
	key := value.String()
	conv, ok := vREyeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: PromiseTemplateOnFulfilled
type PromiseSequenceDisplayOnFulfilledFunc func(value []*Display)

// PromiseSequenceDisplayOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseSequenceDisplayOnFulfilled js.Func

func PromiseSequenceDisplayOnFulfilledToJS(callback PromiseSequenceDisplayOnFulfilledFunc) *PromiseSequenceDisplayOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseSequenceDisplayOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 []*Display // javascript: sequence<VRDisplay> value
		)
		__length0 := args[0].Length()
		__array0 := make([]*Display, __length0, __length0)
		for __idx0 := 0; __idx0 < __length0; __idx0++ {
			var __seq_out0 *Display
			__seq_in0 := args[0].Index(__idx0)
			__seq_out0 = DisplayFromJS(__seq_in0)
			__array0[__idx0] = __seq_out0
		}
		_p0 = __array0
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseSequenceDisplayOnFulfilledFromJS(_value js.Value) PromiseSequenceDisplayOnFulfilledFunc {
	return func(value []*Display) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := js.Global().Get("Array").New(len(value))
		for __idx0, __seq_in0 := range value {
			__seq_out0 := __seq_in0.JSValue()
			_p0.SetIndex(__idx0, __seq_out0)
		}
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnRejected
type PromiseSequenceDisplayOnRejectedFunc func(reason js.Value)

// PromiseSequenceDisplayOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseSequenceDisplayOnRejected js.Func

func PromiseSequenceDisplayOnRejectedToJS(callback PromiseSequenceDisplayOnRejectedFunc) *PromiseSequenceDisplayOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseSequenceDisplayOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromiseSequenceDisplayOnRejectedFromJS(_value js.Value) PromiseSequenceDisplayOnRejectedFunc {
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

// dictionary: VRDisplayEventInit
type DisplayEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Display    *Display
	Reason     DisplayEventReason
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *DisplayEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Display.JSValue()
	out.Set("display", value3)
	value4 := _this.Reason.JSValue()
	out.Set("reason", value4)
	return out
}

// DisplayEventInitFromJS is allocating a new
// DisplayEventInit object and copy all values from
// input javascript object
func DisplayEventInitFromJS(value js.Wrapper) *DisplayEventInit {
	input := value.JSValue()
	var out DisplayEventInit
	var (
		value0 bool               // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool               // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool               // javascript: boolean {composed Composed composed}
		value3 *Display           // javascript: VRDisplay {display Display display}
		value4 DisplayEventReason // javascript: VRDisplayEventReason {reason Reason reason}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = DisplayFromJS(input.Get("display"))
	out.Display = value3
	value4 = DisplayEventReasonFromJS(input.Get("reason"))
	out.Reason = value4
	return &out
}

// dictionary: VRLayerInit
type LayerInit struct {
	Source      *Union
	LeftBounds  []float32
	RightBounds []float32
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *LayerInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Source.JSValue()
	out.Set("source", value0)
	value1 := jsarray.Float32ToJS(_this.LeftBounds)
	out.Set("leftBounds", value1)
	value2 := jsarray.Float32ToJS(_this.RightBounds)
	out.Set("rightBounds", value2)
	return out
}

// LayerInitFromJS is allocating a new
// LayerInit object and copy all values from
// input javascript object
func LayerInitFromJS(value js.Wrapper) *LayerInit {
	input := value.JSValue()
	var out LayerInit
	var (
		value0 *Union    // javascript: Union {source Source source}
		value1 []float32 // javascript: typed-array {leftBounds LeftBounds leftBounds}
		value2 []float32 // javascript: typed-array {rightBounds RightBounds rightBounds}
	)
	if input.Get("source").Type() != js.TypeNull && input.Get("source").Type() != js.TypeUndefined {
		value0 = UnionFromJS(input.Get("source"))
	}
	out.Source = value0
	value1 = jsarray.Float32ToGo(input.Get("leftBounds"))
	out.LeftBounds = value1
	value2 = jsarray.Float32ToGo(input.Get("rightBounds"))
	out.RightBounds = value2
	return &out
}

// class: VRDisplay
type Display struct {
	domcore.EventTarget
}

// DisplayFromJS is casting a js.Wrapper into Display.
func DisplayFromJS(value js.Wrapper) *Display {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Display{}
	ret.Value_JS = input
	return ret
}

// IsConnected returning attribute 'isConnected' with
// type bool (idl: boolean).
func (_this *Display) IsConnected() bool {
	var ret bool
	value := _this.Value_JS.Get("isConnected")
	ret = (value).Bool()
	return ret
}

// IsPresenting returning attribute 'isPresenting' with
// type bool (idl: boolean).
func (_this *Display) IsPresenting() bool {
	var ret bool
	value := _this.Value_JS.Get("isPresenting")
	ret = (value).Bool()
	return ret
}

// Capabilities returning attribute 'capabilities' with
// type DisplayCapabilities (idl: VRDisplayCapabilities).
func (_this *Display) Capabilities() *DisplayCapabilities {
	var ret *DisplayCapabilities
	value := _this.Value_JS.Get("capabilities")
	ret = DisplayCapabilitiesFromJS(value)
	return ret
}

// StageParameters returning attribute 'stageParameters' with
// type StageParameters (idl: VRStageParameters).
func (_this *Display) StageParameters() *StageParameters {
	var ret *StageParameters
	value := _this.Value_JS.Get("stageParameters")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = StageParametersFromJS(value)
	}
	return ret
}

// DisplayId returning attribute 'displayId' with
// type uint (idl: unsigned long).
func (_this *Display) DisplayId() uint {
	var ret uint
	value := _this.Value_JS.Get("displayId")
	ret = (uint)((value).Int())
	return ret
}

// DisplayName returning attribute 'displayName' with
// type string (idl: DOMString).
func (_this *Display) DisplayName() string {
	var ret string
	value := _this.Value_JS.Get("displayName")
	ret = (value).String()
	return ret
}

// DepthNear returning attribute 'depthNear' with
// type float64 (idl: double).
func (_this *Display) DepthNear() float64 {
	var ret float64
	value := _this.Value_JS.Get("depthNear")
	ret = (value).Float()
	return ret
}

// SetDepthNear setting attribute 'depthNear' with
// type float64 (idl: double).
func (_this *Display) SetDepthNear(value float64) {
	input := value
	_this.Value_JS.Set("depthNear", input)
}

// DepthFar returning attribute 'depthFar' with
// type float64 (idl: double).
func (_this *Display) DepthFar() float64 {
	var ret float64
	value := _this.Value_JS.Get("depthFar")
	ret = (value).Float()
	return ret
}

// SetDepthFar setting attribute 'depthFar' with
// type float64 (idl: double).
func (_this *Display) SetDepthFar(value float64) {
	input := value
	_this.Value_JS.Set("depthFar", input)
}

func (_this *Display) GetEyeParameters(whichEye Eye) (_result *EyeParameters) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := whichEye.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getEyeParameters", _args[0:_end]...)
	var (
		_converted *EyeParameters // javascript: VREyeParameters _what_return_name
	)
	_converted = EyeParametersFromJS(_returned)
	_result = _converted
	return
}

func (_this *Display) GetFrameData(frameData *FrameData) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := frameData.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getFrameData", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *Display) GetPose() (_result *Pose) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getPose", _args[0:_end]...)
	var (
		_converted *Pose // javascript: VRPose _what_return_name
	)
	_converted = PoseFromJS(_returned)
	_result = _converted
	return
}

func (_this *Display) ResetPose() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("resetPose", _args[0:_end]...)
	return
}

func (_this *Display) RequestAnimationFrame(callback *htmlcommon.FrameRequestCallback) (_result int) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if callback != nil {
		__callback0 = (*callback).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("requestAnimationFrame", _args[0:_end]...)
	var (
		_converted int // javascript: long _what_return_name
	)
	_converted = (_returned).Int()
	_result = _converted
	return
}

func (_this *Display) CancelAnimationFrame(handle int) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := handle
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("cancelAnimationFrame", _args[0:_end]...)
	return
}

func (_this *Display) RequestPresent(layers []*LayerInit) (_result *javascript.PromiseVoid) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := js.Global().Get("Array").New(len(layers))
	for __idx0, __seq_in0 := range layers {
		__seq_out0 := __seq_in0.JSValue()
		_p0.SetIndex(__idx0, __seq_out0)
	}
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("requestPresent", _args[0:_end]...)
	var (
		_converted *javascript.PromiseVoid // javascript: PromiseVoid _what_return_name
	)
	_converted = javascript.PromiseVoidFromJS(_returned)
	_result = _converted
	return
}

func (_this *Display) ExitPresent() (_result *javascript.PromiseVoid) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("exitPresent", _args[0:_end]...)
	var (
		_converted *javascript.PromiseVoid // javascript: PromiseVoid _what_return_name
	)
	_converted = javascript.PromiseVoidFromJS(_returned)
	_result = _converted
	return
}

func (_this *Display) GetLayers() (_result []*LayerInit) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getLayers", _args[0:_end]...)
	var (
		_converted []*LayerInit // javascript: sequence<VRLayerInit> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*LayerInit, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *LayerInit
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = LayerInitFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *Display) SubmitFrame() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("submitFrame", _args[0:_end]...)
	return
}

// class: VRDisplayCapabilities
type DisplayCapabilities struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DisplayCapabilities) JSValue() js.Value {
	return _this.Value_JS
}

// DisplayCapabilitiesFromJS is casting a js.Wrapper into DisplayCapabilities.
func DisplayCapabilitiesFromJS(value js.Wrapper) *DisplayCapabilities {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &DisplayCapabilities{}
	ret.Value_JS = input
	return ret
}

// HasPosition returning attribute 'hasPosition' with
// type bool (idl: boolean).
func (_this *DisplayCapabilities) HasPosition() bool {
	var ret bool
	value := _this.Value_JS.Get("hasPosition")
	ret = (value).Bool()
	return ret
}

// HasOrientation returning attribute 'hasOrientation' with
// type bool (idl: boolean).
func (_this *DisplayCapabilities) HasOrientation() bool {
	var ret bool
	value := _this.Value_JS.Get("hasOrientation")
	ret = (value).Bool()
	return ret
}

// HasExternalDisplay returning attribute 'hasExternalDisplay' with
// type bool (idl: boolean).
func (_this *DisplayCapabilities) HasExternalDisplay() bool {
	var ret bool
	value := _this.Value_JS.Get("hasExternalDisplay")
	ret = (value).Bool()
	return ret
}

// CanPresent returning attribute 'canPresent' with
// type bool (idl: boolean).
func (_this *DisplayCapabilities) CanPresent() bool {
	var ret bool
	value := _this.Value_JS.Get("canPresent")
	ret = (value).Bool()
	return ret
}

// MaxLayers returning attribute 'maxLayers' with
// type uint (idl: unsigned long).
func (_this *DisplayCapabilities) MaxLayers() uint {
	var ret uint
	value := _this.Value_JS.Get("maxLayers")
	ret = (uint)((value).Int())
	return ret
}

// class: VRDisplayEvent
type DisplayEvent struct {
	domcore.Event
}

// DisplayEventFromJS is casting a js.Wrapper into DisplayEvent.
func DisplayEventFromJS(value js.Wrapper) *DisplayEvent {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &DisplayEvent{}
	ret.Value_JS = input
	return ret
}

func NewVRDisplayEvent(_type string, eventInitDict *DisplayEventInit) (_result *DisplayEvent) {
	_klass := js.Global().Get("VRDisplayEvent")
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
		_converted *DisplayEvent // javascript: VRDisplayEvent _what_return_name
	)
	_converted = DisplayEventFromJS(_returned)
	_result = _converted
	return
}

// Display returning attribute 'display' with
// type Display (idl: VRDisplay).
func (_this *DisplayEvent) Display() *Display {
	var ret *Display
	value := _this.Value_JS.Get("display")
	ret = DisplayFromJS(value)
	return ret
}

// Reason returning attribute 'reason' with
// type DisplayEventReason (idl: VRDisplayEventReason).
func (_this *DisplayEvent) Reason() *DisplayEventReason {
	var ret *DisplayEventReason
	value := _this.Value_JS.Get("reason")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := DisplayEventReasonFromJS(value)
		ret = &__tmp
	}
	return ret
}

// class: VREyeParameters
type EyeParameters struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *EyeParameters) JSValue() js.Value {
	return _this.Value_JS
}

// EyeParametersFromJS is casting a js.Wrapper into EyeParameters.
func EyeParametersFromJS(value js.Wrapper) *EyeParameters {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &EyeParameters{}
	ret.Value_JS = input
	return ret
}

// Offset returning attribute 'offset' with
// type javascript.Float32Array (idl: Float32Array).
func (_this *EyeParameters) Offset() *javascript.Float32Array {
	var ret *javascript.Float32Array
	value := _this.Value_JS.Get("offset")
	ret = javascript.Float32ArrayFromJS(value)
	return ret
}

// FieldOfView returning attribute 'fieldOfView' with
// type FieldOfView (idl: VRFieldOfView).
func (_this *EyeParameters) FieldOfView() *FieldOfView {
	var ret *FieldOfView
	value := _this.Value_JS.Get("fieldOfView")
	ret = FieldOfViewFromJS(value)
	return ret
}

// RenderWidth returning attribute 'renderWidth' with
// type uint (idl: unsigned long).
func (_this *EyeParameters) RenderWidth() uint {
	var ret uint
	value := _this.Value_JS.Get("renderWidth")
	ret = (uint)((value).Int())
	return ret
}

// RenderHeight returning attribute 'renderHeight' with
// type uint (idl: unsigned long).
func (_this *EyeParameters) RenderHeight() uint {
	var ret uint
	value := _this.Value_JS.Get("renderHeight")
	ret = (uint)((value).Int())
	return ret
}

// class: VRFieldOfView
type FieldOfView struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FieldOfView) JSValue() js.Value {
	return _this.Value_JS
}

// FieldOfViewFromJS is casting a js.Wrapper into FieldOfView.
func FieldOfViewFromJS(value js.Wrapper) *FieldOfView {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &FieldOfView{}
	ret.Value_JS = input
	return ret
}

// UpDegrees returning attribute 'upDegrees' with
// type float64 (idl: double).
func (_this *FieldOfView) UpDegrees() float64 {
	var ret float64
	value := _this.Value_JS.Get("upDegrees")
	ret = (value).Float()
	return ret
}

// RightDegrees returning attribute 'rightDegrees' with
// type float64 (idl: double).
func (_this *FieldOfView) RightDegrees() float64 {
	var ret float64
	value := _this.Value_JS.Get("rightDegrees")
	ret = (value).Float()
	return ret
}

// DownDegrees returning attribute 'downDegrees' with
// type float64 (idl: double).
func (_this *FieldOfView) DownDegrees() float64 {
	var ret float64
	value := _this.Value_JS.Get("downDegrees")
	ret = (value).Float()
	return ret
}

// LeftDegrees returning attribute 'leftDegrees' with
// type float64 (idl: double).
func (_this *FieldOfView) LeftDegrees() float64 {
	var ret float64
	value := _this.Value_JS.Get("leftDegrees")
	ret = (value).Float()
	return ret
}

// class: VRFrameData
type FrameData struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FrameData) JSValue() js.Value {
	return _this.Value_JS
}

// FrameDataFromJS is casting a js.Wrapper into FrameData.
func FrameDataFromJS(value js.Wrapper) *FrameData {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &FrameData{}
	ret.Value_JS = input
	return ret
}

func NewVRFrameData() (_result *FrameData) {
	_klass := js.Global().Get("VRFrameData")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *FrameData // javascript: VRFrameData _what_return_name
	)
	_converted = FrameDataFromJS(_returned)
	_result = _converted
	return
}

// Timestamp returning attribute 'timestamp' with
// type float64 (idl: double).
func (_this *FrameData) Timestamp() float64 {
	var ret float64
	value := _this.Value_JS.Get("timestamp")
	ret = (value).Float()
	return ret
}

// LeftProjectionMatrix returning attribute 'leftProjectionMatrix' with
// type javascript.Float32Array (idl: Float32Array).
func (_this *FrameData) LeftProjectionMatrix() *javascript.Float32Array {
	var ret *javascript.Float32Array
	value := _this.Value_JS.Get("leftProjectionMatrix")
	ret = javascript.Float32ArrayFromJS(value)
	return ret
}

// LeftViewMatrix returning attribute 'leftViewMatrix' with
// type javascript.Float32Array (idl: Float32Array).
func (_this *FrameData) LeftViewMatrix() *javascript.Float32Array {
	var ret *javascript.Float32Array
	value := _this.Value_JS.Get("leftViewMatrix")
	ret = javascript.Float32ArrayFromJS(value)
	return ret
}

// RightProjectionMatrix returning attribute 'rightProjectionMatrix' with
// type javascript.Float32Array (idl: Float32Array).
func (_this *FrameData) RightProjectionMatrix() *javascript.Float32Array {
	var ret *javascript.Float32Array
	value := _this.Value_JS.Get("rightProjectionMatrix")
	ret = javascript.Float32ArrayFromJS(value)
	return ret
}

// RightViewMatrix returning attribute 'rightViewMatrix' with
// type javascript.Float32Array (idl: Float32Array).
func (_this *FrameData) RightViewMatrix() *javascript.Float32Array {
	var ret *javascript.Float32Array
	value := _this.Value_JS.Get("rightViewMatrix")
	ret = javascript.Float32ArrayFromJS(value)
	return ret
}

// Pose returning attribute 'pose' with
// type Pose (idl: VRPose).
func (_this *FrameData) Pose() *Pose {
	var ret *Pose
	value := _this.Value_JS.Get("pose")
	ret = PoseFromJS(value)
	return ret
}

// class: VRPose
type Pose struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Pose) JSValue() js.Value {
	return _this.Value_JS
}

// PoseFromJS is casting a js.Wrapper into Pose.
func PoseFromJS(value js.Wrapper) *Pose {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Pose{}
	ret.Value_JS = input
	return ret
}

// Position returning attribute 'position' with
// type javascript.Float32Array (idl: Float32Array).
func (_this *Pose) Position() *javascript.Float32Array {
	var ret *javascript.Float32Array
	value := _this.Value_JS.Get("position")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = javascript.Float32ArrayFromJS(value)
	}
	return ret
}

// LinearVelocity returning attribute 'linearVelocity' with
// type javascript.Float32Array (idl: Float32Array).
func (_this *Pose) LinearVelocity() *javascript.Float32Array {
	var ret *javascript.Float32Array
	value := _this.Value_JS.Get("linearVelocity")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = javascript.Float32ArrayFromJS(value)
	}
	return ret
}

// LinearAcceleration returning attribute 'linearAcceleration' with
// type javascript.Float32Array (idl: Float32Array).
func (_this *Pose) LinearAcceleration() *javascript.Float32Array {
	var ret *javascript.Float32Array
	value := _this.Value_JS.Get("linearAcceleration")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = javascript.Float32ArrayFromJS(value)
	}
	return ret
}

// Orientation returning attribute 'orientation' with
// type javascript.Float32Array (idl: Float32Array).
func (_this *Pose) Orientation() *javascript.Float32Array {
	var ret *javascript.Float32Array
	value := _this.Value_JS.Get("orientation")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = javascript.Float32ArrayFromJS(value)
	}
	return ret
}

// AngularVelocity returning attribute 'angularVelocity' with
// type javascript.Float32Array (idl: Float32Array).
func (_this *Pose) AngularVelocity() *javascript.Float32Array {
	var ret *javascript.Float32Array
	value := _this.Value_JS.Get("angularVelocity")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = javascript.Float32ArrayFromJS(value)
	}
	return ret
}

// AngularAcceleration returning attribute 'angularAcceleration' with
// type javascript.Float32Array (idl: Float32Array).
func (_this *Pose) AngularAcceleration() *javascript.Float32Array {
	var ret *javascript.Float32Array
	value := _this.Value_JS.Get("angularAcceleration")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = javascript.Float32ArrayFromJS(value)
	}
	return ret
}

// class: Promise
type PromiseSequenceDisplay struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseSequenceDisplay) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseSequenceDisplayFromJS is casting a js.Wrapper into PromiseSequenceDisplay.
func PromiseSequenceDisplayFromJS(value js.Wrapper) *PromiseSequenceDisplay {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromiseSequenceDisplay{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseSequenceDisplay) Then(onFulfilled *PromiseSequenceDisplayOnFulfilled, onRejected *PromiseSequenceDisplayOnRejected) (_result *PromiseSequenceDisplay) {
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
		_converted *PromiseSequenceDisplay // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceDisplayFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseSequenceDisplay) Catch(onRejected *PromiseSequenceDisplayOnRejected) (_result *PromiseSequenceDisplay) {
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
		_converted *PromiseSequenceDisplay // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceDisplayFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseSequenceDisplay) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseSequenceDisplay) {
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
		_converted *PromiseSequenceDisplay // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceDisplayFromJS(_returned)
	_result = _converted
	return
}

// class: VRStageParameters
type StageParameters struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *StageParameters) JSValue() js.Value {
	return _this.Value_JS
}

// StageParametersFromJS is casting a js.Wrapper into StageParameters.
func StageParametersFromJS(value js.Wrapper) *StageParameters {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &StageParameters{}
	ret.Value_JS = input
	return ret
}

// SittingToStandingTransform returning attribute 'sittingToStandingTransform' with
// type javascript.Float32Array (idl: Float32Array).
func (_this *StageParameters) SittingToStandingTransform() *javascript.Float32Array {
	var ret *javascript.Float32Array
	value := _this.Value_JS.Get("sittingToStandingTransform")
	ret = javascript.Float32ArrayFromJS(value)
	return ret
}

// SizeX returning attribute 'sizeX' with
// type float32 (idl: float).
func (_this *StageParameters) SizeX() float32 {
	var ret float32
	value := _this.Value_JS.Get("sizeX")
	ret = (float32)((value).Float())
	return ret
}

// SizeZ returning attribute 'sizeZ' with
// type float32 (idl: float).
func (_this *StageParameters) SizeZ() float32 {
	var ret float32
	value := _this.Value_JS.Get("sizeZ")
	ret = (float32)((value).Float())
	return ret
}
