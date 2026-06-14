package public

import (
	"context"
	"fmt"
	"log"

	ring1Contract "github.com/bdra-io/monolith/internal/ring1/protected"
	"github.com/bdra-io/monolith/internal/ring2/protected"
	"github.com/bdra-io/monolith/internal/ring2/pure"
)

type AuditLogServiceImpl struct{}

func NewAuditLogService() protected.AuditLogService {
	return &AuditLogServiceImpl{}
}

func (s *AuditLogServiceImpl) LogOrderCreation(ctx context.Context, order *ring1Contract.OrderDTO) error {
	// Execute Ring 2's internal pure logic boundaries out-of-band
	metrics, err := pure.EvaluateRisk(order.Amount)
	if err != nil {
		return fmt.Errorf("operational intelligence analytics failure: %w", err)
	}

	log.Printf("[AUDIT ENGINE] Ring 2 evaluation complete. OrderID: %s, Risk Score: %.2f, Audit Flagged: %t", 
		order.ID, metrics.RiskScore, metrics.IsFlagged)
		
	return nil
}