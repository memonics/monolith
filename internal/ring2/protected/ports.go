package protected

import (
	"context"

	// SECURE INWARD DEPENDENCY: Ring 2 is legally permitted to import Ring 1's DTOs.
	ring1Contract "github.com/bdra-io/monolith/internal/ring1/protected"
)

type AuditLogService interface {
	LogOrderCreation(ctx context.Context, order *ring1Contract.OrderDTO) error
}