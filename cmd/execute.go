package cmd

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hafez-horoscope-api/config"
	"hafez-horoscope-api/internal/database"
	"log"
)

func initialize() {
	cnf, err := config.LoadConfig("config/config.toml")
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cnf.Database.Username, cnf.Database.Password, cnf.Database.Host, cnf.Database.Port, cnf.Database.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to MariaDB:", err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: cnf.Redis.Host + ":" + cnf.Redis.Port,
		DB:   0,
	})

	database.FillRedis(db, rdb)

}
func Execute() {
	initialize()
}
