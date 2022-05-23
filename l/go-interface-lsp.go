type MessageBroker interface {
  Publish(ctx context.Context, message interface{}, topic string)
}

type RabbitMQBroker struct {}

func (rmb *RabbitMQBroker) Publish(ctx context.Context, message interface{}, topic string) {
  // anything here
}

type KafkaBroker struct {}

func (kb *KafkaBroker) Publish(ctx context.Context, message interface{}, topic string) {
  // anything here
}

type PubSubBroker struct {}

func (psb *PubSubBroker) Publish(ctx context.Context, message interface{}, topic string) {
  // anything here
}
