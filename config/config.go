package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	SMSApiKey    string
	RedisConfig  *redisConfig

	EmailUsername string
	EmailPassword string
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
	smsApiKey, found := getEnv("SMS_API_KEY")
	if !found {
		missing = append(missing, "SMS_API_KEY")
	}

	emailUsername, found := getEnv("EMAIL_USERNAME")
	if !found {
		missing = append(missing, "EMAIL_USERNAME")
	}

	emailPassword, found := getEnv("EMAIL_PASSWORD")
	if !found {
		missing = append(missing, "EMAIL_PASSWORD")
	}

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
		SMSApiKey: smsApiKey,
		RedisConfig:  redisConfig,
		EmailPassword: emailPassword,
		EmailUsername: emailUsername,
	}, nil, nil
}

func getEnv(key string) (string, bool) {
	value := os.Getenv(key)
	if value == "" {
		return "", false
	}

	return value, true
}
