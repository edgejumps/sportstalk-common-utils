package redisfactory

import (
	"context"
	"errors"
	"testing"
)

func TestRedisClientFactoryWithInvalidConfig(t *testing.T) {
	config := "invalid"

	_, err := RedisClientFactory(context.Background(), config)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if !errors.Is(err, ErrInvalidRedisClientConfig) {
		t.Errorf("Expected ErrInvalidRedisClientConfig, got %v", err)
	}
}

func TestRedisClientFactoryWithValidConfig(t *testing.T) {
	config := Config{
		Addr:     "127.0.0.1:6379",
		Password: "simonwang",
	}

	client, err := RedisClientFactory(context.Background(), config)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if client == nil {
		t.Errorf("Expected client, got nil")
	}
}
