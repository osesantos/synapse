package synapse

import (
	"encoding/json"

	"github.com/osesantos/synapse/internal/transport"
)

type SynapseClient struct {
	conn *transport.NatsClient
}

type SynapseMessageHandler func(msg AgentMessage) error

func NewSynapseClient(natsURL string) (*SynapseClient, error) {
	natsClient, err := transport.NewClient(natsURL)
	if err != nil {
		return nil, err
	}
	return NewSynapseClientWithConn(natsClient), nil
}

func NewSynapseClientWithConn(conn *transport.Client) *SynapseClient {
	return &SynapseClient{conn: conn}
}

func (s *SynapseClient) Publish(subject string, msg AgentMessage) error {
	if s.conn == nil {
		return transport.ErrNatsConnectionNotEstablished
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return s.conn.Publish(subject, data)
}

func (s *SynapseClient) Subscribe(subject string, handler SynapseMessageHandler) (*transport.Subscription, error) {
	if s.conn == nil {
		return nil, transport.ErrNatsConnectionNotEstablished
	}

	return s.conn.Subscribe(subject, func(msg *transport.Msg) error {
		var agentMsg AgentMessage
		if err := json.Unmarshal(msg.Data, &msg); err != nil {
			return err
		}

		if err := handler(agentMsg); err != nil {
			return err
		}

		return nil
	})
}

func (s *SynapseClient) Close() error {
	if s.conn == nil {
		return transport.ErrNatsConnectionNotEstablished
	}
	return s.conn.Close()
}
