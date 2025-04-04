package services

import (
	"context"
	"encoding/json"
	"fmt"
	"hafez-horoscope-api/internal/models"
	"hafez-horoscope-api/utils"
	"math/rand"
	"time"
)

type RandomPoemService struct {
	RedisClient *redis.RedisClient
}

func NewRandomPoemService(redisClient *redis.RedisClient) *RandomPoemService {
	return &RandomPoemService{RedisClient: redisClient}
}

func (s *RandomPoemService) GetRandomPoem() (*models.Poem, error) {
	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())
	randID := rand.Intn(495) + 1 // Ensure range is 1-495
	key := fmt.Sprintf("%d", randID)

	jsonValue, err := s.RedisClient.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("error fetching poem from Redis: %w", err)
	}

	var poem models.Poem
	if err := json.Unmarshal([]byte(jsonValue), &poem); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return &poem, nil
}
