package cache_redis_config

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateRedisConfig(config *RedisConfig) error {
	validate := validator.New()

	if err := validate.Struct(config); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			log.Printf("Validation error: %s", err)
		}
		return fmt.Errorf("invalid config: %w", err)
	}

	if config.Port <= 0 || config.Port > 65535 {
		return fmt.Errorf("invalid port: %d", config.Port)
	}

	if !strings.HasPrefix(config.Address, "tcp://") && !strings.HasPrefix(config.Address, "unix://") {
		return fmt.Errorf("invalid address: %s", config.Address)
	}

	return nil
}

type RedisConfig struct {
	Port   int    `json:"port" validate:"gt=0,lte=65535"`
	Address string `json:"address" validate:"required"`
}

type RedisClientConfig struct {
	RedisConfig
	Password string `json:"password" validate:"omitempty,len=40"`
}