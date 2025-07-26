package agentping

// AgentPing is a simple agent that pings a specified subject at regular intervals.

import (
	"time"

	"github.com/osesantos/synapse/synapse"
)

type AgentPing struct {
	Subject     string
	Interval    time.Duration
	StopChannel chan struct{}
	client      *synapse.SynapseClient
}

func NewAgentPing(subject string, interval time.Duration, natsURL string) (*AgentPing, error) {
	client, err := synapse.NewSynapseClient(natsURL)
	if err != nil {
		return nil, err
	}

	return &AgentPing{
		Subject:     subject,
		Interval:    interval,
		StopChannel: make(chan struct{}),
		client:      client,
	}, nil
}

func (a *AgentPing) Start() {
	ticker := time.NewTicker(a.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			msg := synapse.AgentMessage{
				ID:      "ping",
				Type:    "ping",
				Content: "Ping from AgentPing",
			}

			if err := a.client.Publish(a.Subject, msg); err != nil {
				// Handle error (e.g., log it)
			}
		case <-a.StopChannel:
			return
		}
	}
}
