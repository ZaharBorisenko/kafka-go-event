package repository

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type AnalyticsRepository interface {
	LogEvent(ctx context.Context, eventType string, payload string) error
}

type clickHouseRepo struct {
	conn driver.Conn
}

func NewAnalyticsRepository(conn driver.Conn) AnalyticsRepository {
	return &clickHouseRepo{conn: conn}
}

func (c *clickHouseRepo) LogEvent(ctx context.Context, eventType string, payload string) error {
	// 1. Подготавливаем запрос
	batch, err := c.conn.PrepareBatch(ctx, "INSERT INTOпше  events_log (type, data)")
	if err != nil {
		return fmt.Errorf("failed to prepare batch: %w", err)
	}

	// 2. Добавляем данные в батч
	if err := batch.Append(eventType, payload); err != nil {
		return fmt.Errorf("failed to append to batch: %w", err)
	}

	// 3. Отправляем в базу
	if err := batch.Send(); err != nil {
		return fmt.Errorf("failed to send batch: %w", err)
	}
	fmt.Printf("Successfully logged event: %s\n", eventType)
	return nil
}
