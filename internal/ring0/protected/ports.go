package protected

import (
	"context"
	"github.com/bdra-io/monolith/internal/ring0/pure"
)

type UserDTO struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserService interface {
	GetUser(ctx context.Context, id string) (*pure.User, error)
}