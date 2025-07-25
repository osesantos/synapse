package synapse

type AgentMessage struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	Sender        string            `json:"sender"`
	Receiver      string            `json:"receiver"`
	ContextID     string            `json:"context_id"`
	ReplyTo       string            `json:"reply_to"`
	Content       string            `json:"content"`
	Metadata      map[string]string `json:"metadata"`
	Tools         []string          `json:"tools"`
	Stream        bool              `json:"stream"`
	TimeoutMs     int               `json:"timeout_ms"`
	CorrelationID string            `json:"correlation_id"`
	Timestamp     int64             `json:"timestamp"`
}
