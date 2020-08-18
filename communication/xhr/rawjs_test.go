package xhr_test

import (
	"github.com/Programmerino/webapi"
	"github.com/Programmerino/webapi/communication/xhr"
)

func ExampleXMLHttpRequest_ResponseXML() {
	var req *xhr.XMLHttpRequest

	// cast to correct type
	value := webapi.DocumentFromJS(req.ResponseXML())

	// do something with value
	_ = value
}
