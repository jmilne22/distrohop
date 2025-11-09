package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Distro struct {
		Name     string   `mapstructure:"name"`
		Family   string   `mapstructure:"family"`
		Packages []string `mapstructure:"packages"`
	} `mapstructure:"package_managers"`
}

var Cfg Config

func LoadConfig() error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("$HOME/")

	// ---
	// Reading and Unmarshaling
	// ---

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config file not found")
		} else {
			log.Fatalf("Error reading config file: %s", err)
		}
	}

	if err := v.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	return nil
}
