package events

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	orderdomain "shop-api/domain/order"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
	// pool   *pgxpool.Pool
	repo orderdomain.Repository
}

func NewConsumer(brokers []string, groupID string, repo orderdomain.Repository) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			Topic:    "order.events",
			GroupID:  groupID,
			MinBytes: 10,
			MaxBytes: 10e6,
			MaxWait:  time.Second,
		}),
		repo: repo,
	}
}

func (c *Consumer) Start(ctx context.Context) {
	slog.Info("consumer started", "topic", "order.events")

	// orderDate := event.CreatedAt.Truncate(24 * time.Hour)

	// if _, err := c.pool.Exec(ctx,
	// 	`INSERT INTO daily_purchases (order_date, purchases) VALUES ($1, 1)
	//  ON CONFLICT (order_date) DO UPDATE SET purchases = daily_purchases.purchases + 1`,
	// 	orderDate,
	// ); err != nil {
	// 	slog.Error("upsert daily purchase error", "error", err)
	// }

	// if err := c.repo.UpsertDailyPurchase(ctx, orderDate); err != nil {
	// 	slog.Error("upsert daily purchase error", "error", err)
	// }

	for {
		msg, err := c.reader.ReadMessage(ctx)

		if err != nil {
			if ctx.Err() != nil {
				return
			}

			slog.Error("consumer read error", "error", err)

			continue
		}

		var event OrderCreated

		if err := json.Unmarshal(msg.Value, &event); err != nil {
			slog.Error("consumer unmarshal error", "error", err)

			continue
		}

		slog.Info("order created",
			"order_id", event.OrderID,
			"user_id", event.UserID,
			"partition", msg.Partition,
			"offset", msg.Offset,
		)

		orderDate := event.CreatedAt.Truncate(24 * time.Hour)

		if err := c.repo.UpsertDailyPurchase(ctx, orderDate); err != nil {
			slog.Error("upsert daily purchase error", "error", err)
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}

func StartConsumer(ctx context.Context, brokers []string, groupID string, repo orderdomain.Repository) {
	c := NewConsumer(brokers, groupID, repo)

	defer func() {
		if err := c.Close(); err != nil {
			slog.Error("consumer close error", "error", err)
		}
	}()

	c.Start(ctx)
}
