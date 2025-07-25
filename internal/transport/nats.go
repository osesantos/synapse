package transport

import "fmt"

type NatsClient struct {
	conn *nats.Conn
}

func (t *NatsClient) Publish(subject string, msg []byte) error {
	if t.conn == nil {
		return fmt.Errorf("Nats connection is not established")
	}
	return t.conn.Publish(subject, msg)
}

func (t *NatsClient) Subscribe(subject string, cb nats.MsgHandler) error {
	if t.conn == nil {
		return fmt.Errorf("Nats connection is not established")
	}
	return t.conn.Subscribe(subject, cb)
}
