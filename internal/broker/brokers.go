package broker

import (
	"context"
	"time"
)

type BrokerConsume interface {
	SetReadDeadline(time.Time)
}

type BrokerProduce interface {
	SetWriteDeadline(time.Time)
}

type Message struct {
	Type        string `json:"type"`
	Payload     []byte `json:"payload"`
	ID          string `json:"id"`
	Queue       string `json:"queue"`
	Timeout     string `json:"timeout"`
	Deadline    string `json:"deadline"`
	UniqueKey   string `json:"unique_key"`
	CompletedAt string `json:"completed_at"`
}

type Broker interface {
	Ping() error
	Close() error
	Enqueue(ctx context.Context, msg *Message) error
	Dequeue(qnames ...string) (*Message, time.Time, error)
	Done(ctx context.Context, msg *Message) error
	MarkAsComplete(ctx context.Context, msg *Message) error
}
