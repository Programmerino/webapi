// Code generated by webidl-bind. DO NOT EDIT.

package csp

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
)

// using following types:
// domcore.Event

// source idl files:
// CSP.idl

// transform files:
// CSP.go.md

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

// enum: SecurityPolicyViolationEventDisposition
type SecurityPolicyViolationEventDisposition int

const (
	EnforceSecurityPolicyViolationEventDisposition SecurityPolicyViolationEventDisposition = iota
	ReportSecurityPolicyViolationEventDisposition
)

var securityPolicyViolationEventDispositionToWasmTable = []string{
	"enforce", "report",
}

var securityPolicyViolationEventDispositionFromWasmTable = map[string]SecurityPolicyViolationEventDisposition{
	"enforce": EnforceSecurityPolicyViolationEventDisposition, "report": ReportSecurityPolicyViolationEventDisposition,
}

// JSValue is converting this enum into a javascript object
func (this *SecurityPolicyViolationEventDisposition) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this SecurityPolicyViolationEventDisposition) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(securityPolicyViolationEventDispositionToWasmTable) {
		return securityPolicyViolationEventDispositionToWasmTable[idx]
	}
	panic("unknown input value")
}

// SecurityPolicyViolationEventDispositionFromJS is converting a javascript value into
// a SecurityPolicyViolationEventDisposition enum value.
func SecurityPolicyViolationEventDispositionFromJS(value js.Value) SecurityPolicyViolationEventDisposition {
	key := value.String()
	conv, ok := securityPolicyViolationEventDispositionFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: SecurityPolicyViolationEventInit
type SecurityPolicyViolationEventInit struct {
	Bubbles            bool
	Cancelable         bool
	Composed           bool
	DocumentURL        string
	Referrer           string
	BlockedURL         string
	EffectiveDirective string
	OriginalPolicy     string
	SourceFile         string
	Sample             string
	Disposition        SecurityPolicyViolationEventDisposition
	StatusCode         int
	Lineno             uint
	Colno              uint
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *SecurityPolicyViolationEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.DocumentURL
	out.Set("documentURL", value3)
	value4 := _this.Referrer
	out.Set("referrer", value4)
	value5 := _this.BlockedURL
	out.Set("blockedURL", value5)
	value6 := _this.EffectiveDirective
	out.Set("effectiveDirective", value6)
	value7 := _this.OriginalPolicy
	out.Set("originalPolicy", value7)
	value8 := _this.SourceFile
	out.Set("sourceFile", value8)
	value9 := _this.Sample
	out.Set("sample", value9)
	value10 := _this.Disposition.JSValue()
	out.Set("disposition", value10)
	value11 := _this.StatusCode
	out.Set("statusCode", value11)
	value12 := _this.Lineno
	out.Set("lineno", value12)
	value13 := _this.Colno
	out.Set("colno", value13)
	return out
}

// SecurityPolicyViolationEventInitFromJS is allocating a new
// SecurityPolicyViolationEventInit object and copy all values from
// input javascript object
func SecurityPolicyViolationEventInitFromJS(value js.Wrapper) *SecurityPolicyViolationEventInit {
	input := value.JSValue()
	var out SecurityPolicyViolationEventInit
	var (
		value0  bool                                    // javascript: boolean {bubbles Bubbles bubbles}
		value1  bool                                    // javascript: boolean {cancelable Cancelable cancelable}
		value2  bool                                    // javascript: boolean {composed Composed composed}
		value3  string                                  // javascript: USVString {documentURL DocumentURL documentURL}
		value4  string                                  // javascript: USVString {referrer Referrer referrer}
		value5  string                                  // javascript: USVString {blockedURL BlockedURL blockedURL}
		value6  string                                  // javascript: DOMString {effectiveDirective EffectiveDirective effectiveDirective}
		value7  string                                  // javascript: DOMString {originalPolicy OriginalPolicy originalPolicy}
		value8  string                                  // javascript: USVString {sourceFile SourceFile sourceFile}
		value9  string                                  // javascript: DOMString {sample Sample sample}
		value10 SecurityPolicyViolationEventDisposition // javascript: SecurityPolicyViolationEventDisposition {disposition Disposition disposition}
		value11 int                                     // javascript: unsigned short {statusCode StatusCode statusCode}
		value12 uint                                    // javascript: unsigned long {lineno Lineno lineno}
		value13 uint                                    // javascript: unsigned long {colno Colno colno}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = (input.Get("documentURL")).String()
	out.DocumentURL = value3
	value4 = (input.Get("referrer")).String()
	out.Referrer = value4
	value5 = (input.Get("blockedURL")).String()
	out.BlockedURL = value5
	value6 = (input.Get("effectiveDirective")).String()
	out.EffectiveDirective = value6
	value7 = (input.Get("originalPolicy")).String()
	out.OriginalPolicy = value7
	value8 = (input.Get("sourceFile")).String()
	out.SourceFile = value8
	value9 = (input.Get("sample")).String()
	out.Sample = value9
	value10 = SecurityPolicyViolationEventDispositionFromJS(input.Get("disposition"))
	out.Disposition = value10
	value11 = (input.Get("statusCode")).Int()
	out.StatusCode = value11
	value12 = (uint)((input.Get("lineno")).Int())
	out.Lineno = value12
	value13 = (uint)((input.Get("colno")).Int())
	out.Colno = value13
	return &out
}

// class: SecurityPolicyViolationEvent
type SecurityPolicyViolationEvent struct {
	domcore.Event
}

// SecurityPolicyViolationEventFromJS is casting a js.Wrapper into SecurityPolicyViolationEvent.
func SecurityPolicyViolationEventFromJS(value js.Wrapper) *SecurityPolicyViolationEvent {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &SecurityPolicyViolationEvent{}
	ret.Value_JS = input
	return ret
}

func NewSecurityPolicyViolationEvent(_type string, eventInitDict *SecurityPolicyViolationEventInit) (_result *SecurityPolicyViolationEvent) {
	_klass := js.Global().Get("SecurityPolicyViolationEvent")
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
		_converted *SecurityPolicyViolationEvent // javascript: SecurityPolicyViolationEvent _what_return_name
	)
	_converted = SecurityPolicyViolationEventFromJS(_returned)
	_result = _converted
	return
}

// DocumentURL returning attribute 'documentURL' with
// type string (idl: USVString).
func (_this *SecurityPolicyViolationEvent) DocumentURL() string {
	var ret string
	value := _this.Value_JS.Get("documentURL")
	ret = (value).String()
	return ret
}

// DocumentURI returning attribute 'documentURI' with
// type string (idl: USVString).
func (_this *SecurityPolicyViolationEvent) DocumentURI() string {
	var ret string
	value := _this.Value_JS.Get("documentURI")
	ret = (value).String()
	return ret
}

// Referrer returning attribute 'referrer' with
// type string (idl: USVString).
func (_this *SecurityPolicyViolationEvent) Referrer() string {
	var ret string
	value := _this.Value_JS.Get("referrer")
	ret = (value).String()
	return ret
}

// BlockedURL returning attribute 'blockedURL' with
// type string (idl: USVString).
func (_this *SecurityPolicyViolationEvent) BlockedURL() string {
	var ret string
	value := _this.Value_JS.Get("blockedURL")
	ret = (value).String()
	return ret
}

// BlockedURI returning attribute 'blockedURI' with
// type string (idl: USVString).
func (_this *SecurityPolicyViolationEvent) BlockedURI() string {
	var ret string
	value := _this.Value_JS.Get("blockedURI")
	ret = (value).String()
	return ret
}

// EffectiveDirective returning attribute 'effectiveDirective' with
// type string (idl: DOMString).
func (_this *SecurityPolicyViolationEvent) EffectiveDirective() string {
	var ret string
	value := _this.Value_JS.Get("effectiveDirective")
	ret = (value).String()
	return ret
}

// ViolatedDirective returning attribute 'violatedDirective' with
// type string (idl: DOMString).
func (_this *SecurityPolicyViolationEvent) ViolatedDirective() string {
	var ret string
	value := _this.Value_JS.Get("violatedDirective")
	ret = (value).String()
	return ret
}

// OriginalPolicy returning attribute 'originalPolicy' with
// type string (idl: DOMString).
func (_this *SecurityPolicyViolationEvent) OriginalPolicy() string {
	var ret string
	value := _this.Value_JS.Get("originalPolicy")
	ret = (value).String()
	return ret
}

// SourceFile returning attribute 'sourceFile' with
// type string (idl: USVString).
func (_this *SecurityPolicyViolationEvent) SourceFile() string {
	var ret string
	value := _this.Value_JS.Get("sourceFile")
	ret = (value).String()
	return ret
}

// Sample returning attribute 'sample' with
// type string (idl: DOMString).
func (_this *SecurityPolicyViolationEvent) Sample() string {
	var ret string
	value := _this.Value_JS.Get("sample")
	ret = (value).String()
	return ret
}

// Disposition returning attribute 'disposition' with
// type SecurityPolicyViolationEventDisposition (idl: SecurityPolicyViolationEventDisposition).
func (_this *SecurityPolicyViolationEvent) Disposition() SecurityPolicyViolationEventDisposition {
	var ret SecurityPolicyViolationEventDisposition
	value := _this.Value_JS.Get("disposition")
	ret = SecurityPolicyViolationEventDispositionFromJS(value)
	return ret
}

// StatusCode returning attribute 'statusCode' with
// type int (idl: unsigned short).
func (_this *SecurityPolicyViolationEvent) StatusCode() int {
	var ret int
	value := _this.Value_JS.Get("statusCode")
	ret = (value).Int()
	return ret
}

// Lineno returning attribute 'lineno' with
// type uint (idl: unsigned long).
func (_this *SecurityPolicyViolationEvent) Lineno() uint {
	var ret uint
	value := _this.Value_JS.Get("lineno")
	ret = (uint)((value).Int())
	return ret
}

// LineNumber returning attribute 'lineNumber' with
// type uint (idl: unsigned long).
func (_this *SecurityPolicyViolationEvent) LineNumber() uint {
	var ret uint
	value := _this.Value_JS.Get("lineNumber")
	ret = (uint)((value).Int())
	return ret
}

// Colno returning attribute 'colno' with
// type uint (idl: unsigned long).
func (_this *SecurityPolicyViolationEvent) Colno() uint {
	var ret uint
	value := _this.Value_JS.Get("colno")
	ret = (uint)((value).Int())
	return ret
}

// ColumnNumber returning attribute 'columnNumber' with
// type uint (idl: unsigned long).
func (_this *SecurityPolicyViolationEvent) ColumnNumber() uint {
	var ret uint
	value := _this.Value_JS.Get("columnNumber")
	ret = (uint)((value).Int())
	return ret
}
