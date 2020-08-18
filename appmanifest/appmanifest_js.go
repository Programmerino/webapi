// Code generated by webidl-bind. DO NOT EDIT.

package appmanifest

import "syscall/js"

import (
	"github.com/Programmerino/webapi/appmanifest/appmenifestres"
	"github.com/Programmerino/webapi/dom/domcore"
	"github.com/Programmerino/webapi/html/htmlcommon"
	"github.com/Programmerino/webapi/javascript"
	"github.com/Programmerino/webapi/media/orientation"
	"github.com/Programmerino/webapi/serviceworker"
)

// using following types:
// appmenifestres.ExternalApplicationResource
// appmenifestres.ImageResource
// domcore.Event
// domcore.EventInit
// htmlcommon.WorkerType
// javascript.PromiseFinally
// orientation.OrientationLockType
// serviceworker.ServiceWorkerUpdateViaCache

// source idl files:
// appmanifest.idl
// promises.idl

// transform files:
// appmanifest.go.md
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

// enum: AppBannerPromptOutcome
type AppBannerPromptOutcome int

const (
	AcceptedAppBannerPromptOutcome AppBannerPromptOutcome = iota
	DismissedAppBannerPromptOutcome
)

var appBannerPromptOutcomeToWasmTable = []string{
	"accepted", "dismissed",
}

var appBannerPromptOutcomeFromWasmTable = map[string]AppBannerPromptOutcome{
	"accepted": AcceptedAppBannerPromptOutcome, "dismissed": DismissedAppBannerPromptOutcome,
}

// JSValue is converting this enum into a javascript object
func (this *AppBannerPromptOutcome) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this AppBannerPromptOutcome) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(appBannerPromptOutcomeToWasmTable) {
		return appBannerPromptOutcomeToWasmTable[idx]
	}
	panic("unknown input value")
}

// AppBannerPromptOutcomeFromJS is converting a javascript value into
// a AppBannerPromptOutcome enum value.
func AppBannerPromptOutcomeFromJS(value js.Value) AppBannerPromptOutcome {
	key := value.String()
	conv, ok := appBannerPromptOutcomeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: DisplayModeType
type DisplayModeType int

const (
	FullscreenDisplayModeType DisplayModeType = iota
	StandaloneDisplayModeType
	MinimalUiDisplayModeType
	BrowserDisplayModeType
)

var displayModeTypeToWasmTable = []string{
	"fullscreen", "standalone", "minimal-ui", "browser",
}

var displayModeTypeFromWasmTable = map[string]DisplayModeType{
	"fullscreen": FullscreenDisplayModeType, "standalone": StandaloneDisplayModeType, "minimal-ui": MinimalUiDisplayModeType, "browser": BrowserDisplayModeType,
}

// JSValue is converting this enum into a javascript object
func (this *DisplayModeType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this DisplayModeType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(displayModeTypeToWasmTable) {
		return displayModeTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// DisplayModeTypeFromJS is converting a javascript value into
// a DisplayModeType enum value.
func DisplayModeTypeFromJS(value js.Value) DisplayModeType {
	key := value.String()
	conv, ok := displayModeTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: TextDirectionType
type TextDirectionType int

const (
	LtrTextDirectionType TextDirectionType = iota
	RtlTextDirectionType
	AutoTextDirectionType
)

var textDirectionTypeToWasmTable = []string{
	"ltr", "rtl", "auto",
}

var textDirectionTypeFromWasmTable = map[string]TextDirectionType{
	"ltr": LtrTextDirectionType, "rtl": RtlTextDirectionType, "auto": AutoTextDirectionType,
}

// JSValue is converting this enum into a javascript object
func (this *TextDirectionType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this TextDirectionType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(textDirectionTypeToWasmTable) {
		return textDirectionTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// TextDirectionTypeFromJS is converting a javascript value into
// a TextDirectionType enum value.
func TextDirectionTypeFromJS(value js.Value) TextDirectionType {
	key := value.String()
	conv, ok := textDirectionTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: PromiseTemplateOnFulfilled
type PromisePromptResponseObjectOnFulfilledFunc func(value *PromptResponseObject)

// PromisePromptResponseObjectOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromisePromptResponseObjectOnFulfilled js.Func

func PromisePromptResponseObjectOnFulfilledToJS(callback PromisePromptResponseObjectOnFulfilledFunc) *PromisePromptResponseObjectOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromisePromptResponseObjectOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *PromptResponseObject // javascript: PromptResponseObject value
		)
		_p0 = PromptResponseObjectFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromisePromptResponseObjectOnFulfilledFromJS(_value js.Value) PromisePromptResponseObjectOnFulfilledFunc {
	return func(value *PromptResponseObject) {
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
type PromisePromptResponseObjectOnRejectedFunc func(reason js.Value)

// PromisePromptResponseObjectOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromisePromptResponseObjectOnRejected js.Func

func PromisePromptResponseObjectOnRejectedToJS(callback PromisePromptResponseObjectOnRejectedFunc) *PromisePromptResponseObjectOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromisePromptResponseObjectOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromisePromptResponseObjectOnRejectedFromJS(_value js.Value) PromisePromptResponseObjectOnRejectedFunc {
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

// dictionary: PromptResponseObject
type PromptResponseObject struct {
	UserChoice AppBannerPromptOutcome
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *PromptResponseObject) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.UserChoice.JSValue()
	out.Set("userChoice", value0)
	return out
}

// PromptResponseObjectFromJS is allocating a new
// PromptResponseObject object and copy all values from
// input javascript object
func PromptResponseObjectFromJS(value js.Wrapper) *PromptResponseObject {
	input := value.JSValue()
	var out PromptResponseObject
	var (
		value0 AppBannerPromptOutcome // javascript: AppBannerPromptOutcome {userChoice UserChoice userChoice}
	)
	value0 = AppBannerPromptOutcomeFromJS(input.Get("userChoice"))
	out.UserChoice = value0
	return &out
}

// dictionary: ServiceWorkerRegistrationObject
type ServiceWorkerRegistrationObject struct {
	Src            string
	Scope          string
	Type           htmlcommon.WorkerType
	UpdateViaCache serviceworker.ServiceWorkerUpdateViaCache
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ServiceWorkerRegistrationObject) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Src
	out.Set("src", value0)
	value1 := _this.Scope
	out.Set("scope", value1)
	value2 := _this.Type.JSValue()
	out.Set("type", value2)
	value3 := _this.UpdateViaCache.JSValue()
	out.Set("update_via_cache", value3)
	return out
}

// ServiceWorkerRegistrationObjectFromJS is allocating a new
// ServiceWorkerRegistrationObject object and copy all values from
// input javascript object
func ServiceWorkerRegistrationObjectFromJS(value js.Wrapper) *ServiceWorkerRegistrationObject {
	input := value.JSValue()
	var out ServiceWorkerRegistrationObject
	var (
		value0 string                                    // javascript: USVString {src Src src}
		value1 string                                    // javascript: USVString {scope Scope scope}
		value2 htmlcommon.WorkerType                     // javascript: WorkerType {type Type _type}
		value3 serviceworker.ServiceWorkerUpdateViaCache // javascript: ServiceWorkerUpdateViaCache {update_via_cache UpdateViaCache updateViaCache}
	)
	value0 = (input.Get("src")).String()
	out.Src = value0
	value1 = (input.Get("scope")).String()
	out.Scope = value1
	value2 = htmlcommon.WorkerTypeFromJS(input.Get("type"))
	out.Type = value2
	value3 = serviceworker.ServiceWorkerUpdateViaCacheFromJS(input.Get("update_via_cache"))
	out.UpdateViaCache = value3
	return &out
}

// dictionary: WebAppManifest
type WebAppManifest struct {
	Dir                       TextDirectionType
	Lang                      string
	Name                      string
	ShortName                 string
	Description               string
	Icons                     []*appmenifestres.ImageResource
	Screenshots               []*appmenifestres.ImageResource
	Categories                []string
	IarcRatingId              string
	StartUrl                  string
	Display                   DisplayModeType
	Orientation               orientation.OrientationLockType
	ThemeColor                string
	BackgroundColor           string
	Scope                     string
	Serviceworker             *ServiceWorkerRegistrationObject
	RelatedApplications       []*appmenifestres.ExternalApplicationResource
	PreferRelatedApplications bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *WebAppManifest) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Dir.JSValue()
	out.Set("dir", value0)
	value1 := _this.Lang
	out.Set("lang", value1)
	value2 := _this.Name
	out.Set("name", value2)
	value3 := _this.ShortName
	out.Set("short_name", value3)
	value4 := _this.Description
	out.Set("description", value4)
	value5 := js.Global().Get("Array").New(len(_this.Icons))
	for __idx5, __seq_in5 := range _this.Icons {
		__seq_out5 := __seq_in5.JSValue()
		value5.SetIndex(__idx5, __seq_out5)
	}
	out.Set("icons", value5)
	value6 := js.Global().Get("Array").New(len(_this.Screenshots))
	for __idx6, __seq_in6 := range _this.Screenshots {
		__seq_out6 := __seq_in6.JSValue()
		value6.SetIndex(__idx6, __seq_out6)
	}
	out.Set("screenshots", value6)
	value7 := js.Global().Get("Array").New(len(_this.Categories))
	for __idx7, __seq_in7 := range _this.Categories {
		__seq_out7 := __seq_in7
		value7.SetIndex(__idx7, __seq_out7)
	}
	out.Set("categories", value7)
	value8 := _this.IarcRatingId
	out.Set("iarc_rating_id", value8)
	value9 := _this.StartUrl
	out.Set("start_url", value9)
	value10 := _this.Display.JSValue()
	out.Set("display", value10)
	value11 := _this.Orientation.JSValue()
	out.Set("orientation", value11)
	value12 := _this.ThemeColor
	out.Set("theme_color", value12)
	value13 := _this.BackgroundColor
	out.Set("background_color", value13)
	value14 := _this.Scope
	out.Set("scope", value14)
	value15 := _this.Serviceworker.JSValue()
	out.Set("serviceworker", value15)
	value16 := js.Global().Get("Array").New(len(_this.RelatedApplications))
	for __idx16, __seq_in16 := range _this.RelatedApplications {
		__seq_out16 := __seq_in16.JSValue()
		value16.SetIndex(__idx16, __seq_out16)
	}
	out.Set("related_applications", value16)
	value17 := _this.PreferRelatedApplications
	out.Set("prefer_related_applications", value17)
	return out
}

// WebAppManifestFromJS is allocating a new
// WebAppManifest object and copy all values from
// input javascript object
func WebAppManifestFromJS(value js.Wrapper) *WebAppManifest {
	input := value.JSValue()
	var out WebAppManifest
	var (
		value0  TextDirectionType                             // javascript: TextDirectionType {dir Dir dir}
		value1  string                                        // javascript: DOMString {lang Lang lang}
		value2  string                                        // javascript: USVString {name Name name}
		value3  string                                        // javascript: USVString {short_name ShortName shortName}
		value4  string                                        // javascript: USVString {description Description description}
		value5  []*appmenifestres.ImageResource               // javascript: sequence<ImageResource> {icons Icons icons}
		value6  []*appmenifestres.ImageResource               // javascript: sequence<ImageResource> {screenshots Screenshots screenshots}
		value7  []string                                      // javascript: sequence<USVString> {categories Categories categories}
		value8  string                                        // javascript: DOMString {iarc_rating_id IarcRatingId iarcRatingId}
		value9  string                                        // javascript: USVString {start_url StartUrl startUrl}
		value10 DisplayModeType                               // javascript: DisplayModeType {display Display display}
		value11 orientation.OrientationLockType               // javascript: OrientationLockType {orientation Orientation orientation}
		value12 string                                        // javascript: USVString {theme_color ThemeColor themeColor}
		value13 string                                        // javascript: USVString {background_color BackgroundColor backgroundColor}
		value14 string                                        // javascript: USVString {scope Scope scope}
		value15 *ServiceWorkerRegistrationObject              // javascript: ServiceWorkerRegistrationObject {serviceworker Serviceworker serviceworker}
		value16 []*appmenifestres.ExternalApplicationResource // javascript: sequence<ExternalApplicationResource> {related_applications RelatedApplications relatedApplications}
		value17 bool                                          // javascript: boolean {prefer_related_applications PreferRelatedApplications preferRelatedApplications}
	)
	value0 = TextDirectionTypeFromJS(input.Get("dir"))
	out.Dir = value0
	value1 = (input.Get("lang")).String()
	out.Lang = value1
	value2 = (input.Get("name")).String()
	out.Name = value2
	value3 = (input.Get("short_name")).String()
	out.ShortName = value3
	value4 = (input.Get("description")).String()
	out.Description = value4
	__length5 := input.Get("icons").Length()
	__array5 := make([]*appmenifestres.ImageResource, __length5, __length5)
	for __idx5 := 0; __idx5 < __length5; __idx5++ {
		var __seq_out5 *appmenifestres.ImageResource
		__seq_in5 := input.Get("icons").Index(__idx5)
		__seq_out5 = appmenifestres.ImageResourceFromJS(__seq_in5)
		__array5[__idx5] = __seq_out5
	}
	value5 = __array5
	out.Icons = value5
	__length6 := input.Get("screenshots").Length()
	__array6 := make([]*appmenifestres.ImageResource, __length6, __length6)
	for __idx6 := 0; __idx6 < __length6; __idx6++ {
		var __seq_out6 *appmenifestres.ImageResource
		__seq_in6 := input.Get("screenshots").Index(__idx6)
		__seq_out6 = appmenifestres.ImageResourceFromJS(__seq_in6)
		__array6[__idx6] = __seq_out6
	}
	value6 = __array6
	out.Screenshots = value6
	__length7 := input.Get("categories").Length()
	__array7 := make([]string, __length7, __length7)
	for __idx7 := 0; __idx7 < __length7; __idx7++ {
		var __seq_out7 string
		__seq_in7 := input.Get("categories").Index(__idx7)
		__seq_out7 = (__seq_in7).String()
		__array7[__idx7] = __seq_out7
	}
	value7 = __array7
	out.Categories = value7
	value8 = (input.Get("iarc_rating_id")).String()
	out.IarcRatingId = value8
	value9 = (input.Get("start_url")).String()
	out.StartUrl = value9
	value10 = DisplayModeTypeFromJS(input.Get("display"))
	out.Display = value10
	value11 = orientation.OrientationLockTypeFromJS(input.Get("orientation"))
	out.Orientation = value11
	value12 = (input.Get("theme_color")).String()
	out.ThemeColor = value12
	value13 = (input.Get("background_color")).String()
	out.BackgroundColor = value13
	value14 = (input.Get("scope")).String()
	out.Scope = value14
	value15 = ServiceWorkerRegistrationObjectFromJS(input.Get("serviceworker"))
	out.Serviceworker = value15
	__length16 := input.Get("related_applications").Length()
	__array16 := make([]*appmenifestres.ExternalApplicationResource, __length16, __length16)
	for __idx16 := 0; __idx16 < __length16; __idx16++ {
		var __seq_out16 *appmenifestres.ExternalApplicationResource
		__seq_in16 := input.Get("related_applications").Index(__idx16)
		__seq_out16 = appmenifestres.ExternalApplicationResourceFromJS(__seq_in16)
		__array16[__idx16] = __seq_out16
	}
	value16 = __array16
	out.RelatedApplications = value16
	value17 = (input.Get("prefer_related_applications")).Bool()
	out.PreferRelatedApplications = value17
	return &out
}

// class: BeforeInstallPromptEvent
type BeforeInstallPromptEvent struct {
	domcore.Event
}

// BeforeInstallPromptEventFromJS is casting a js.Wrapper into BeforeInstallPromptEvent.
func BeforeInstallPromptEventFromJS(value js.Wrapper) *BeforeInstallPromptEvent {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &BeforeInstallPromptEvent{}
	ret.Value_JS = input
	return ret
}

func NewBeforeInstallPromptEvent(_type string, eventInitDict *domcore.EventInit) (_result *BeforeInstallPromptEvent) {
	_klass := js.Global().Get("BeforeInstallPromptEvent")
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
		_converted *BeforeInstallPromptEvent // javascript: BeforeInstallPromptEvent _what_return_name
	)
	_converted = BeforeInstallPromptEventFromJS(_returned)
	_result = _converted
	return
}

func (_this *BeforeInstallPromptEvent) Prompt() (_result *PromisePromptResponseObject) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("prompt", _args[0:_end]...)
	var (
		_converted *PromisePromptResponseObject // javascript: Promise _what_return_name
	)
	_converted = PromisePromptResponseObjectFromJS(_returned)
	_result = _converted
	return
}

// class: Promise
type PromisePromptResponseObject struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromisePromptResponseObject) JSValue() js.Value {
	return _this.Value_JS
}

// PromisePromptResponseObjectFromJS is casting a js.Wrapper into PromisePromptResponseObject.
func PromisePromptResponseObjectFromJS(value js.Wrapper) *PromisePromptResponseObject {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromisePromptResponseObject{}
	ret.Value_JS = input
	return ret
}

func (_this *PromisePromptResponseObject) Then(onFulfilled *PromisePromptResponseObjectOnFulfilled, onRejected *PromisePromptResponseObjectOnRejected) (_result *PromisePromptResponseObject) {
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
		_converted *PromisePromptResponseObject // javascript: Promise _what_return_name
	)
	_converted = PromisePromptResponseObjectFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromisePromptResponseObject) Catch(onRejected *PromisePromptResponseObjectOnRejected) (_result *PromisePromptResponseObject) {
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
		_converted *PromisePromptResponseObject // javascript: Promise _what_return_name
	)
	_converted = PromisePromptResponseObjectFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromisePromptResponseObject) Finally(onFinally *javascript.PromiseFinally) (_result *PromisePromptResponseObject) {
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
		_converted *PromisePromptResponseObject // javascript: Promise _what_return_name
	)
	_converted = PromisePromptResponseObjectFromJS(_returned)
	_result = _converted
	return
}
