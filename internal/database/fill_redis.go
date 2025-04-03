package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"hafez-horoscope-api/internal/models"
	"log"
)

func FillRedis(db *gorm.DB, rdb *redis.Client) {
	ctx := context.Background()

	// Fetch poems from MariaDB
	var poems []models.Poem
	if err := db.Table("horoscope.hafez").Find(&poems).Error; err != nil {
		log.Fatal("Error fetching poems:", err)
	}

	for _, poem := range poems {
		// Convert poem struct to JSON
		jsonValue, err := json.Marshal(poem)
		if err != nil {
			log.Println("Error marshaling JSON:", err)
			continue
		}

		// Store in Redis with ID as the key
		key := fmt.Sprintf("%d", poem.ID)
		err = rdb.Set(ctx, key, jsonValue, 0).Err()
		if err != nil {
			log.Println("Error storing in Redis:", err)
		} else {
			log.Println("Stored poem in Redis with key:", key)
		}
	}
}
