package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	ServerConfig struct {
		ServerPort string `mapstructure:"SERVER_PORT"`

		DatabaseHost     string `mapstructure:"DATABASE_HOST"`
		DatabasePort     string `mapstructure:"DATABASE_PORT"`
		DatabaseUsername string `mapstructure:"DATABASE_USERNAME"`
		DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
		DatabaseName     string `mapstructure:"DATABASE_NAME"`

		SecretKeyToken string `mapstructure:"SECRET_KEY_TOKEN"`
	}
)

func LoadServerConfig() (*ServerConfig, error) {
	serverConfig := ServerConfig{}

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return &serverConfig, fmt.Errorf("error reading env file, %v", err)
	}

	if err := viper.Unmarshal(&serverConfig); err != nil {
		return &serverConfig, fmt.Errorf("unknow error, %v", err)
	}

	return &serverConfig, nil
}
