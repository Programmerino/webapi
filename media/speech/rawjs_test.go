package speech_test

import (
	"github.com/Programmerino/webapi"
	"github.com/Programmerino/webapi/media/speech"
)

func ExampleSpeechRecognitionEvent_Emma() {
	var event *speech.SpeechRecognitionEvent

	// cast to correct type
	value := webapi.DocumentFromJS(event.Emma())

	// do something with value
	_ = value
}
