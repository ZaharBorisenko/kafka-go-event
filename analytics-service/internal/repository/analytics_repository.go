package repository

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type AnalyticsRepository interface {
	LogEvent(ctx context.Context, eventType string, payload string) error
	SaveDeposit(ctx context.Context, userID string, amount float64, opType string) error
}

type clickHouseRepo struct {
	conn driver.Conn
}

func NewAnalyticsRepository(conn driver.Conn) AnalyticsRepository {
	return &clickHouseRepo{conn: conn}
}

func (c *clickHouseRepo) LogEvent(ctx context.Context, eventType string, payload string) error {
	batch, err := c.conn.PrepareBatch(ctx, "INSERT INTO events_log (type, data)")
	if err != nil {
		return fmt.Errorf("failed to prepare batch: %w", err)
	}
	if err := batch.Append(eventType, payload); err != nil {
		return fmt.Errorf("failed to append to batch: %w", err)
	}
	if err := batch.Send(); err != nil {
		return fmt.Errorf("failed to send batch: %w", err)
	}
	fmt.Printf("Successfully logged event: %s\n", eventType)
	return nil
}

func (c *clickHouseRepo) SaveDeposit(ctx context.Context, userID string, amount float64, opType string) error {
	batch, err := c.conn.PrepareBatch(ctx, "INSERT INTO deposits (user_id, amount, type)")
	if err != nil {
		return fmt.Errorf("failed to prepare batch: %w", err)
	}

	if err := batch.Append(userID, amount, opType); err != nil {
		return fmt.Errorf("failed to append to batch: %w", err)
	}

	if err := batch.Send(); err != nil {
		return fmt.Errorf("failed to send batch: %w", err)
	}

	fmt.Println("Successfully SaveDeposit", opType)

	return nil
}
