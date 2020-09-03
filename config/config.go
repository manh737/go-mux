package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config is the server configuration structure.
// all fields will be filled with environment variables.
type Config struct {
	ServerHost    string // address that server will listening on
	MongoUser     string // mongo db username
	MongoPassword string // mongo db password
	MongoHost     string // host that mongo db listening on
	MongoPort     string // port that mongo db listening on
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
	value, ok = viper.Get("mongo_user").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	config.MongoUser = value
	value, ok = viper.Get("mongo_password").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	config.MongoPassword = value
	value, ok = viper.Get("mongo_host").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	config.MongoHost = value
	value, ok = viper.Get("mongo_port").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	config.MongoPort = value
}

// MongoURI will generate mongo db connect uri
func (config *Config) MongoURI() string {
	if config.MongoUser == "" && config.MongoPassword == "" {
		return fmt.Sprintf("mongodb://%s:%s",
			config.MongoHost,
			config.MongoPort,
		)

	}
	return fmt.Sprintf("mongodb://%s:%s@%s:%s",
		config.MongoUser,
		config.MongoPassword,
		config.MongoHost,
		config.MongoPort,
	)
}

// NewConfig will create and initialize config struct
func NewConfig() *Config {
	config := new(Config)
	config.initialize()
	return config
}