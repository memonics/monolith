package pure

import (
	"errors"
	"time"
)

var ErrInvalidAmount = errors.New("order amount must be greater than zero")

type Order struct {
	ID        string
	UserID    string
	Amount    float64
	Status    string
	CreatedAt time.Time
}

// CreateOrder executes core domain logic with absolute zero infrastructure concerns.
func CreateOrder(id string, userID string, amount float64) (*Order, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	return &Order{
		ID:        id,
		UserID:    userID,
		Amount:    amount,
		Status:    "PENDING",
		CreatedAt: time.Now().UTC(), // Permitted standard utility usage (non-I/O)
	}, nil
}