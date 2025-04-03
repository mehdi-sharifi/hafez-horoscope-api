package cmd

import (
	"hafez-horoscope-api/config"
)

func Execute() {
	_, err := config.LoadConfig("config/config.toml")
	if err != nil {
		panic(err)
	}
}
