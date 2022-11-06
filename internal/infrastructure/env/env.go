package env

import (
	"log"

	goEnv "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Environment struct {
	Port string `env:"PORT"`
}

func GetEnvironment() *Environment {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var environment Environment

	_, err = goEnv.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}

	return &environment
}
