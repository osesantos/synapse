package transport

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

type NatsClient struct {
	conn *nats.Conn
}

type Subscription = nats.Subscription
type Msg = nats.Msg
type MsgHandler func(msg *Msg) error

var ErrNatsConnectionNotEstablished = errors.New("Nats connection is not established")

func NewNatsClient(natsURL string) (*NatsClient, error) {
	natsClient, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}
	return &NatsClient{conn: natsClient}, nil
}

func (t *NatsClient) Publish(subject string, msg []byte) error {
	if t.conn == nil {
		return fmt.Errorf("Nats connection is not established")
	}
	return t.conn.Publish(subject, msg)
}

func (t *NatsClient) Subscribe(subject string, handler MsgHandler) (*nats.Subscription, error) {
	if t.conn == nil {
		return nil, fmt.Errorf("Nats connection is not established")
	}
	return t.conn.Subscribe(subject, func(msg *nats.Msg) {
		if err := handler(msg); err != nil {
			fmt.Printf("Error handling message: %v\n", err)
		}
	})
}

func (t *NatsClient) Close() error {
	if t.conn == nil {
		return ErrNatsConnectionNotEstablished
	}
	t.conn.Close()
	return nil
}
