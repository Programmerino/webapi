// Code generated by webidl-bind. DO NOT EDIT.

package fromelement

import "syscall/js"

import (
	"github.com/gowebapi/webapi/html/canvas"
	"github.com/gowebapi/webapi/media/capture/local"
)

// using following types:
// canvas.HTMLCanvasElement
// local.MediaStreamTrack

// source idl files:
// mediacapture-fromelement.idl

// transform files:
// mediacapture-fromelement.go.md

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

// class: CanvasCaptureMediaStreamTrack
type CanvasCaptureMediaStreamTrack struct {
	local.MediaStreamTrack
}

// CanvasCaptureMediaStreamTrackFromJS is casting a js.Wrapper into CanvasCaptureMediaStreamTrack.
func CanvasCaptureMediaStreamTrackFromJS(value js.Wrapper) *CanvasCaptureMediaStreamTrack {
	input := value.JSValue()
	if typ := input.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &CanvasCaptureMediaStreamTrack{}
	ret.Value_JS = input
	return ret
}

// Canvas returning attribute 'canvas' with
// type canvas.HTMLCanvasElement (idl: HTMLCanvasElement).
func (_this *CanvasCaptureMediaStreamTrack) Canvas() *canvas.HTMLCanvasElement {
	var ret *canvas.HTMLCanvasElement
	value := _this.Value_JS.Get("canvas")
	ret = canvas.HTMLCanvasElementFromJS(value)
	return ret
}

func (_this *CanvasCaptureMediaStreamTrack) RequestFrame() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("requestFrame", _args[0:_end]...)
	return
}
