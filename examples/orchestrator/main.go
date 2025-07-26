package orchestrator

// Orchestrator is a simple orchestrator that manages agents and their interactions.

import (
	"fmt"

	"github.com/osesantos/synapse/synapse"
)

type Orchestrator struct {
	Agents      map[string]*synapse.SynapseClient
	NatsURL     string
	StopChannel chan struct{}
}

func NewOrchestrator(natsURL string) *Orchestrator {
	return &Orchestrator{
		Agents:      make(map[string]*synapse.SynapseClient),
		NatsURL:     natsURL,
		StopChannel: make(chan struct{}),
	}
}

func (o *Orchestrator) RegisterAgent(name string) error {
	client, err := synapse.NewSynapseClient(o.NatsURL)
	if err != nil {
		return fmt.Errorf("failed to create Synapse client for agent %s: %w", name, err)
	}

	o.Agents[name] = client
	return nil
}

func (o *Orchestrator) PublishToAgent(agentName, subject string, msg synapse.AgentMessage) error {
	client, exists := o.Agents[agentName]
	if !exists {
		return fmt.Errorf("agent %s not registered", agentName)
	}

	return client.Publish(subject, msg)
}

func (o *Orchestrator) SubscribeToAgent(agentName, subject string, handler synapse.SynapseMessageHandler) error {
	client, exists := o.Agents[agentName]
	if !exists {
		return fmt.Errorf("agent %s not registered", agentName)
	}

	_, err := client.Subscribe(subject, handler)
	if err != nil {
		return fmt.Errorf("failed to subscribe to agent %s on subject %s: %w", agentName, subject, err)
	}

	return nil
}

func (o *Orchestrator) Stop() {
	close(o.StopChannel)
	for _, client := range o.Agents {
		client.Close()
	}
	o.Agents = make(map[string]*synapse.SynapseClient)
}
