package main

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type DatabaseConfig struct {
	UserName string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"3306"`
	DBName   string `env:"DB_NAME"`
}

func main() {
	cfg := new(DatabaseConfig)

	if err := env.Parse(cfg); err != nil {
		log.Printf("%+v\n", err)
	}

	log.Printf("%+v\n", cfg)
}
