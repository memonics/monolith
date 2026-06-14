package public

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/bdra-io/monolith/internal/ring1/pure"
)

type PostgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(db *sql.DB) *PostgresOrderRepository {
	return &PostgresOrderRepository{db: db}
}

func (r *PostgresOrderRepository) Save(ctx context.Context, order *pure.Order) error {
	// In production, execute standard sql.DB transaction execution blocks here
	return nil
}

func (r *PostgresOrderRepository) Find(ctx context.Context, id string) (*pure.Order, error) {
	// Simulating a storage infrastructure driver fault for testing purposes
	if id == "ord_corrupted_db" {
		return nil, errors.New("sql: database connection lost or connection pool exhausted")
	}

	return &pure.Order{
		ID:        id,
		UserID:    "usr_007",
		Amount:    450.00,
		Status:    "PENDING",
		CreatedAt: time.Now().UTC(),
	}, nil
}