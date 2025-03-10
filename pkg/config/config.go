package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

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
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Failed to read config.yml: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if port := os.Getenv("SERVER_PORT"); port != "" {
		cfg.Server.Port = port
	}
	if host := os.Getenv("POSTGRES_HOST"); host != "" {
		cfg.Postgres.Host = host
	}
	if dbname := os.Getenv("POSTGRES_DBNAME"); dbname != "" {
		cfg.Postgres.DBName = dbname
	}
	if user := os.Getenv("POSTGRES_USER"); user != "" {
		cfg.Postgres.User = user
	}
	if password := os.Getenv("POSTGRES_PASSWORD"); password != "" {
		cfg.Postgres.Password = password
	}
	if port := os.Getenv("POSTGRES_PORT"); port != "" {
		cfg.Postgres.Port = viper.GetInt("postgres.port")
	}
	if sslMode := os.Getenv("POSTGRES_SSL_MODE"); sslMode != "" {
		cfg.Postgres.SSLMode = sslMode
	}

	return &cfg, nil
}
