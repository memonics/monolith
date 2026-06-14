package pure

import (
	"errors"
)

var ErrNegativeEvaluation = errors.New("analytical metrics cannot evaluate a negative transaction volume")

type RiskMetrics struct {
	IsFlagged bool
	RiskScore float64
}

// EvaluateRisk runs pure mathematical business rules to flag anomalous transactions.
func EvaluateRisk(amount float64) (RiskMetrics, error) {
	if amount < 0 {
		return RiskMetrics{}, ErrNegativeEvaluation
	}

	// Invariant: Single transactions over $10,000 are automatically flagged for review
	isFlagged := amount > 10000.0
	score := amount * 0.02

	return RiskMetrics{
		IsFlagged: isFlagged,
		RiskScore: score,
	}, nil
}