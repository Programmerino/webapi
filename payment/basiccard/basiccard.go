// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package basiccard

import js "github.com/Programmerino/webapi/core/js"

import (
	"github.com/Programmerino/webapi/payment/request"
)

// using following types:
// request.AddressErrors
// request.PaymentAddress

// source idl files:
// payment-method-basic-card.idl

// transform files:
// payment-method-basic-card.go.md

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

// dictionary: BasicCardChangeDetails
type BasicCardChangeDetails struct {
	BillingAddress *request.PaymentAddress
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *BasicCardChangeDetails) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.BillingAddress.JSValue()
	out.Set("billingAddress", value0)
	return out
}

// BasicCardChangeDetailsFromJS is allocating a new
// BasicCardChangeDetails object and copy all values from
// input javascript object
func BasicCardChangeDetailsFromJS(value js.Wrapper) *BasicCardChangeDetails {
	input := value.JSValue()
	var out BasicCardChangeDetails
	var (
		value0 *request.PaymentAddress // javascript: PaymentAddress {billingAddress BillingAddress billingAddress}
	)
	if input.Get("billingAddress").Type() != js.TypeNull && input.Get("billingAddress").Type() != js.TypeUndefined {
		value0 = request.PaymentAddressFromJS(input.Get("billingAddress"))
	}
	out.BillingAddress = value0
	return &out
}

// dictionary: BasicCardErrors
type BasicCardErrors struct {
	CardNumber       string
	CardholderName   string
	CardSecurityCode string
	ExpiryMonth      string
	ExpiryYear       string
	BillingAddress   *request.AddressErrors
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *BasicCardErrors) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.CardNumber
	out.Set("cardNumber", value0)
	value1 := _this.CardholderName
	out.Set("cardholderName", value1)
	value2 := _this.CardSecurityCode
	out.Set("cardSecurityCode", value2)
	value3 := _this.ExpiryMonth
	out.Set("expiryMonth", value3)
	value4 := _this.ExpiryYear
	out.Set("expiryYear", value4)
	value5 := _this.BillingAddress.JSValue()
	out.Set("billingAddress", value5)
	return out
}

// BasicCardErrorsFromJS is allocating a new
// BasicCardErrors object and copy all values from
// input javascript object
func BasicCardErrorsFromJS(value js.Wrapper) *BasicCardErrors {
	input := value.JSValue()
	var out BasicCardErrors
	var (
		value0 string                 // javascript: DOMString {cardNumber CardNumber cardNumber}
		value1 string                 // javascript: DOMString {cardholderName CardholderName cardholderName}
		value2 string                 // javascript: DOMString {cardSecurityCode CardSecurityCode cardSecurityCode}
		value3 string                 // javascript: DOMString {expiryMonth ExpiryMonth expiryMonth}
		value4 string                 // javascript: DOMString {expiryYear ExpiryYear expiryYear}
		value5 *request.AddressErrors // javascript: AddressErrors {billingAddress BillingAddress billingAddress}
	)
	value0 = (input.Get("cardNumber")).String()
	out.CardNumber = value0
	value1 = (input.Get("cardholderName")).String()
	out.CardholderName = value1
	value2 = (input.Get("cardSecurityCode")).String()
	out.CardSecurityCode = value2
	value3 = (input.Get("expiryMonth")).String()
	out.ExpiryMonth = value3
	value4 = (input.Get("expiryYear")).String()
	out.ExpiryYear = value4
	value5 = request.AddressErrorsFromJS(input.Get("billingAddress"))
	out.BillingAddress = value5
	return &out
}

// dictionary: BasicCardRequest
type BasicCardRequest struct {
	SupportedNetworks []string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *BasicCardRequest) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.SupportedNetworks))
	for __idx0, __seq_in0 := range _this.SupportedNetworks {
		__seq_out0 := __seq_in0
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("supportedNetworks", value0)
	return out
}

// BasicCardRequestFromJS is allocating a new
// BasicCardRequest object and copy all values from
// input javascript object
func BasicCardRequestFromJS(value js.Wrapper) *BasicCardRequest {
	input := value.JSValue()
	var out BasicCardRequest
	var (
		value0 []string // javascript: sequence<DOMString> {supportedNetworks SupportedNetworks supportedNetworks}
	)
	__length0 := input.Get("supportedNetworks").Length()
	__array0 := make([]string, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 string
		__seq_in0 := input.Get("supportedNetworks").Index(__idx0)
		__seq_out0 = (__seq_in0).String()
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.SupportedNetworks = value0
	return &out
}

// dictionary: BasicCardResponse
type BasicCardResponse struct {
	CardNumber       string
	CardholderName   string
	CardSecurityCode string
	ExpiryMonth      string
	ExpiryYear       string
	BillingAddress   *request.PaymentAddress
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *BasicCardResponse) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.CardNumber
	out.Set("cardNumber", value0)
	value1 := _this.CardholderName
	out.Set("cardholderName", value1)
	value2 := _this.CardSecurityCode
	out.Set("cardSecurityCode", value2)
	value3 := _this.ExpiryMonth
	out.Set("expiryMonth", value3)
	value4 := _this.ExpiryYear
	out.Set("expiryYear", value4)
	value5 := _this.BillingAddress.JSValue()
	out.Set("billingAddress", value5)
	return out
}

// BasicCardResponseFromJS is allocating a new
// BasicCardResponse object and copy all values from
// input javascript object
func BasicCardResponseFromJS(value js.Wrapper) *BasicCardResponse {
	input := value.JSValue()
	var out BasicCardResponse
	var (
		value0 string                  // javascript: DOMString {cardNumber CardNumber cardNumber}
		value1 string                  // javascript: DOMString {cardholderName CardholderName cardholderName}
		value2 string                  // javascript: DOMString {cardSecurityCode CardSecurityCode cardSecurityCode}
		value3 string                  // javascript: DOMString {expiryMonth ExpiryMonth expiryMonth}
		value4 string                  // javascript: DOMString {expiryYear ExpiryYear expiryYear}
		value5 *request.PaymentAddress // javascript: PaymentAddress {billingAddress BillingAddress billingAddress}
	)
	value0 = (input.Get("cardNumber")).String()
	out.CardNumber = value0
	value1 = (input.Get("cardholderName")).String()
	out.CardholderName = value1
	value2 = (input.Get("cardSecurityCode")).String()
	out.CardSecurityCode = value2
	value3 = (input.Get("expiryMonth")).String()
	out.ExpiryMonth = value3
	value4 = (input.Get("expiryYear")).String()
	out.ExpiryYear = value4
	if input.Get("billingAddress").Type() != js.TypeNull && input.Get("billingAddress").Type() != js.TypeUndefined {
		value5 = request.PaymentAddressFromJS(input.Get("billingAddress"))
	}
	out.BillingAddress = value5
	return &out
}
