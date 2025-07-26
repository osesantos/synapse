# 🧠 Synapse

**Synapse** is a lightweight, pub/sub-based communication protocol and Go client for intelligent agents, LLMs, and autonomous systems.

Built for **contextual messaging**, **fan-out/fan-in**, **streaming**, and **modular orchestration** — Synapse lets your agents think, act, and respond like neurons in a high-performance AI brain.

> Think: fast, semantic, decentralized message passing — designed for LLMs and beyond.

---

## 🚀 Features

- 🧠 Agent-oriented message format (sender, receiver, context, tools, etc.)
- 📡 Pub/Sub communication (transport-agnostic)
- 🔁 Fan-out and fan-in friendly
- 🧩 Modular and pluggable by design
- ⚙️ Supports streaming, retries, and metadata
- 💡 Simple Go API for building real-time agent ecosystems

---

## 📦 Install

```bash
go get github.com/osesantos/synapse
```

---

## ✨ Example

```go
import "github.com/osesantos/synapse"

client, _ := synapse.NewSynapseClient("nats://localhost:4222")

msg := synapse.AgentMessage{
    ID:        "msg-001",
    Sender:    "agent.weather",
    Receiver:  "agent.calendar",
    Type:      "question",
    Content:   "What is the weather today?",
    ReplyTo:   "mcp.responses",
    ContextID: "session-abc",
}

client.Publish("agents.request", msg)
```

---

## 🧠 Message Format

```json
{
  "id": "msg-001",
  "type": "question",
  "sender": "agent.weather",
  "receiver": "agent.calendar",
  "content": "What is the weather today?",
  "reply_to": "mcp.responses",
  "context_id": "session-abc",
  "stream": true,
  "metadata": {
    "lang": "pt",
    "origin": "gomind"
  }
}
```

---

## 🔌 Transport Agnostic

Synapse works over any pub/sub backend.  
By default it uses [NATS](https://nats.io), but can be extended to work with other message systems like MQTT, Redis Streams, Kafka, or your own custom transport.

---

## 💬 License

MIT — built with ❤️ for agent-based systems.

