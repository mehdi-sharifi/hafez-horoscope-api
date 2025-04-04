package services

import (
	"context"
	"encoding/json"
	"fmt"
	"hafez-horoscope-api/config"
	"hafez-horoscope-api/internal/models"
	"hafez-horoscope-api/utils"
	"log"
	"math/rand"
	"time"
)

type RandomPoemService struct {
	RedisClient *redis.RedisClient
}

func NewRandomPoemService(redisClient *redis.RedisClient) *RandomPoemService {
	return &RandomPoemService{RedisClient: redisClient}
}

// Load config at package level to avoid redundant file reads
var cnf, err = config.LoadConfig("config/config.toml")

func init() {
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
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

	// Use randID for URLs since filenames follow id.mp3 & id.jpg format
	poem.AudioURL = fmt.Sprintf("%s/%d.mp3", cnf.Minio.AudioEndpoint, randID)
	poem.ImageURL = fmt.Sprintf("%s/%d.jpg", cnf.Minio.ImageEndpoint, randID)

	return &poem, nil
}
