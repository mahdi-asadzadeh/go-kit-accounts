package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadSettings(mode string) {
	var err error

	if mode == "dev" {
		err = godotenv.Load("../pkg/config/envs/dev.env")

	} else {
		err = godotenv.Load("../pkg/config/envs/pro.env")
	}
	if err != nil {
		log.Fatal(err)
	}
}