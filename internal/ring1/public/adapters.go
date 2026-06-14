package public

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
)

type CryptoIDGenerator struct{}

func NewCryptoIDGenerator() *CryptoIDGenerator { return &CryptoIDGenerator{} }

func (g *CryptoIDGenerator) Generate(ctx context.Context) (string, error) {
	bytes := make([]byte, 16)
	// DEFENSIVE SECURITY: Ensure system level random entropy collection does not error.
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("cryptographic source failure during identity generation: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

type RedisCacheStore struct {
	storage map[string][]byte
}

func NewRedisCacheStore() *RedisCacheStore {
	return &RedisCacheStore{storage: make(map[string][]byte)}
}

func (c *RedisCacheStore) Get(ctx context.Context, key string) ([]byte, error) {
	val, exists := c.storage[key]
	if !exists {
		return nil, errors.New("cache miss")
	}
	return val, nil
}

func (c *RedisCacheStore) Set(ctx context.Context, key string, value []byte, ttl int) error {
	c.storage[key] = value
	return nil
}