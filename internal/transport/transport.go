package transport

type Client interface {
	Publish(subject string, data []byte) error
	Subscribe(subject string, handler func(msg *Msg) error) (*Subscription, error)
	Close() error
}
