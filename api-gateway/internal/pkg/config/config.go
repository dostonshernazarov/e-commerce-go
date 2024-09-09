package config

import (
	"os"
	"time"
)

type webAddress struct {
	Host string
	Port string
}

type Config struct {
	APP         string
	Environment string
	LogLevel    string
	APIToken    string
	ChatID      string
	Server      struct {
		Host         string
		Port         string
		ReadTimeout  string
		WriteTimeout string
		IdleTimeout  string
	}
	Context struct {
		Timeout string
	}
	Redis struct {
		Host     string
		Port     string
		Password string
		Name     string
	}
	Token struct {
		AccessTTL  time.Duration
		RefreshTTL time.Duration
		SignInKey  string
	}
	Minio struct {
		Host       string
		Port       string
		AccessKey  string
		SecretKey  string
		Location   string
		BucketName string
	}
	Kafka struct {
		Address []string
		Topic   struct {
			UserCreateTopic string
		}
	}
	OTLPCollector webAddress
}

func NewConfig() (*Config, error) {
	var config Config

	// general configuration
	config.APP = getEnv("APP", "app")
	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("LOG_LEVEL", "debug")
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "7s")

	// server configuration
	config.Server.Host = getEnv("SERVER_HOST", "api-service")
	config.Server.Port = getEnv("SERVER_PORT", ":9070")
	config.Server.ReadTimeout = getEnv("SERVER_READ_TIMEOUT", "10s")
	config.Server.WriteTimeout = getEnv("SERVER_WRITE_TIMEOUT", "10s")
	config.Server.IdleTimeout = getEnv("SERVER_IDLE_TIMEOUT", "120s")

	// redis configuration
	config.Redis.Host = getEnv("REDIS_HOST", "redis-db")
	config.Redis.Port = getEnv("REDIS_PORT", "6379")
	config.Redis.Password = getEnv("REDIS_PASSWORD", "")
	config.Redis.Name = getEnv("REDIS_DATABASE", "0")

	// minio configuration
	config.Minio.Host = getEnv("MINIO_HOST", "localhost")
	config.Minio.Port = getEnv("MINIO_PORT", ":9000")
	config.Minio.AccessKey = getEnv("MINIO_ROOT_USER", "minioadmin")
	config.Minio.SecretKey = getEnv("MINIO_ROOT_PASSWORD", "minioadmin")
	config.Minio.BucketName = getEnv("MINIO_DEFAULT_BUCKETS", "resumes")

	// otlp collector configuration
	config.OTLPCollector.Host = getEnv("OTLP_COLLECTOR_HOST", "localhost")
	config.OTLPCollector.Port = getEnv("OTLP_COLLECTOR_PORT", ":4317")

	return &config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
