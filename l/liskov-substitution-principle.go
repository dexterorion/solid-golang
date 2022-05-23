func sendMessage(mBroker MessageBroker, ctx context.Context, message interface{}, topic string) {
  mBroker.Publish(ctx, message, topic)
}

func main() {
  rabbitMQBroker := &RabbitMQBroker{}
  kafkaBroker := &KafkaBroker{}
  pubSubBroker := &PubSubBroker{}
  
  ctx := context.Background()
  message := map[string]string{
    "name": "Test",
  }
  topic := "Test"
  
  // send 1
  sendMessage(rabbitMQBroker, ctx, message, topic)
  
  // send 2
  sendMessage(kafkaBroker, ctx, message, topic)
  
  // send 3
  sendMessage(pubSubBroker, ctx, message, topic)
}
