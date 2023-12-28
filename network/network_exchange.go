package p2p

type Message struct {
	MessageType string      `json:"type"`
	MessageBody interface{} `json:"body"`
}

func (m *Message) UnmarshalJSON(buffer []byte) error {
	return nil
}
