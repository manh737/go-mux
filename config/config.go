package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config is the server configuration structure.
// all fields will be filled with environment variables.
type Config struct {
	ServerHost string // address that server will listening on
}

// initialize will read environment variables and save them in config structure fields
func (config *Config) initialize() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	// read environment variables
	value, ok := viper.Get("server_host").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	config.ServerHost = value
}

// NewConfig will create and initialize config struct
func NewConfig() *Config {
	config := new(Config)
	config.initialize()
	return config
}
