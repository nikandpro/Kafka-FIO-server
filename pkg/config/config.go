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
	// cfg, err := LoadConfig(".env")
	// if err != nil {
	// 	log.Fatal("? Could not load environment variables", err)
	// }
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

	if err := viper.BindEnv("db"); err != nil {
		return err
	}

	cfg.DBPath = viper.GetString("db")

	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigType("env")
	viper.SetConfigName("main")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
