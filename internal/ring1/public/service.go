package public

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync/atomic"

	// SECURE INWARD DEPENDENCY FLOW: Ring 1 explicitly imports Ring 0's Protected 
	// interface boundary layer. Ring 0 never imports Ring 1.
	ring0Contract "github.com/bdra-io/monolith/internal/ring0/protected"
	"github.com/bdra-io/monolith/internal/ring1/protected"
	"github.com/bdra-io/monolith/internal/ring1/pure"
)

type OrderServiceImpl struct {
	repo        protected.OrderRepository
	idGen       protected.IDGenerator
	cache       protected.CacheStore
	identityRing ring0Contract.UserService // Cross-ring interface dependency injection
	isAbsorbed  atomic.Bool
}

func NewOrderService(
	r protected.OrderRepository, 
	id protected.IDGenerator, 
	c protected.CacheStore,
	identity ring0Contract.UserService,
) *OrderServiceImpl {
	s := &OrderServiceImpl{repo: r, idGen: id, cache: c, identityRing: identity}
	s.isAbsorbed.Store(false)
	return s
}

func (s *OrderServiceImpl) SetAbsorbedMode(active bool) {
	s.isAbsorbed.Store(active)
}

func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, userID string, amount float64) (*protected.OrderDTO, error) {
	if s.isAbsorbed.Load() {
		return nil, protected.NewWriteRejectedError("ring1", "transactional-core")
	}

	// EXECUTE CROSS-RING INVARIANT CHECK: Inbound request validation against Ring 0
	user, err := s.identityRing.GetUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("identity verification failed across ring boundary: %w", err)
	}
	if !user.IsActive {
		return nil, errors.New("transaction rejected: target identity profile is suspended")
	}

	id, err := s.idGen.Generate(ctx)
	if err != nil {
		return nil, err
	}

	domainOrder, err := pure.CreateOrder(id, userID, amount)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(ctx, domainOrder); err != nil {
		return nil, err
	}

	return &protected.OrderDTO{ID: domainOrder.ID, Amount: domainOrder.Amount, Status: domainOrder.Status}, nil
}

func (s *OrderServiceImpl) GetOrderCached(ctx context.Context, orderID string) (*protected.OrderDTO, error) {
	order, err := s.repo.Find(ctx, orderID)
	if err != nil {
		cachedBytes, cacheErr := s.cache.Get(ctx, orderID)
		if cacheErr == nil {
			var dto protected.OrderDTO
			if json.Unmarshal(cachedBytes, &dto) == nil {
				return &dto, nil
			}
		}
		return nil, err
	}

	dto := &protected.OrderDTO{ID: order.ID, Amount: order.Amount, Status: order.Status}
	if bytes, marshalErr := json.Marshal(dto); marshalErr == nil {
		_ = s.cache.Set(ctx, orderID, bytes, 300)
	}
	return dto, nil
}