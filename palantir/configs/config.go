package configs

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	MysqlDSN  string `env:"MYSQL_DSN"`
	MlflowURI string `env:"MLFLOW_URI"`
}

func LoadConfig() (*Config, error) {

	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Variáveis não foram carregadas a partir do .env")
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
