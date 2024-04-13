package config

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Redis  RedisConfig
	Secret string
}

type RedisConfig struct {
	Password string
	Addr     string
	Username string
	Db       int
}

type ApiContext struct {
	Config Config
	Client *redis.Client
}

func Get() Config {

	return Config{
		Redis: RedisConfig{
			Password: getEnv("REDIS_PW"),
			Addr:     getEnv("REDIS_ADDR"),
			Username: getEnv("REDIS_USR"),
			Db:       3,
		},
		Secret: getEnv("SECRET"),
	}
}

func getEnv(key string) string {
	value, set := os.LookupEnv(key)
	if !set {
		log.Fatalf("Config variable %s was missing\n", key)
	}
	return value
}
