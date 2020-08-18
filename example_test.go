package webapi_test

import (
	"fmt"

	"github.com/Programmerino/webapi"
	"github.com/Programmerino/webapi/html"
	"github.com/Programmerino/webapi/html/htmlevent"
)

func Example_buttonTest() {
	element := webapi.GetWindow().Document().GetElementById("foo")
	button := html.HTMLButtonElementFromJS(element)
	button.SetInnerText("Press me!")

	count := 1
	callback := func(event *htmlevent.MouseEvent, currentTarget *html.HTMLElement) {
		button.SetInnerText(fmt.Sprint("Count: ", count))
		count++
	}
	jsFunction := button.SetOnClick(callback)

	_ = jsFunction

	// prevent to program to terminate
	// c := make(chan struct{}, 0)
	// <-c
}
