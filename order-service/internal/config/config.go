package config

import (
	"fmt"
)

type Config struct {
	App struct {
		Name        string
		Environment string
	}

	Postgres struct {
		Host     string
		Port     string
		Username string
		Password string
		Database string
		SSLMode  string
	}
	OrderServer struct {
		Http struct {
			Host string
			Port string
		}
	}
	UserServer struct {
		Http struct {
            Host string
            Port string
        }
	}
	ProductServer struct {
		Http struct {
			Host string
            Port string
		}
	}
}

func (c Config) LoadConfig() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Postgres.Username,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.Database,
		c.Postgres.SSLMode,
	)
}
