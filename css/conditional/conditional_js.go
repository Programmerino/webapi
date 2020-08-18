// Code generated by webidl-bind. DO NOT EDIT.

package conditional

import "syscall/js"

import (
	"github.com/Programmerino/webapi/css/cssom"
)

// using following types:
// cssom.CSSGroupingRule
// cssom.MediaList

// source idl files:
// css-conditional.idl

// transform files:
// css-conditional.go.md

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

// class: CSSConditionRule
type CSSConditionRule struct {
	cssom.CSSGroupingRule
}

// CSSConditionRuleFromJS is casting a js.Wrapper into CSSConditionRule.
func CSSConditionRuleFromJS(value js.Wrapper) *CSSConditionRule {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &CSSConditionRule{}
	ret.Value_JS = input
	return ret
}

// ConditionText returning attribute 'conditionText' with
// type string (idl: DOMString).
func (_this *CSSConditionRule) ConditionText() string {
	var ret string
	value := _this.Value_JS.Get("conditionText")
	ret = (value).String()
	return ret
}

// SetConditionText setting attribute 'conditionText' with
// type string (idl: DOMString).
func (_this *CSSConditionRule) SetConditionText(value string) {
	input := value
	_this.Value_JS.Set("conditionText", input)
}

// class: CSSMediaRule
type CSSMediaRule struct {
	CSSConditionRule
}

// CSSMediaRuleFromJS is casting a js.Wrapper into CSSMediaRule.
func CSSMediaRuleFromJS(value js.Wrapper) *CSSMediaRule {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &CSSMediaRule{}
	ret.Value_JS = input
	return ret
}

// Media returning attribute 'media' with
// type cssom.MediaList (idl: MediaList).
func (_this *CSSMediaRule) Media() *cssom.MediaList {
	var ret *cssom.MediaList
	value := _this.Value_JS.Get("media")
	ret = cssom.MediaListFromJS(value)
	return ret
}

// class: CSSSupportsRule
type CSSSupportsRule struct {
	CSSConditionRule
}

// CSSSupportsRuleFromJS is casting a js.Wrapper into CSSSupportsRule.
func CSSSupportsRuleFromJS(value js.Wrapper) *CSSSupportsRule {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &CSSSupportsRule{}
	ret.Value_JS = input
	return ret
}
