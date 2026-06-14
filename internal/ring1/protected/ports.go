package protected

import (
	"context"
	
	// INTRA-RING IMPORT RULE: Permitted under BDRA Lite constraints. 
	// The Protected layer of a Ring is explicitly allowed to import the Pure 
	// layer of the SAME Ring to expose core validation models directly.
	"github.com/bdra-io/monolith/internal/ring1/pure"
)

type OrderDTO struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

type OrderService interface {
	PlaceOrder(ctx context.Context, userID string, amount float64) (*OrderDTO, error)
	GetOrderCached(ctx context.Context, orderID string) (*OrderDTO, error)
}

type OrderRepository interface {
	Save(ctx context.Context, order *pure.Order) error
	Find(ctx context.Context, id string) (*pure.Order, error)
}

type CacheStore interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte, ttl int) error
}

type IDGenerator interface {
	Generate(ctx context.Context) (string, error)
}