package configuration

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string `env:"PORT" envDefault:"9090"`
	DBHost      string `env:"DB_HOST"`
	DBPort      string `env:"DB_PORT"`
	DBDatabase  string `env:"DB_DATABASE"`
	DBUser      string `env:"DB_USER"`
	DBPassword  string `env:"DB_PASSWORD"`
	KafkaURI    string `env:"KAFKA_URI"`
	RabbitMQURI string `env:"RABBIT_MQ_URI"`
}

func Load() (*Config, error) {
	config := Config{}
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[WARN] .env file not found. Loading from system environment")
	}
	config.Port = os.Getenv("PORT")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBDatabase = os.Getenv("DB_DATABASE")
	config.KafkaURI = os.Getenv("KAFKA_URI")
	config.RabbitMQURI = os.Getenv("RABBIT_MQ_URI")
	return &config, nil
}
