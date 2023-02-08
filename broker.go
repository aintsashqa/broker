package broker

import (
	"context"
)

type Broker struct {
	stream chan *Message
	subs   map[string][]Subscriber
}

func New(size uint16) Broker {
	return Broker{
		stream: make(chan *Message, size),
		subs:   make(map[string][]Subscriber),
	}
}

func (b *Broker) Sub(topic string, sub Subscriber) {
	b.subs[topic] = append(b.subs[topic], sub)
}

func (b *Broker) Pub(msg *Message) {
	b.stream <- msg
}

func (b *Broker) Start(ctx context.Context) {
	defer func() { close(b.stream) }()

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-b.stream:
			subs := b.subs[msg.topic]
			for _, sub := range subs {
				sub.Proceed(msg)
			}
		}
	}
}

func (b *Broker) Run(ctx context.Context) {
	go b.Start(ctx)
}
