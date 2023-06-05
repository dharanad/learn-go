package gcf

import (
	"context"
	"encoding/json"
	"github.com/cloudevents/sdk-go/v2/event"
	"io"
	"log"
	"os"
	"testing"
)

func TestHelloPubSubFunction(t *testing.T) {
	r, w, _ := os.Pipe()
	log.SetOutput(w)
	originalFlags := log.Flags()
	// removing data & time from log output
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	mockEvent := mockCloudEventFromString("Dharan", nil)
	err := HelloPubSubFunction(context.Background(), mockEvent)
	if err != nil {
		t.Errorf("Expected a nil error")
	}
	w.Close()
	log.SetOutput(os.Stderr)
	log.SetFlags(originalFlags)

	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	want := "Hello, Dharan\n"
	if got := string(out); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func mockCloudEventFromStruct(data any, attribute map[any]any) event.Event {
	byteData, _ := json.Marshal(data)
	type pubSubMessageData struct {
		Data       []byte      `json:"data"`
		Attributes map[any]any `json:"attributes,omitempty"`
	}
	type publishedPubSubMessage struct {
		Message pubSubMessageData
	}
	e := event.New()
	e.SetDataContentType("application/json")
	_ = e.SetData(e.DataContentType(), publishedPubSubMessage{
		Message: pubSubMessageData{
			Data:       byteData,
			Attributes: attribute,
		},
	})
	return e
}

func mockCloudEventFromString(data string, attribute map[any]any) event.Event {
	type pubSubMessageData struct {
		Data       []byte      `json:"data"`
		Attributes map[any]any `json:"attributes,omitempty"`
	}
	type publishedPubSubMessage struct {
		Message pubSubMessageData
	}
	e := event.New()
	e.SetDataContentType("application/json")
	_ = e.SetData(e.DataContentType(), publishedPubSubMessage{
		Message: pubSubMessageData{
			Data:       []byte(data),
			Attributes: attribute,
		},
	})
	return e
}
