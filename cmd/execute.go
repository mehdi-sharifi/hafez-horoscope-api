package cmd

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hafez-horoscope-api/config"
	"hafez-horoscope-api/internal/api"
	"hafez-horoscope-api/internal/services"
	"hafez-horoscope-api/utils"
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

func startServer() {
	router := api.SetupRouter()
	err = router.Listen(fmt.Sprintf(":%s", cnf.Server.Port))
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func cacheWarmer() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cnf.Database.Username, cnf.Database.Password, cnf.Database.Host, cnf.Database.Port, cnf.Database.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to MariaDB:", err)
	}
	services.FillRedis(db, redisClient)

}

func Execute() {
	cacheWarmer()
	startServer()
}
