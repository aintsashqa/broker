package broker_test

import (
	"context"
	"io"
	"testing"

	"github.com/aintsashqa/broker"
)

type BenchSub struct{}

func NewBenchSub() broker.Subscriber {
	return &BenchSub{}
}

func (s *BenchSub) Proceed(io.Reader) {}

func BenchmarkBrokerStartWithZeroSubs(b *testing.B) {
	var (
		brok = broker.New(1000)
		msg  = broker.NewMessage("test-topic", []byte("message"))
	)

	brok.Run(context.Background())

	for i := 0; i < b.N; i++ {
		brok.Pub(&msg)
	}
}

func BenchmarkBrokerStartWithOneSub(b *testing.B) {
	var (
		brok = broker.New(1000)
		sub  = NewBenchSub()
		msg  = broker.NewMessage("test-topic", []byte("message"))
	)

	brok.Sub("test-topic", sub)
	brok.Run(context.Background())

	for i := 0; i < b.N; i++ {
		brok.Pub(&msg)
	}
}
