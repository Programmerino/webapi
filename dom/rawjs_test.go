package dom_test

import (
	"github.com/Programmerino/webapi"
	"github.com/Programmerino/webapi/dom"
	"github.com/Programmerino/webapi/html"
)

func ExampleElement_AssignedSlot() {
	var element *dom.Element

	// cast to correct type
	value := html.HTMLSlotElementFromJS(element.AssignedSlot())

	// do something with value
	_ = value
}

func ExampleText_AssignedSlot() {
	var text *dom.Text

	// cast to correct type
	value := html.HTMLSlotElementFromJS(text.AssignedSlot())

	// do something with value
	_ = value
}

func ExampleNode_OwnerDocument() {
	var node *dom.Node

	// cast to correct type
	value := webapi.DocumentFromJS(node.OwnerDocument())

	// do something with value
	_ = value
}
