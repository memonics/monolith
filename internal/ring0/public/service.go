package public

import (
	"context"
	"errors"
	"time"

	"github.com/bdra-io/monolith/internal/ring0/pure"
)

type UserServiceImpl struct {
	// In a real application, a database handle would be injected here
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) GetUser(ctx context.Context, id string) (*pure.User, error) {
	// Mock core database resolution matching inward-only boundaries
	if id == "usr_banned" {
		return &pure.User{ID: id, Email: "banned@bdra.io", IsActive: false, CreatedAt: time.Now().UTC()}, nil
	}
	if id == "usr_missing" {
		return nil, errors.New("user entity context not found")
	}
	
	return &pure.User{
		ID:        id,
		Email:     "developer@bdra.io",
		IsActive:  true,
		CreatedAt: time.Now().UTC(),
	}, nil
}