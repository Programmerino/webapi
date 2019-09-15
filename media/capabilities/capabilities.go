// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package capabilities

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/javascript"
	"github.com/gowebapi/webapi/media/encrypted"
)

// using following types:
// encrypted.MediaKeySystemAccess
// encrypted.MediaKeysRequirement
// javascript.PromiseFinally

// source idl files:
// media-capabilities.idl
// promises.idl

// transform files:
// media-capabilities.go.md
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

// enum: MediaDecodingType
type MediaDecodingType int

const (
	FileMediaDecodingType MediaDecodingType = iota
	MediaSourceMediaDecodingType
)

var mediaDecodingTypeToWasmTable = []string{
	"file", "media-source",
}

var mediaDecodingTypeFromWasmTable = map[string]MediaDecodingType{
	"file": FileMediaDecodingType, "media-source": MediaSourceMediaDecodingType,
}

// JSValue is converting this enum into a javascript object
func (this *MediaDecodingType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this MediaDecodingType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(mediaDecodingTypeToWasmTable) {
		return mediaDecodingTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// MediaDecodingTypeFromJS is converting a javascript value into
// a MediaDecodingType enum value.
func MediaDecodingTypeFromJS(value js.Value) MediaDecodingType {
	key := value.String()
	conv, ok := mediaDecodingTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: MediaEncodingType
type MediaEncodingType int

const (
	RecordMediaEncodingType MediaEncodingType = iota
	TransmissionMediaEncodingType
)

var mediaEncodingTypeToWasmTable = []string{
	"record", "transmission",
}

var mediaEncodingTypeFromWasmTable = map[string]MediaEncodingType{
	"record": RecordMediaEncodingType, "transmission": TransmissionMediaEncodingType,
}

// JSValue is converting this enum into a javascript object
func (this *MediaEncodingType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this MediaEncodingType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(mediaEncodingTypeToWasmTable) {
		return mediaEncodingTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// MediaEncodingTypeFromJS is converting a javascript value into
// a MediaEncodingType enum value.
func MediaEncodingTypeFromJS(value js.Value) MediaEncodingType {
	key := value.String()
	conv, ok := mediaEncodingTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: ScreenColorGamut
type ScreenColorGamut int

const (
	SrgbScreenColorGamut ScreenColorGamut = iota
	P3ScreenColorGamut
	Rec2020ScreenColorGamut
)

var screenColorGamutToWasmTable = []string{
	"srgb", "p3", "rec2020",
}

var screenColorGamutFromWasmTable = map[string]ScreenColorGamut{
	"srgb": SrgbScreenColorGamut, "p3": P3ScreenColorGamut, "rec2020": Rec2020ScreenColorGamut,
}

// JSValue is converting this enum into a javascript object
func (this *ScreenColorGamut) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ScreenColorGamut) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(screenColorGamutToWasmTable) {
		return screenColorGamutToWasmTable[idx]
	}
	panic("unknown input value")
}

// ScreenColorGamutFromJS is converting a javascript value into
// a ScreenColorGamut enum value.
func ScreenColorGamutFromJS(value js.Value) ScreenColorGamut {
	key := value.String()
	conv, ok := screenColorGamutFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: PromiseTemplateOnFulfilled
type PromiseMediaCapabilitiesDecodingInfoOnFulfilledFunc func(value *MediaCapabilitiesDecodingInfo)

// PromiseMediaCapabilitiesDecodingInfoOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseMediaCapabilitiesDecodingInfoOnFulfilled js.Func

func PromiseMediaCapabilitiesDecodingInfoOnFulfilledToJS(callback PromiseMediaCapabilitiesDecodingInfoOnFulfilledFunc) *PromiseMediaCapabilitiesDecodingInfoOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseMediaCapabilitiesDecodingInfoOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *MediaCapabilitiesDecodingInfo // javascript: MediaCapabilitiesDecodingInfo value
		)
		_p0 = MediaCapabilitiesDecodingInfoFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseMediaCapabilitiesDecodingInfoOnFulfilledFromJS(_value js.Value) PromiseMediaCapabilitiesDecodingInfoOnFulfilledFunc {
	return func(value *MediaCapabilitiesDecodingInfo) {
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
type PromiseMediaCapabilitiesDecodingInfoOnRejectedFunc func(reason js.Value)

// PromiseMediaCapabilitiesDecodingInfoOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseMediaCapabilitiesDecodingInfoOnRejected js.Func

func PromiseMediaCapabilitiesDecodingInfoOnRejectedToJS(callback PromiseMediaCapabilitiesDecodingInfoOnRejectedFunc) *PromiseMediaCapabilitiesDecodingInfoOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseMediaCapabilitiesDecodingInfoOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromiseMediaCapabilitiesDecodingInfoOnRejectedFromJS(_value js.Value) PromiseMediaCapabilitiesDecodingInfoOnRejectedFunc {
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
type PromiseMediaCapabilitiesInfoOnFulfilledFunc func(value *MediaCapabilitiesInfo)

// PromiseMediaCapabilitiesInfoOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseMediaCapabilitiesInfoOnFulfilled js.Func

func PromiseMediaCapabilitiesInfoOnFulfilledToJS(callback PromiseMediaCapabilitiesInfoOnFulfilledFunc) *PromiseMediaCapabilitiesInfoOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseMediaCapabilitiesInfoOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *MediaCapabilitiesInfo // javascript: MediaCapabilitiesInfo value
		)
		_p0 = MediaCapabilitiesInfoFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseMediaCapabilitiesInfoOnFulfilledFromJS(_value js.Value) PromiseMediaCapabilitiesInfoOnFulfilledFunc {
	return func(value *MediaCapabilitiesInfo) {
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
type PromiseMediaCapabilitiesInfoOnRejectedFunc func(reason js.Value)

// PromiseMediaCapabilitiesInfoOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseMediaCapabilitiesInfoOnRejected js.Func

func PromiseMediaCapabilitiesInfoOnRejectedToJS(callback PromiseMediaCapabilitiesInfoOnRejectedFunc) *PromiseMediaCapabilitiesInfoOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseMediaCapabilitiesInfoOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromiseMediaCapabilitiesInfoOnRejectedFromJS(_value js.Value) PromiseMediaCapabilitiesInfoOnRejectedFunc {
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

// dictionary: AudioConfiguration
type AudioConfiguration struct {
	ContentType string
	Channels    string
	Bitrate     int
	Samplerate  uint
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *AudioConfiguration) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.ContentType
	out.Set("contentType", value0)
	value1 := _this.Channels
	out.Set("channels", value1)
	value2 := _this.Bitrate
	out.Set("bitrate", value2)
	value3 := _this.Samplerate
	out.Set("samplerate", value3)
	return out
}

// AudioConfigurationFromJS is allocating a new
// AudioConfiguration object and copy all values from
// input javascript object
func AudioConfigurationFromJS(value js.Wrapper) *AudioConfiguration {
	input := value.JSValue()
	var out AudioConfiguration
	var (
		value0 string // javascript: DOMString {contentType ContentType contentType}
		value1 string // javascript: DOMString {channels Channels channels}
		value2 int    // javascript: unsigned long long {bitrate Bitrate bitrate}
		value3 uint   // javascript: unsigned long {samplerate Samplerate samplerate}
	)
	value0 = (input.Get("contentType")).String()
	out.ContentType = value0
	value1 = (input.Get("channels")).String()
	out.Channels = value1
	value2 = (input.Get("bitrate")).Int()
	out.Bitrate = value2
	value3 = (uint)((input.Get("samplerate")).Int())
	out.Samplerate = value3
	return &out
}

// dictionary: MediaCapabilitiesDecodingInfo
type MediaCapabilitiesDecodingInfo struct {
	Supported       bool
	Smooth          bool
	PowerEfficient  bool
	KeySystemAccess *encrypted.MediaKeySystemAccess
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *MediaCapabilitiesDecodingInfo) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Supported
	out.Set("supported", value0)
	value1 := _this.Smooth
	out.Set("smooth", value1)
	value2 := _this.PowerEfficient
	out.Set("powerEfficient", value2)
	value3 := _this.KeySystemAccess.JSValue()
	out.Set("keySystemAccess", value3)
	return out
}

// MediaCapabilitiesDecodingInfoFromJS is allocating a new
// MediaCapabilitiesDecodingInfo object and copy all values from
// input javascript object
func MediaCapabilitiesDecodingInfoFromJS(value js.Wrapper) *MediaCapabilitiesDecodingInfo {
	input := value.JSValue()
	var out MediaCapabilitiesDecodingInfo
	var (
		value0 bool                            // javascript: boolean {supported Supported supported}
		value1 bool                            // javascript: boolean {smooth Smooth smooth}
		value2 bool                            // javascript: boolean {powerEfficient PowerEfficient powerEfficient}
		value3 *encrypted.MediaKeySystemAccess // javascript: MediaKeySystemAccess {keySystemAccess KeySystemAccess keySystemAccess}
	)
	value0 = (input.Get("supported")).Bool()
	out.Supported = value0
	value1 = (input.Get("smooth")).Bool()
	out.Smooth = value1
	value2 = (input.Get("powerEfficient")).Bool()
	out.PowerEfficient = value2
	value3 = encrypted.MediaKeySystemAccessFromJS(input.Get("keySystemAccess"))
	out.KeySystemAccess = value3
	return &out
}

// dictionary: MediaCapabilitiesInfo
type MediaCapabilitiesInfo struct {
	Supported      bool
	Smooth         bool
	PowerEfficient bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *MediaCapabilitiesInfo) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Supported
	out.Set("supported", value0)
	value1 := _this.Smooth
	out.Set("smooth", value1)
	value2 := _this.PowerEfficient
	out.Set("powerEfficient", value2)
	return out
}

// MediaCapabilitiesInfoFromJS is allocating a new
// MediaCapabilitiesInfo object and copy all values from
// input javascript object
func MediaCapabilitiesInfoFromJS(value js.Wrapper) *MediaCapabilitiesInfo {
	input := value.JSValue()
	var out MediaCapabilitiesInfo
	var (
		value0 bool // javascript: boolean {supported Supported supported}
		value1 bool // javascript: boolean {smooth Smooth smooth}
		value2 bool // javascript: boolean {powerEfficient PowerEfficient powerEfficient}
	)
	value0 = (input.Get("supported")).Bool()
	out.Supported = value0
	value1 = (input.Get("smooth")).Bool()
	out.Smooth = value1
	value2 = (input.Get("powerEfficient")).Bool()
	out.PowerEfficient = value2
	return &out
}

// dictionary: MediaCapabilitiesKeySystemConfiguration
type MediaCapabilitiesKeySystemConfiguration struct {
	KeySystem             string
	InitDataType          string
	AudioRobustness       string
	VideoRobustness       string
	DistinctiveIdentifier encrypted.MediaKeysRequirement
	PersistentState       encrypted.MediaKeysRequirement
	SessionTypes          []string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *MediaCapabilitiesKeySystemConfiguration) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.KeySystem
	out.Set("keySystem", value0)
	value1 := _this.InitDataType
	out.Set("initDataType", value1)
	value2 := _this.AudioRobustness
	out.Set("audioRobustness", value2)
	value3 := _this.VideoRobustness
	out.Set("videoRobustness", value3)
	value4 := _this.DistinctiveIdentifier.JSValue()
	out.Set("distinctiveIdentifier", value4)
	value5 := _this.PersistentState.JSValue()
	out.Set("persistentState", value5)
	value6 := js.Global().Get("Array").New(len(_this.SessionTypes))
	for __idx6, __seq_in6 := range _this.SessionTypes {
		__seq_out6 := __seq_in6
		value6.SetIndex(__idx6, __seq_out6)
	}
	out.Set("sessionTypes", value6)
	return out
}

// MediaCapabilitiesKeySystemConfigurationFromJS is allocating a new
// MediaCapabilitiesKeySystemConfiguration object and copy all values from
// input javascript object
func MediaCapabilitiesKeySystemConfigurationFromJS(value js.Wrapper) *MediaCapabilitiesKeySystemConfiguration {
	input := value.JSValue()
	var out MediaCapabilitiesKeySystemConfiguration
	var (
		value0 string                         // javascript: DOMString {keySystem KeySystem keySystem}
		value1 string                         // javascript: DOMString {initDataType InitDataType initDataType}
		value2 string                         // javascript: DOMString {audioRobustness AudioRobustness audioRobustness}
		value3 string                         // javascript: DOMString {videoRobustness VideoRobustness videoRobustness}
		value4 encrypted.MediaKeysRequirement // javascript: MediaKeysRequirement {distinctiveIdentifier DistinctiveIdentifier distinctiveIdentifier}
		value5 encrypted.MediaKeysRequirement // javascript: MediaKeysRequirement {persistentState PersistentState persistentState}
		value6 []string                       // javascript: sequence<DOMString> {sessionTypes SessionTypes sessionTypes}
	)
	value0 = (input.Get("keySystem")).String()
	out.KeySystem = value0
	value1 = (input.Get("initDataType")).String()
	out.InitDataType = value1
	value2 = (input.Get("audioRobustness")).String()
	out.AudioRobustness = value2
	value3 = (input.Get("videoRobustness")).String()
	out.VideoRobustness = value3
	value4 = encrypted.MediaKeysRequirementFromJS(input.Get("distinctiveIdentifier"))
	out.DistinctiveIdentifier = value4
	value5 = encrypted.MediaKeysRequirementFromJS(input.Get("persistentState"))
	out.PersistentState = value5
	__length6 := input.Get("sessionTypes").Length()
	__array6 := make([]string, __length6, __length6)
	for __idx6 := 0; __idx6 < __length6; __idx6++ {
		var __seq_out6 string
		__seq_in6 := input.Get("sessionTypes").Index(__idx6)
		__seq_out6 = (__seq_in6).String()
		__array6[__idx6] = __seq_out6
	}
	value6 = __array6
	out.SessionTypes = value6
	return &out
}

// dictionary: MediaConfiguration
type MediaConfiguration struct {
	Video *VideoConfiguration
	Audio *AudioConfiguration
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *MediaConfiguration) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Video.JSValue()
	out.Set("video", value0)
	value1 := _this.Audio.JSValue()
	out.Set("audio", value1)
	return out
}

// MediaConfigurationFromJS is allocating a new
// MediaConfiguration object and copy all values from
// input javascript object
func MediaConfigurationFromJS(value js.Wrapper) *MediaConfiguration {
	input := value.JSValue()
	var out MediaConfiguration
	var (
		value0 *VideoConfiguration // javascript: VideoConfiguration {video Video video}
		value1 *AudioConfiguration // javascript: AudioConfiguration {audio Audio audio}
	)
	value0 = VideoConfigurationFromJS(input.Get("video"))
	out.Video = value0
	value1 = AudioConfigurationFromJS(input.Get("audio"))
	out.Audio = value1
	return &out
}

// dictionary: MediaDecodingConfiguration
type MediaDecodingConfiguration struct {
	Video                  *VideoConfiguration
	Audio                  *AudioConfiguration
	Type                   MediaDecodingType
	KeySystemConfiguration *MediaCapabilitiesKeySystemConfiguration
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *MediaDecodingConfiguration) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Video.JSValue()
	out.Set("video", value0)
	value1 := _this.Audio.JSValue()
	out.Set("audio", value1)
	value2 := _this.Type.JSValue()
	out.Set("type", value2)
	value3 := _this.KeySystemConfiguration.JSValue()
	out.Set("keySystemConfiguration", value3)
	return out
}

// MediaDecodingConfigurationFromJS is allocating a new
// MediaDecodingConfiguration object and copy all values from
// input javascript object
func MediaDecodingConfigurationFromJS(value js.Wrapper) *MediaDecodingConfiguration {
	input := value.JSValue()
	var out MediaDecodingConfiguration
	var (
		value0 *VideoConfiguration                      // javascript: VideoConfiguration {video Video video}
		value1 *AudioConfiguration                      // javascript: AudioConfiguration {audio Audio audio}
		value2 MediaDecodingType                        // javascript: MediaDecodingType {type Type _type}
		value3 *MediaCapabilitiesKeySystemConfiguration // javascript: MediaCapabilitiesKeySystemConfiguration {keySystemConfiguration KeySystemConfiguration keySystemConfiguration}
	)
	value0 = VideoConfigurationFromJS(input.Get("video"))
	out.Video = value0
	value1 = AudioConfigurationFromJS(input.Get("audio"))
	out.Audio = value1
	value2 = MediaDecodingTypeFromJS(input.Get("type"))
	out.Type = value2
	value3 = MediaCapabilitiesKeySystemConfigurationFromJS(input.Get("keySystemConfiguration"))
	out.KeySystemConfiguration = value3
	return &out
}

// dictionary: MediaEncodingConfiguration
type MediaEncodingConfiguration struct {
	Video *VideoConfiguration
	Audio *AudioConfiguration
	Type  MediaEncodingType
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *MediaEncodingConfiguration) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Video.JSValue()
	out.Set("video", value0)
	value1 := _this.Audio.JSValue()
	out.Set("audio", value1)
	value2 := _this.Type.JSValue()
	out.Set("type", value2)
	return out
}

// MediaEncodingConfigurationFromJS is allocating a new
// MediaEncodingConfiguration object and copy all values from
// input javascript object
func MediaEncodingConfigurationFromJS(value js.Wrapper) *MediaEncodingConfiguration {
	input := value.JSValue()
	var out MediaEncodingConfiguration
	var (
		value0 *VideoConfiguration // javascript: VideoConfiguration {video Video video}
		value1 *AudioConfiguration // javascript: AudioConfiguration {audio Audio audio}
		value2 MediaEncodingType   // javascript: MediaEncodingType {type Type _type}
	)
	value0 = VideoConfigurationFromJS(input.Get("video"))
	out.Video = value0
	value1 = AudioConfigurationFromJS(input.Get("audio"))
	out.Audio = value1
	value2 = MediaEncodingTypeFromJS(input.Get("type"))
	out.Type = value2
	return &out
}

// dictionary: VideoConfiguration
type VideoConfiguration struct {
	ContentType string
	Width       uint
	Height      uint
	Bitrate     int
	Framerate   string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *VideoConfiguration) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.ContentType
	out.Set("contentType", value0)
	value1 := _this.Width
	out.Set("width", value1)
	value2 := _this.Height
	out.Set("height", value2)
	value3 := _this.Bitrate
	out.Set("bitrate", value3)
	value4 := _this.Framerate
	out.Set("framerate", value4)
	return out
}

// VideoConfigurationFromJS is allocating a new
// VideoConfiguration object and copy all values from
// input javascript object
func VideoConfigurationFromJS(value js.Wrapper) *VideoConfiguration {
	input := value.JSValue()
	var out VideoConfiguration
	var (
		value0 string // javascript: DOMString {contentType ContentType contentType}
		value1 uint   // javascript: unsigned long {width Width width}
		value2 uint   // javascript: unsigned long {height Height height}
		value3 int    // javascript: unsigned long long {bitrate Bitrate bitrate}
		value4 string // javascript: DOMString {framerate Framerate framerate}
	)
	value0 = (input.Get("contentType")).String()
	out.ContentType = value0
	value1 = (uint)((input.Get("width")).Int())
	out.Width = value1
	value2 = (uint)((input.Get("height")).Int())
	out.Height = value2
	value3 = (input.Get("bitrate")).Int()
	out.Bitrate = value3
	value4 = (input.Get("framerate")).String()
	out.Framerate = value4
	return &out
}

// class: MediaCapabilities
type MediaCapabilities struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *MediaCapabilities) JSValue() js.Value {
	return _this.Value_JS
}

// MediaCapabilitiesFromJS is casting a js.Wrapper into MediaCapabilities.
func MediaCapabilitiesFromJS(value js.Wrapper) *MediaCapabilities {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &MediaCapabilities{}
	ret.Value_JS = input
	return ret
}

func (_this *MediaCapabilities) DecodingInfo(configuration *MediaDecodingConfiguration) (_result *PromiseMediaCapabilitiesDecodingInfo) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := configuration.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("decodingInfo", _args[0:_end]...)
	var (
		_converted *PromiseMediaCapabilitiesDecodingInfo // javascript: Promise _what_return_name
	)
	_converted = PromiseMediaCapabilitiesDecodingInfoFromJS(_returned)
	_result = _converted
	return
}

func (_this *MediaCapabilities) EncodingInfo(configuration *MediaEncodingConfiguration) (_result *PromiseMediaCapabilitiesInfo) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := configuration.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("encodingInfo", _args[0:_end]...)
	var (
		_converted *PromiseMediaCapabilitiesInfo // javascript: Promise _what_return_name
	)
	_converted = PromiseMediaCapabilitiesInfoFromJS(_returned)
	_result = _converted
	return
}

// class: Promise
type PromiseMediaCapabilitiesDecodingInfo struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseMediaCapabilitiesDecodingInfo) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseMediaCapabilitiesDecodingInfoFromJS is casting a js.Wrapper into PromiseMediaCapabilitiesDecodingInfo.
func PromiseMediaCapabilitiesDecodingInfoFromJS(value js.Wrapper) *PromiseMediaCapabilitiesDecodingInfo {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromiseMediaCapabilitiesDecodingInfo{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseMediaCapabilitiesDecodingInfo) Then(onFulfilled *PromiseMediaCapabilitiesDecodingInfoOnFulfilled, onRejected *PromiseMediaCapabilitiesDecodingInfoOnRejected) (_result *PromiseMediaCapabilitiesDecodingInfo) {
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
		_converted *PromiseMediaCapabilitiesDecodingInfo // javascript: Promise _what_return_name
	)
	_converted = PromiseMediaCapabilitiesDecodingInfoFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseMediaCapabilitiesDecodingInfo) Catch(onRejected *PromiseMediaCapabilitiesDecodingInfoOnRejected) (_result *PromiseMediaCapabilitiesDecodingInfo) {
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
		_converted *PromiseMediaCapabilitiesDecodingInfo // javascript: Promise _what_return_name
	)
	_converted = PromiseMediaCapabilitiesDecodingInfoFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseMediaCapabilitiesDecodingInfo) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseMediaCapabilitiesDecodingInfo) {
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
		_converted *PromiseMediaCapabilitiesDecodingInfo // javascript: Promise _what_return_name
	)
	_converted = PromiseMediaCapabilitiesDecodingInfoFromJS(_returned)
	_result = _converted
	return
}

// class: Promise
type PromiseMediaCapabilitiesInfo struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseMediaCapabilitiesInfo) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseMediaCapabilitiesInfoFromJS is casting a js.Wrapper into PromiseMediaCapabilitiesInfo.
func PromiseMediaCapabilitiesInfoFromJS(value js.Wrapper) *PromiseMediaCapabilitiesInfo {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromiseMediaCapabilitiesInfo{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseMediaCapabilitiesInfo) Then(onFulfilled *PromiseMediaCapabilitiesInfoOnFulfilled, onRejected *PromiseMediaCapabilitiesInfoOnRejected) (_result *PromiseMediaCapabilitiesInfo) {
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
		_converted *PromiseMediaCapabilitiesInfo // javascript: Promise _what_return_name
	)
	_converted = PromiseMediaCapabilitiesInfoFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseMediaCapabilitiesInfo) Catch(onRejected *PromiseMediaCapabilitiesInfoOnRejected) (_result *PromiseMediaCapabilitiesInfo) {
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
		_converted *PromiseMediaCapabilitiesInfo // javascript: Promise _what_return_name
	)
	_converted = PromiseMediaCapabilitiesInfoFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseMediaCapabilitiesInfo) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseMediaCapabilitiesInfo) {
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
		_converted *PromiseMediaCapabilitiesInfo // javascript: Promise _what_return_name
	)
	_converted = PromiseMediaCapabilitiesInfoFromJS(_returned)
	_result = _converted
	return
}

// class: ScreenLuminance
type ScreenLuminance struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *ScreenLuminance) JSValue() js.Value {
	return _this.Value_JS
}

// ScreenLuminanceFromJS is casting a js.Wrapper into ScreenLuminance.
func ScreenLuminanceFromJS(value js.Wrapper) *ScreenLuminance {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &ScreenLuminance{}
	ret.Value_JS = input
	return ret
}

// Min returning attribute 'min' with
// type float64 (idl: double).
func (_this *ScreenLuminance) Min() float64 {
	var ret float64
	value := _this.Value_JS.Get("min")
	ret = (value).Float()
	return ret
}

// Max returning attribute 'max' with
// type float64 (idl: double).
func (_this *ScreenLuminance) Max() float64 {
	var ret float64
	value := _this.Value_JS.Get("max")
	ret = (value).Float()
	return ret
}

// MaxAverage returning attribute 'maxAverage' with
// type float64 (idl: double).
func (_this *ScreenLuminance) MaxAverage() float64 {
	var ret float64
	value := _this.Value_JS.Get("maxAverage")
	ret = (value).Float()
	return ret
}
