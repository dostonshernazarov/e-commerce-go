package load

import "github.com/spf13/viper"

type Postgres struct {
	HOST     string
	PORT     int
	DBNAME   string
	PASSWORD string
}

type Redis struct {
	HOST string
	PORT int
}

type Config struct {
	Postgres Postgres
	Redis    Redis

	UserServiceHost string
	UserServicePort int
}

func LOAD(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := Config{
		Postgres: Postgres{
			HOST:     viper.GetString("postgres.host"),
			PORT:     viper.GetInt("postgres.port"),
			DBNAME:   viper.GetString("postgres.dbname"),
			PASSWORD: viper.GetString("postgres.password"),
		},
		Redis: Redis{
			HOST: viper.GetString("redis.host"),
			PORT: viper.GetInt("redis.port"),
		},

		UserServiceHost: viper.GetString("service.host"),
		UserServicePort: viper.GetInt("service.port"),
	}

	return &conf, nil
}
