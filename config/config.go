package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort string

	MongoURI string
}

func GetConfig() *Config {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	httpPort, ok := viper.Get("HTTP_PORT").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	mongodbURI, ok := viper.Get("MONGODB_URI").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return &Config{
		HTTPPort: httpPort,
		MongoURI: mongodbURI,
	}
}
