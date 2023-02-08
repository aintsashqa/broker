package broker

type Message struct {
	topic string
	bytes []byte
}

func NewMessage(topic string, bytes []byte) Message {
	return Message{
		topic: topic,
		bytes: bytes,
	}
}

func (m *Message) Read(buff []byte) (int, error) {
	// buff = append(buff, m.bytes...)
	return len(buff), nil
}
