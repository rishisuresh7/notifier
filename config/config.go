package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	RedisConfig  *redisConfig
}

func NewConfig() (*Config, error) {
	conf, missing, err := getConfigFromEnv()
	if err != nil {
		return nil, fmt.Errorf("NewConfig: %s", err.Error())
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("NewConfig: missing env argument(s): %s", strings.Join(missing, ", "))
	}

	return conf, nil
}

func getConfigFromEnv() (*Config, []string, error) {
	var missing []string

	redisConfig := newRedisConfig(&missing)

	if len(missing) > 0 {
		return nil, missing, nil
	}

	redisPort, err := strconv.Atoi(redisConfig.portString)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value provided for REDIS_PORT")
	}

	redisDatabase, err := strconv.Atoi(redisConfig.dbString)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value provided for REDIS_DATABASE_ID")
	}

	redisConfig.Port = redisPort
	redisConfig.Database = redisDatabase

	return &Config{
		RedisConfig:  redisConfig,
	}, nil, nil
}

func getEnv(key string) (string, bool) {
	value := os.Getenv(key)
	if value == "" {
		return "", false
	}

	return value, true
}
