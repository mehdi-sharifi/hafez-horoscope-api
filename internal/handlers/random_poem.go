package handlers

import (
	"github.com/gofiber/fiber/v2"
	"hafez-horoscope-api/config"
	"hafez-horoscope-api/internal/services"
	redis "hafez-horoscope-api/utils"
	"log"
)

var cnf, err = config.LoadConfig("config/config.toml")
var redisClient *redis.RedisClient

func init() {
	if err != nil {
		log.Fatal(err)
	}
	redisClient, err = redis.GetRedisClient(cnf)
}

func GetRandomPoem(c *fiber.Ctx) error {

	randomPoemService := services.NewRandomPoemService(redisClient)

	// Fetch a random poem
	poem, err := randomPoemService.GetRandomPoem()
	if err != nil {
		log.Fatal("Error fetching random poem:", err)
	}
	return c.JSON(poem)
}
