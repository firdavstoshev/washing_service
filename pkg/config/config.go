package config

import "github.com/spf13/viper"

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	Port     int    `mapstructure:"port"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
	Server   ServerConfig   `mapstructure:"server"`
}

func Init() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
