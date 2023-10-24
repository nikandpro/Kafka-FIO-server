package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DBPath string

	Messages Messages
}

type Messages struct {
	Errors
}

type Errors struct {
	Default string `mapstructure:"default"`
}

func Init() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {

		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	godotenv.Load(".env")

	if err := viper.BindEnv("connection_db"); err != nil {
		return err
	}

	cfg.DBPath = viper.GetString("connection_db")

	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
