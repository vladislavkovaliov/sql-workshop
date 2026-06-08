package events

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Balancer: &kafka.Hash{},
		},
	}
}

func (p *Producer) PublishOrderCreated(ctx context.Context, event OrderCreated) error {
	value, err := json.Marshal(event)

	if err != nil {
		return fmt.Errorf("marshal event: %w", err)
	}

	return p.writer.WriteMessages(ctx, kafka.Message{
		Topic: "order.events",
		Key:   []byte(fmt.Sprintf("%d", event.OrderID)),
		Value: value,
	})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
