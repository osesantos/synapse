package test

import (
	"testing"
	"time"

	"github.com/osesantos/synapse/synapse"
)

func TestSynapseClient(t *testing.T) {
	natsURL := "nats://localhost:4222"
	client, err := synapse.NewSynapseClient(natsURL)
	if err != nil {
		t.Fatalf("Failed to create Synapse client: %v", err)
	}
	defer client.Close()

	subject := "test.subject"
	msg := synapse.AgentMessage{
		ID:      "test",
		Type:    "test",
		Content: "Hello, Synapse!",
	}

	// Test publishing a message
	if err := client.Publish(subject, msg); err != nil {
		t.Errorf("Failed to publish message: %v", err)
	}

	// Test subscribing to a subject
	handler := func(msg synapse.AgentMessage) error {
		if msg.Content != "Hello, Synapse!" {
			t.Errorf("Received unexpected message content: %s", msg.Content)
		}
		return nil
	}

	subscription, err := client.Subscribe(subject, handler)
	if err != nil {
		t.Fatalf("Failed to subscribe to subject %s: %v", subject, err)
	}
	defer subscription.Unsubscribe()

	time.Sleep(1 * time.Second) // Allow time for the message to be processed
}

func TestSynapseClientErrorHandling(t *testing.T) {
	natsURL := "nats://localhost:4222"
	client, err := synapse.NewSynapseClient(natsURL)
	if err != nil {
		t.Fatalf("Failed to create Synapse client: %v", err)
	}
	defer client.Close()

	// Test publishing with an invalid subject
	invalidSubject := ""
	msg := synapse.AgentMessage{
		ID:      "test",
		Type:    "test",
		Content: "This should fail",
	}

	if err := client.Publish(invalidSubject, msg); err == nil {
		t.Error("Expected error when publishing to an invalid subject, but got none")
	}

	// Test subscribing with an invalid handler
	_, err = client.Subscribe(invalidSubject, nil)
	if err == nil {
		t.Error("Expected error when subscribing with a nil handler, but got none")
	}
}

func TestSynapseClientClose(t *testing.T) {
	natsURL := "nats://localhost:4222"
	client, err := synapse.NewSynapseClient(natsURL)
	if err != nil {
		t.Fatalf("Failed to create Synapse client: %v", err)
	}

	// Test closing the client
	if err := client.Close(); err != nil {
		t.Errorf("Failed to close Synapse client: %v", err)
	}

	// Ensure that further operations fail after closing
	if err := client.Publish("test.subject", synapse.AgentMessage{}); err == nil {
		t.Error("Expected error when publishing after closing the client, but got none")
	}
}
