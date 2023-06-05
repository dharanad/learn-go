package gcf

import (
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"log"
)

func HelloPubSubFunction(ctx context.Context, e event.Event) error {
	var msg PublishedMessageData
	if err := e.DataAs(&msg); err != nil {
		return err
	}
	data := string(msg.Message.Data)
	log.Printf(fmt.Sprintf("Hello, %s", data))
	return nil
}
