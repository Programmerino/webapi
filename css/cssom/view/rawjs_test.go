package view_test

import (
	"github.com/Programmerino/webapi/css/cssom/view"
	"github.com/Programmerino/webapi/dom"
)

func ExampleCaretPosition_OffsetNode() {
	var caret *view.CaretPosition

	// cast to correct type
	value := dom.NodeFromJS(caret.OffsetNode())

	// do something with value
	_ = value
}
