package env

import (
	"log"

	goEnv "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Environment struct {
	Port          int    `env:"PORT"`
	AppEnv        string `env:"ENVIRONMENT"`
	CollectorUrl  string `env:"COLLECTOR_URL"`
	ServiceName   string `env:"SERVICE_NAME"`
	GitServiceUrl string `env:"GIT_SERVICE_URL"`
}

func GetEnvironment() *Environment {
	_ = godotenv.Load()

	var environment Environment
	_, err := goEnv.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}

	return &environment
}
