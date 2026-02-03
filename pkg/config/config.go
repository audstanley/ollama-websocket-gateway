package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`
	Ollama struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"ollama"`
	Logging struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"logging"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.ollama-gateway")
	viper.AddConfigPath("/etc/ollama-gateway")

	viper.SetEnvPrefix("GATEWAY")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("server.port", "11436")
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("ollama.url", "http://localhost:11434")
	viper.SetDefault("logging.level", "info")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
