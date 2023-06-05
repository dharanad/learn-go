package gcf

type PubSubMessage struct {
	Data []byte `json:"data"`
}

type PublishedMessageData struct {
	Message PubSubMessage
}
